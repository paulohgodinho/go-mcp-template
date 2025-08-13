package prompts

import (
	"log"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

type PromptDefinition struct {
	Prompt  mcp.Prompt
	Handler server.PromptHandlerFunc
}

var Prompts []PromptDefinition

func registerAsPrompt(prompt mcp.Prompt, handler server.PromptHandlerFunc) {
	Prompts = append(Prompts, PromptDefinition{
		Prompt:  prompt,
		Handler: handler,
	})
}

func RegisterPrompts(s *server.MCPServer) {
	for _, p := range Prompts {
		log.Printf("Registering prompt: %s", p.Prompt.Name)
		s.AddPrompt(p.Prompt, p.Handler)
	}
}
