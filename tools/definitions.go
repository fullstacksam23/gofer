package tools

import "github.com/openai/openai-go/v3"

var ReadFileDefinition = openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
	Name:        "Read",
	Description: openai.String("Read and return the contents of a file"),
	Parameters: openai.FunctionParameters{
		"type": "object",
		"properties": map[string]any{
			"file_path": map[string]any{
				"type":        "string",
				"description": "The path to the file to read",
			},
		},
		"required": []string{"file_path"},
	},
})

var WriteFileDefinition = openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
	Name:        "Write",
	Description: openai.String("Write content to a file"),
	Parameters: openai.FunctionParameters{
		"type": "object",
		"properties": map[string]any{
			"file_path": map[string]any{
				"type":        "string",
				"description": "The path of the file to write to",
			},
			"content": map[string]any{
				"type":        "string",
				"description": "The content to write to the file",
			},
		},
		"required": []string{"file_path", "content"},
	},
})
var BashDefinition = openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
	Name:        "Bash",
	Description: openai.String("Execute a shell command"),
	Parameters: openai.FunctionParameters{
		"type": "object",
		"properties": map[string]any{
			"command": map[string]any{
				"type":        "string",
				"description": "The command to execute",
			},
		},
		"required": []string{"command"},
	},
})
