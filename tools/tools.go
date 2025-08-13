package tools

import (
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type ToolDefinition struct {
	Tool    mcp.Tool
	Handler server.ToolHandlerFunc
}

var Tools []ToolDefinition

func registerAsTool(tool mcp.Tool, handler server.ToolHandlerFunc) {
	Tools = append(Tools, ToolDefinition{
		Tool:    tool,
		Handler: handler,
	})
}

func RegisterTools(s *server.MCPServer) {
	for _, t := range Tools {
		log.Printf("Registering tool: %s", t.Tool.Name)
		s.AddTool(t.Tool, t.Handler)
	}
}
