package main

import (
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
)

func gpt(secret, message string) {
	client := openai.NewClient(secret)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: message,
				},
			},
		},
	)

	if err != nil {
		return
	}

	fmt.Println(resp.Choices[0].Message.Content)
}