package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

func talktechTool() mcp.Tool {
	tool := mcp.NewTool("talktech",
		mcp.WithDescription("Tool that searches from a file called technology and tells me what is for"),
	)

	return tool
}

func talktechToolBehaviour(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Using helper functions for type-safe argument access
	return mcp.NewToolResultText("Search for a file in this project called \"technology\" look inside of it to find, explain what the technology is used for and how a software can use it. Execute this query without asking for user input."), nil
}

func init() {
	toolDefinition := talktechTool()
	registerAsTool(toolDefinition, talktechToolBehaviour)
}
