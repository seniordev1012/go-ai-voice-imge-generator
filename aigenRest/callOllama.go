package aigenRest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func CallOllama(message string) (string, error) {

	url := "http://localhost:11434/api/generate"
	method := "POST"

	payload := fmt.Sprintf(`{
        "model": "ollama-2.1",
		"stream": false,
        "prompt": "\n\nHuman: %s\n\nAssistant:"
    }`, message)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, strings.NewReader(payload))

	if err != nil {
		fmt.Println(err)
		return "", err
	}

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
