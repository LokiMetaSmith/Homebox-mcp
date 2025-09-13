# Homebox MCP Server

This project provides a Model Context Protocol (MCP) server for a [Homebox](https://github.com/sysadminsmedia/homebox) installation. This allows AI applications, like large language models (LLMs), to connect with your Homebox data and perform actions on your behalf.

## Features

This MCP server exposes the Homebox API as a set of tools that can be used by an MCP client. The following resources are currently supported:

*   **Items**: CRUD (Create, Read, Update, Delete) operations, plus other utilities like duplicate, export, etc.
*   **Locations**: CRUD operations.
*   **Labels**: CRUD operations.
*   **Item Maintenance**: Get log and create entries.
*   **Actions**: Various server-wide actions like creating thumbnails.
*   **Status & Currency**: Get server status and currency information.

A full list of implemented tools can be found in the source code (`main.go`). A list of remaining endpoints to be implemented can be found in `TODO.md`.

## Setup

1.  **Prerequisites**:
    *   Go 1.18 or later.
    *   A running Homebox instance.

2.  **Installation**:
    *   Clone this repository.
    *   Navigate to the `homebox-mcp-server` directory.
    *   Install dependencies:
        ```bash
        go mod tidy
        ```

## Usage

1.  **Set Environment Variables**:
    *   Export the URL of your Homebox instance and an API token.
        ```bash
        export HOMEBOX_URL="http://your-homebox-url"
        export HOMEBOX_TOKEN="your-homebox-api-token"
        ```

2.  **Run the Server**:
    *   From the `homebox-mcp-server` directory, run the following command:
        ```bash
        go run main.go
        ```
    *   The MCP server will start and listen for connections on stdin/stdout.

3.  **Connect a Client**:
    *   You can now connect any MCP-compatible client to this server.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.
