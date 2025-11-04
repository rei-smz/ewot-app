package MCPHandler

import (
	"context"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"os"
)

func SystemPromptHandler(_ context.Context, _ *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	systemPromptText, err := ReadPrompt("system_prompt.txt")
	if err != nil {
		return nil, mcp.ResourceNotFoundError("system_prompt.txt")
	}

	return &mcp.GetPromptResult{
		Description: "system prompt",
		Messages: []*mcp.PromptMessage{
			{
				Role:    "system",
				Content: &mcp.TextContent{Text: systemPromptText},
			},
		},
	}, nil
}

func RDFSummaryHandler(_ context.Context, _ *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	rdfSummaryText, err := ReadPrompt("rdf_summary.txt")
	if err != nil {
		return nil, mcp.ResourceNotFoundError("rdf_summary.txt")
	}

	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{
				URI:      "file://rdf_summary.txt",
				Text:     rdfSummaryText,
				MIMEType: "text/plain",
			},
		},
	}, nil
}

func ReadPrompt(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
