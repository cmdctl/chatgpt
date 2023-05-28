package main

import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

func constructMessages(content string, delimiter string) []openai.ChatCompletionMessage {
	chunks := strings.Split(content, delimiter)
	messages := make([]openai.ChatCompletionMessage, len(chunks))
	for i, chunk := range chunks {
		// role is either "user" or "ai" if i is even or odd
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
	content, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal("Error reading from stdin")
	}
	key := os.Getenv("OPENAI_KEY")
	client := openai.NewClient(key)
	response, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model:       openai.GPT3Dot5Turbo,
		Messages:    constructMessages(string(content), delimiter),
		Temperature: 0,
	})

	if err != nil {
		log.Fatalf("Error creating completion: %v", err)
	}

	os.Stdout.WriteString(response.Choices[0].Message.Content + "\n")
}
