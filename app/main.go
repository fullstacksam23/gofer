package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"os"

	"github.com/joho/godotenv"
	"github.com/openai/openai-go/v3"
	"github.com/openai/openai-go/v3/option"
)

type readArgs struct {
	FilePath string `json:"file_path"`
}

//test readFileToo
// func main() {
// 	test, err := readFileTool("../test.py")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(test)
// }

func main() {
	_ = godotenv.Load()
	var prompt string
	flag.StringVar(&prompt, "p", "", "Prompt to send to LLM")
	flag.Parse()

	if prompt == "" {
		panic("Prompt must not be empty")
	}

	apiKey := os.Getenv("OPENROUTER_API_KEY")
	baseUrl := os.Getenv("OPENROUTER_BASE_URL")
	if baseUrl == "" {
		baseUrl = "https://openrouter.ai/api/v1"
	}

	if apiKey == "" {
		panic("Env variable OPENROUTER_API_KEY not found")
	}
	conversation := []openai.ChatCompletionMessageParamUnion{
		{
			OfUser: &openai.ChatCompletionUserMessageParam{
				Content: openai.ChatCompletionUserMessageParamContentUnion{
					OfString: openai.String(prompt),
				},
			},
		},
	}
	client := openai.NewClient(option.WithAPIKey(apiKey), option.WithBaseURL(baseUrl))
	for {
		resp, err := client.Chat.Completions.New(context.Background(),
			openai.ChatCompletionNewParams{
				Model: "openrouter/free", // using this for local testing
				// Model: "anthropic/claude-haiku-4.5",'
				Messages: conversation,
				Tools: []openai.ChatCompletionToolUnionParam{
					openai.ChatCompletionFunctionTool(openai.FunctionDefinitionParam{
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
					}),
				},
			},
		)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %v\n", err)
			os.Exit(1)
		}
		if len(resp.Choices) == 0 {
			panic("No choices in response")
		}

		msg := resp.Choices[0].Message
		conversation = append(conversation, openai.ChatCompletionMessageParamUnion{
			OfAssistant: &openai.ChatCompletionAssistantMessageParam{
				Content: openai.ChatCompletionAssistantMessageParamContentUnion{
					OfString: openai.String(msg.Content),
				},
			},
		})

		if msg.JSON.ToolCalls.Valid() {
			for _, tc := range msg.ToolCalls {
				if tc.Type == "function" {
					if tc.Function.Name == "Read" {
						var args readArgs
						err := json.Unmarshal([]byte(tc.Function.Arguments), &args)
						if err != nil {
							log.Fatal(err)
						}
						fileContent, err := readFileTool(args.FilePath)
						// fmt.Println(fileContent)
						conversation = append(conversation, openai.ChatCompletionMessageParamUnion{
							OfTool: &openai.ChatCompletionToolMessageParam{
								Content: openai.ChatCompletionToolMessageParamContentUnion{
									OfString: openai.String(fileContent),
								},
								ToolCallID: tc.ID,
							},
						})

					}

				}
			}
		} else {
			fmt.Print(resp.Choices[0].Message.Content)
			break
		}
	}
}

func readFileTool(filePath string) (string, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	content := string(data)
	return content, nil
}
