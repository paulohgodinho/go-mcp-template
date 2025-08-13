package main

import (
	"fmt"
	"log"

	"mcpdev/prompts"
	"mcpdev/tools"

	"github.com/mark3labs/mcp-go/server"
)

func main() {

	s := server.NewMCPServer(
		"Calculator Demo",
		"1.0.0",
		server.WithToolCapabilities(false),
		server.WithRecovery(),
	)

	tools.RegisterTools(s)
	prompts.RegisterPrompts(s)

	// address := "127.0.0.1:58999"
	// streamableServer := server.NewStreamableHTTPServer(s)
	// log.Println("Starting server on", address)
	// streamableServer.Start(address)

	log.Println("Starting server on stdin/stdout")
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
