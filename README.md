# Homebox MCP Server

This repository contains the Homebox MCP Server, a Go application that acts as a Model Context Protocol (MCP) server for a [Homebox](https://github.com/sysadminsmedia/homebox) installation.

## Overview

The Homebox MCP Server allows AI applications, like large language models (LLMs), to connect with your Homebox data and perform actions on your behalf. It exposes the Homebox API as a set of tools that can be used by an MCP client.

For more detailed information about the server, including setup, usage, and a full list of supported features, please see the [README.md file in the `homebox-mcp-server` directory](./homebox-mcp-server/README.md).

## Getting Started

To get started with the Homebox MCP Server, please refer to the detailed instructions in the [server's README.md file](./homebox-mcp-server/README.md).

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.

## Proxmox Install Script

This repository includes an installation script for Proxmox VE hosts, `homebox-install.sh`. This script automates the process of setting up a new LXC container and installing both Homebox and the Homebox MCP Server.

### What it Does

The script will:

1.  Prompt you for configuration details for a new LXC container.
2.  Create a new Debian 11 based LXC container.
3.  Install Docker and Docker Compose inside the new container.
4.  Download and start the official Homebox Docker container.
5.  Copy the `homebox-mcp-server` into the container.
6.  Install Go and the necessary dependencies for the MCP server.
7.  Set up a systemd service to run the `homebox-mcp-server` automatically.

### How to Use

1.  Clone this repository to your Proxmox VE host.
2.  Make the script executable: `chmod +x homebox-install.sh`
3.  Run the script: `./homebox-install.sh`
4.  Follow the on-screen prompts to configure your new Homebox instance.
