# Agent Instructions

This document provides instructions for AI agents working on the Homebox MCP Server project.

## Project Overview

This project is a Go application that acts as a Model Context Protocol (MCP) server for a [Homebox](https://github.com/sysadminsmedia/homebox) installation. The goal is to implement the entire Homebox API specification as a set of tools that can be used by an MCP client.

## Development Workflow

1.  **Understand the API**: The primary source of truth for the API is the Swagger 2.0 specification provided by the user. All new tools should be based on this specification.
2.  **Implement Tools Incrementally**: The API is large. Implement tools for one resource at a time (e.g., "Items", "Locations", "Labels").
3.  **Follow Existing Patterns**: The existing code in `main.go` establishes a clear pattern for defining input/output structs and tool functions. Follow this pattern for all new tools.
4.  **Update `main.go`**: All new tools should be registered in the `main` function.
5.  **Update `TODO.md`**: As you implement new endpoints, make sure to mark them as complete in the `TODO.md` file.

## Environment Notes

*   **Go Version**: The project uses Go 1.18 or later.
*   **Dependencies**: Dependencies are managed with Go modules. Run `go mod tidy` in the `homebox-mcp-server` directory to install or update dependencies.
*   **Running the Server**: To run the server, you need to set the `HOMEBOX_URL` and `HOMEBOX_TOKEN` environment variables. Then, from the `homebox-mcp-server` directory, run `go run main.go`.
*   **Execution Environment**: The execution environment for `run_in_bash_session` has been unreliable. Do not attempt to compile or run the code unless you have a clear understanding of the file system context. Focus on writing correct code that can be submitted for the user to run and test.
