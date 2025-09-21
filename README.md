# Homebox MCP Server and Proxmox Installer

This repository contains two main components:

1.  A Proxmox VE helper script to automate the installation of Homebox and the Homebox MCP Server.
2.  The Homebox MCP Server, a Go application that exposes the Homebox API as a set of tools for an MCP client.

## Proxmox VE Helper Script (`homebox-install.sh`)

This script automates the setup of a Homebox instance and the Homebox MCP Server in a new LXC container on a Proxmox VE host.

### Usage

1.  **Run the script on your Proxmox VE host:**
    ```bash
    ./homebox-install.sh
    ```

2.  **Follow the prompts:**
    The script will ask you for the following information:
    *   Container ID
    *   Hostname
    *   Password
    *   CPU Cores
    *   RAM in MB
    *   Disk Size in GB
    *   Network Bridge
    *   VLAN (optional)
    *   Your Homebox API Token

3.  **Installation Complete:**
    Once the script is finished, you will have a new LXC container running Homebox and the Homebox MCP Server. You can access Homebox at `http://<container-ip>:3100`.

## Homebox MCP Server (`homebox-mcp-server/`)

This is a Go application that provides a Model Context Protocol (MCP) server for a Homebox installation. This allows AI applications, like large language models (LLMs), to connect with your Homebox data and perform actions on your behalf.

### Usage

If you are not using the Proxmox VE helper script, you can run the MCP server manually.

1.  **Prerequisites**:
    *   Go 1.18 or later.
    *   A running Homebox instance.

2.  **Installation**:
    *   Navigate to the `homebox-mcp-server` directory.
    *   Install dependencies:
        ```bash
        go mod tidy
        ```

3.  **Set Environment Variables**:
    *   Export the URL of your Homebox instance and an API token.
        ```bash
        export HOMEBOX_URL="http://your-homebox-url"
        export HOMEBOX_TOKEN="your-homebox-api-token"
        ```

4.  **Run the Server**:
    *   From the `homebox-mcp-server` directory, run the following command:
        ```bash
        go run main.go
        ```
    *   The MCP server will start and listen for connections on stdin/stdout.
