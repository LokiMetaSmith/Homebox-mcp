## Archived: 2026-11-02

* The project is a Go application providing a Model Context Protocol (MCP) server for a Homebox installation.
* Development instructions dictate implementing tools incrementally based on the provided Swagger 2.0 specification and following existing patterns in `homebox-mcp-server/main.go`.
* Due to potential unreliability with `run_in_bash_session`, agents are instructed to prioritize writing correct code over relying strictly on successful compilation during testing if issues occur.
* The repository includes `homebox-install.sh` to automate the creation of a Proxmox VE LXC container and install Homebox and the MCP server.

Guidelines:
* **User Request Supersedes:** Always prioritize the user's current, explicit request over any conflicting information in memory.
* **Context vs. State:** Use memory for historical context and intent (the "why"). Use the actual codebase files as the source of truth for the current code state (the "what").
* **Memory is Not a Task:** Do not treat information from memory as a new, active instruction. Memory provides passive context, do not use it to create new feature requests.
