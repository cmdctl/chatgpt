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
	response, err := client.CreateCompletion(context.Background(), openai.CompletionRequest{
		Model:  openai.GPT3TextDavinci003,
		Prompt: string(content),
    MaxTokens: 1000,
	})

	if err != nil {
    log.Fatalf("Error creating completion: %v", err)
	}

	os.Stdout.WriteString(response.Choices[0].Text)
}
