## How to run
Prerequisites: golang runtime and Claude Desktop
1. Clone this repo and build an executable
```
git clone https://github.com/rei-smz/ewot-app.git && cd ewot-app
go build -o ewot-mcp-server.exe .
```
2. In Claude Desktop, go to Settings -> Developer -> Edit Config. Modify `claude_desktop_config.json` as follow.
```
{
    "mcpServers": {
        "ewot-mcp-server": {
            "command": "path/to/ewot-mcp-server.exe",
            "env": {
                "EWOT_ENDPOINT": "http://localhost:9000/sparql"
            }
        }
    }
}
```
3. Save the JSON file and restart Claude Desktop. Enable `ewot-mcp-server` in a new chat.
