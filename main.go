package main

import (
	"context"
	"ewot-app/MCPHandler"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"log"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "sparql", Version: "v0.0.1"}, nil)

	systemPrompt := mcp.Prompt{
		Name:        "system",
		Description: "system prompt",
	}
	server.AddPrompt(&systemPrompt, MCPHandler.SystemPromptHandler)

	rdfSummaryRes := mcp.Resource{
		Name:        "rdf-summary",
		Description: "rdf summary resource",
		MIMEType:    "text/plain",
	}
	server.AddResource(&rdfSummaryRes, MCPHandler.RDFSummaryHandler)

	mcp.AddTool(server, &mcp.Tool{Name: "query", Description: "send sparql query"}, MCPHandler.SPARQuery)

	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
