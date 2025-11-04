package MCPHandler

import (
	"context"
	"ewot-app/Query"
	"fmt"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"os"
)

type Input struct {
	Query string `json:"query" jsonschema:"SPARQL query text"`
}

type Output struct {
	Answer string `json:"answer" jsonschema:"SPARQL query answer"`
}

func SPARQuery(ctx context.Context, req *mcp.CallToolRequest, input Input) (*mcp.CallToolResult, Output, error) {
	connectorInstance := Query.GetSPARQLConnector()
	answer, err := connectorInstance.SendQuery(input.Query)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error sending query: %v, error: %v\n", input.Query, err)
		return nil, Output{}, err
	}

	return nil, Output{Answer: answer}, nil
}
