package aigenRest

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"github.com/sashabaranov/go-openai"
	"image/png"
	"os"
)

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

	fmt.Println("Image saved to", imageFilePath)
	return imageFilePath, nil
}
