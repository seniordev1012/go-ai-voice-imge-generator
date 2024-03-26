package aigenRest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
)

func CallClaude(message string) (string, error) {

	url := "https://api.anthropic.com/v1/complete"
	method := "POST"

	payload := fmt.Sprintf(`{
        "model": "claude-2.1",
        "max_tokens_to_sample": 1024,
        "prompt": "\n\nHuman: %s\n\nAssistant:"
    }`, message)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return "", err
	}
	req.Header.Add("x-api-key", os.Getenv("ANTHROPIC"))
	req.Header.Add("anthropic-version", "2023-06-01")
	req.Header.Add("content-type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	defer res.Body.Close()

	var response map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	completion, ok := response["completion"].(string)
	if !ok {
		return "", fmt.Errorf("completion not found in response")
	}

	return completion, nil
}
