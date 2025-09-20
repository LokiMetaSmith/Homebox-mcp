#!/usr/bin/env bash

# This script creates a new Proxmox VE LXC container and installs Homebox and the Homebox MCP Server.
#
# Copyright (c) 2025 The Proxmox VE Helper-Scripts Authors (https://github.com/community-scripts/ProxmoxVE)
#
# Permission is hereby granted, free of charge, to any person obtaining a copy
# of this software and associated documentation files (the "Software"), to deal
# in the Software without restriction, including without limitation the rights
# to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
# copies of the Software, and to permit persons to whom the Software is
# furnished to do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all
# copies or substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
# IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
# FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
# AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
# LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
# SOFTWARE.

# --- Configuration ---
YW_OK() { echo -e " \e[32m✔\e[0m ${1}"; }
YW_ERR() { echo -e " \e[31m✖\e[0m ${1}"; }
YW_INFO() { echo -e " \e[36mℹ\e[0m ${1}"; }
YW_WARN() { echo -e " \e[33m⚠\e[0m ${1}"; }

# --- Spinner ---
SPINNER="⠋⠙⠹⠸⠼⠴⠦⠧⠇⠏"
show_spinner() {
  local pid=$1
  local delay=0.1
  local i=0
  while ps -p $pid >/dev/null; do
    i=$(((i + 1) % ${#SPINNER}))
    printf "\r \e[36m%s\e[0m" "${SPINNER:$i:1}"
    sleep $delay
  done
  printf "\r\033[K"
}

# --- Check for Proxmox ---
if ! command -v pveversion >/dev/null 2>&1; then
  YW_ERR "This script must be run on a Proxmox VE host."
  exit 1
fi

# --- Main Logic ---
main() {
  YW_INFO "Starting Homebox LXC Container Setup..."

  # --- Get Container Configuration ---
  read -p "Enter Container ID: " CT_ID
  read -p "Enter Hostname: " CT_HOSTNAME
  read -p "Enter Password: " CT_PASSWORD
  read -p "Enter CPU Cores: " CT_CORES
  read -p "Enter RAM in MB: " CT_RAM
  read -p "Enter Disk Size in GB: " CT_DISK_SIZE
  read -p "Enter Bridge: " CT_BRIDGE
  read -p "Enter VLAN (optional): " CT_VLAN
  read -p "Enter your Homebox API Token: " HOMEBOX_TOKEN

  # --- Create LXC Container ---
  YW_INFO "Creating LXC Container..."
  pct create $CT_ID /var/lib/vz/template/cache/debian-11-standard_11.3-1_amd64.tar.gz \
    --hostname $CT_HOSTNAME \
    --password $CT_PASSWORD \
    --cores $CT_CORES \
    --memory $CT_RAM \
    --swap 0 \
    --rootfs local-lvm:$CT_DISK_SIZE \
    --net0 name=eth0,bridge=$CT_BRIDGE,ip=dhcp \
    --features nesting=1,keyctl=1 \
    --onboot 1 \
    --start 1 &> /dev/null &

  SPINNER_PID=$!
  show_spinner $SPINNER_PID
  wait $SPINNER_PID
  YW_OK "LXC Container created."

  # --- Install Docker ---
  YW_INFO "Installing Docker..."
  pct exec $CT_ID -- bash -c "apt-get update && apt-get install -y curl" &> /dev/null
  pct exec $CT_ID -- bash -c "curl -fsSL https://get.docker.com -o get-docker.sh" &> /dev/null
  pct exec $CT_ID -- bash -c "sh get-docker.sh" &> /dev/null &

  SPINNER_PID=$!
  show_spinner $SPINNER_PID
  wait $SPINNER_PID
  YW_OK "Docker installed."

  # --- Install Docker Compose ---
  YW_INFO "Installing Docker Compose..."
  pct exec $CT_ID -- bash -c "curl -L 'https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)' -o /usr/local/bin/docker-compose" &> /dev/null
  pct exec $CT_ID -- bash -c "chmod +x /usr/local/bin/docker-compose" &> /dev/null &

  SPINNER_PID=$!
  show_spinner $SPINNER_PID
  wait $SPINNER_PID
  YW_OK "Docker Compose installed."

  # --- Setup Homebox ---
  YW_INFO "Setting up Homebox..."
  pct exec $CT_ID -- bash -c "mkdir -p /opt/homebox" &> /dev/null
  pct exec $CT_ID -- bash -c "curl -L 'https://raw.githubusercontent.com/sysadminsmedia/homebox/main/docker-compose.yml' -o /opt/homebox/docker-compose.yml" &> /dev/null
  pct exec $CT_ID -- bash -c "mkdir -p /path/to/data/folder" &> /dev/null
  pct exec $CT_ID -- bash -c "chown 65532:65532 -R /path/to/data/folder" &> /dev/null
  pct exec $CT_ID -- bash -c "cd /opt/homebox && docker-compose up -d" &> /dev/null &

  SPINNER_PID=$!
  show_spinner $SPINNER_PID
  wait $SPINNER_PID
  YW_OK "Homebox setup complete."

  # --- Setup homebox-mcp-server ---
  YW_INFO "Setting up homebox-mcp-server..."
  pct push $CT_ID homebox-mcp-server /opt/homebox-mcp-server &> /dev/null

  YW_INFO "Installing Go..."
  pct exec $CT_ID -- bash -c "apt-get update && apt-get install -y golang" &> /dev/null &
  SPINNER_PID=$!
  show_spinner $SPINNER_PID
  wait $SPINNER_PID
  YW_OK "Go installed."

  YW_INFO "Installing dependencies for homebox-mcp-server..."
  pct exec $CT_ID -- bash -c "cd /opt/homebox-mcp-server && go mod tidy" &> /dev/null &
  SPINNER_PID=$!
  show_spinner $SPINNER_PID
  wait $SPINNER_PID
  YW_OK "Dependencies installed."

  YW_INFO "Creating systemd service for homebox-mcp-server..."
  pct exec $CT_ID -- bash -c "cat <<EOF > /etc/systemd/system/homebox-mcp.service
[Unit]
Description=Homebox MCP Server
After=network.target

[Service]
Type=simple
User=root
WorkingDirectory=/opt/homebox-mcp-server
ExecStart=/usr/bin/go run main.go
Environment=\"HOMEBOX_URL=http://localhost:3100\"
Environment=\"HOMEBOX_TOKEN=$HOMEBOX_TOKEN\"
Restart=on-failure

[Install]
WantedBy=multi-user.target
EOF" &> /dev/null

  YW_INFO "Starting homebox-mcp-server service..."
  pct exec $CT_ID -- bash -c "systemctl daemon-reload && systemctl enable --now homebox-mcp.service" &> /dev/null &
  SPINNER_PID=$!
  show_spinner $SPINNER_PID
  wait $SPINNER_PID
  YW_OK "homebox-mcp-server service started."


  YW_INFO "Homebox is running at http://<container-ip>:3100"
  YW_INFO "Homebox MCP Server is running."
}

main
