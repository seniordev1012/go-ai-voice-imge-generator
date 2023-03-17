package main

import (
	"bufio"
	"context"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"os"
	"strings"
)

// makeApiCall is a function that makes an API call to the Ron Swanson Quote API
// and returns the quote as a string.
// The function returns an error if the API call fails.
// The function returns a string if the API call is successful.
// The function returns an empty string if the API call is successful but the
// quote is empty.
// e.g makeApiCall() (string, error)
// usage: quote, err := makeApiCall()
func makeApiCall(prompt string) (string, error) {
	OpenaiApiKey := os.Getenv("OPENAI")
	client := openai.NewClient(OpenaiApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleUser,
					Content: prompt,
				},
			},
		},
	)

	if err != nil {
		fmt.Printf("ChatCompletion error: %v\n", err)
		return "", nil
	}

	fmt.Println(resp.Choices[0].Message.Content)
	return resp.Choices[0].Message.Content, nil
}

func chatGPT3() (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI"))
	messages := make([]openai.ChatCompletionMessage, 0)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Conversation")
	fmt.Println("---------------------")

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleUser,
			Content: text,
		})

		resp, err := client.CreateChatCompletion(
			context.Background(),
			openai.ChatCompletionRequest{
				Model:    openai.GPT3Dot5Turbo,
				Messages: messages,
			},
		)

		if err != nil {
			fmt.Printf("ChatCompletion error: %v\n", err)
			continue
		}

		content := resp.Choices[0].Message.Content
		messages = append(messages, openai.ChatCompletionMessage{
			Role:    openai.ChatMessageRoleAssistant,
			Content: content,
		})
		fmt.Println(content)
	}
}
