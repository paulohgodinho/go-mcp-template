package prompts

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
)

func readMePrompt() mcp.Prompt {
	prompt := mcp.NewPrompt("generateReadMe",
		mcp.WithPromptDescription("Generate a README file"),
		mcp.WithArgument("project_name",
			mcp.ArgumentDescription("Name of the project"),
		),
		mcp.WithArgument("author",
			mcp.ArgumentDescription("Author of the project"),
		),
	)

	return prompt
}

func readMePromptMeBehaviour(ctx context.Context, request mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	projectName := request.Params.Arguments["project_name"]
	author := request.Params.Arguments["author"]

	if projectName == "" {
		projectName = "Unknown Project"
	}
	if author == "" {
		author = "Unknown Author"
	}

	return mcp.NewGetPromptResult(
		"Generate a README file",
		[]mcp.PromptMessage{
			mcp.NewPromptMessage(
				mcp.RoleAssistant,
				mcp.NewTextContent(fmt.Sprintf("# %s\n\n## Author: %s\n\n## Description:\n\nThis is a README file for the %s project.", projectName, author, projectName)),
			),
		},
	), nil
}

func init() {
	prompt := readMePrompt()
	registerAsPrompt(prompt, readMePromptMeBehaviour)
}
