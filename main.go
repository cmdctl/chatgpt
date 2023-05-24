package main

import (
	"context"
	"io/ioutil"
	"log"
	"os"

	openai "github.com/sashabaranov/go-openai"
)

func main() {
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
    log.Fatal("Error reading from stdin")
	}
	key := os.Getenv("OPENAI_KEY")
	client := openai.NewClient(key)
	response, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:            openai.GPT3Dot5Turbo,
		Messages:         []openai.ChatCompletionMessage{
      {
      	Role:    "user",
      	Content: string(content),
      },
    },
		MaxTokens:        4000,
		Temperature:      0,
	})

	if err != nil {
    log.Fatalf("Error creating completion: %v", err)
	}

	os.Stdout.WriteString(response.Choices[0].Message.Content)
}

