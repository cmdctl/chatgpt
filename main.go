package main

import (
	"context"
	"flag"
	"io"
	"log"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func constructMessages(content string, delimiter string) []openai.ChatCompletionMessage {
	chunks := strings.Split(content, delimiter)
	messages := make([]openai.ChatCompletionMessage, len(chunks))
	for i, chunk := range chunks {
		role := "user"
		if i%2 == 0 {
			role = "assistant"
		}

		messages[i] = openai.ChatCompletionMessage{
			Role:    role,
			Content: chunk,
		}
	}
	return messages
}
func main() {
	var delimiter string
	flag.StringVar(&delimiter, "d", "", "Delimiter to use between messages")
	flag.Parse()
	content, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Error reading from stdin")
	}
	key := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(key)

  request := openai.ChatCompletionRequest{
		Model:       "gpt-4o",
		Messages:    constructMessages(string(content), delimiter),
		Temperature: 0.4,
	}

	response, err := client.CreateChatCompletion(context.Background(), request)

	if err != nil {
		log.Fatalf("Error creating completion: %v", err)
	}

	os.Stdout.WriteString(response.Choices[0].Message.Content + "\n")
}
