# Homebox MCP Server and Proxmox Installer

This repository contains two main components:

1.  `homebox-mcp-server`: A server that exposes the Homebox API as a set of tools for the Model-Context Protocol (MCP).
2.  `homebox-install.sh`: A script to automate the installation of Homebox and the `homebox-mcp-server` on a Proxmox VE host.

## Quick Install

This script will create a new Proxmox LXC container and install Homebox and the `homebox-mcp-server`.

1.  Make sure you are on a Proxmox VE host.
2.  Run the script:

    ```bash
    ./homebox-install.sh
    ```
3.  The script will prompt you for the container configuration details.

## `homebox-install.sh`

The `homebox-install.sh` script automates the setup of a complete Homebox environment. It will guide you through the process of creating a new LXC container and installing all the necessary components.

## `homebox-mcp-server`

For detailed instructions on how to install and run the `homebox-mcp-server` manually, please refer to the `README.md` in the `homebox-mcp-server` directory:

[homebox-mcp-server/README.md](./homebox-mcp-server/README.md)
