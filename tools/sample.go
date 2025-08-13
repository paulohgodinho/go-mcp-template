package tools

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func sampleTool() mcp.Tool {
	tool := mcp.NewTool("sample",
		mcp.WithDescription("Tool that says hello to a defined name"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("The name to greet"),
		),
	)

	return tool
}

func sampleToolBehaviour(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Using helper functions for type-safe argument access
	name, err := request.RequireString("name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(fmt.Sprintf("Sample tool executed with input: %s - Hi %s, hope you are well!", name, name)), nil
}

func init() {
	toolDefinition := sampleTool()
	registerAsTool(toolDefinition, sampleToolBehaviour)
}
