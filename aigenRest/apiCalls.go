package aigenRest

import (
	"bufio"
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	openai "github.com/sashabaranov/go-openai"
	"image/png"
	"io"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
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

func ronSwan() (string, error) {

	url := "https://ron-swanson-quotes.herokuapp.com/v2/quotes"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println(string(body))
	return string(body[1 : len(body)-1]), nil

}

func ImageGenerationCall(prompt string) (string, error) {
	c := openai.NewClient(os.Getenv("OPENAI"))
	ctx := context.Background()

	// Example image as base64
	reqBase64 := openai.ImageRequest{
		Prompt:         prompt,
		Size:           openai.CreateImageSize512x512,
		ResponseFormat: openai.CreateImageResponseFormatB64JSON,
		N:              1,
	}

	respBase64, err := c.CreateImage(ctx, reqBase64)
	if err != nil {
		fmt.Printf("Image creation error: %v\n", err)
		return "", nil
	}

	imgBytes, err := base64.StdEncoding.DecodeString(respBase64.Data[0].B64JSON)
	if err != nil {
		fmt.Printf("Base64 decode error: %v\n", err)
		return "", nil
	}

	r := bytes.NewReader(imgBytes)
	imgData, err := png.Decode(r)
	if err != nil {
		fmt.Printf("PNG decode error: %v\n", err)
		return "", nil
	}
	imageFilePath := randomImageName()

	file, err := os.Create(imageFilePath)
	if err != nil {
		fmt.Printf("File creation error: %v\n", err)
		return "", nil
	}
	defer file.Close()

	if err := png.Encode(file, imgData); err != nil {
		fmt.Printf("PNG encode error: %v\n", err)
		return "", nil
	}

	fmt.Println("The image was saved as example.png")
	return imageFilePath, nil
}

func randomImageName() string {
	timestamp := time.Now().Format("02-01-2006-15-04-05")
	timestamp = strings.Replace(timestamp, "-", "", -1)
	timestamp = strings.Replace(timestamp, ":", "", -1)

	letterPool := "abcdefghijklmnopqrstuvwxyzABZDEFGHIJKLMNOPQRSTUVWXYZ"
	randomLetter := string(letterPool[rand.Intn(len(letterPool))])

	var letters strings.Builder
	for i := 0; i < 5; i++ {
		letters.WriteByte(letterPool[rand.Intn(len(letterPool))])
	}

	randomName := timestamp + randomLetter + letters.String()
	randomFileName := randomName + ".png"
	imageFileName := "dalleAssets/image_" + randomFileName
	if !strings.HasSuffix(imageFileName, ".png") {
		imageFileName += ".png"
	}
	return imageFileName
}
