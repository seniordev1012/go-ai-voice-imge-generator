package aigenRest

import (
	"context"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"os"
)

// MakeApiCall is a function that makes an API call to the Ron Swanson Quote API
// and returns the quote as a string.
// The function returns an error if the API call fails.
// The function returns a string if the API call is successful.
// The function returns an empty string if the API call is successful but the
// quote is empty.
// e.g MakeApiCall() (string, error)
// usage: quote, err := MakeApiCall()
func MakeApiCall(prompt string) (string, error) {
	OpenaiApiKey := os.Getenv("OPENAI")
	client := openai.NewClient(OpenaiApiKey)
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo0301,
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
