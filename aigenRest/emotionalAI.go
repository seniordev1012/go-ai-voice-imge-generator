package aigenRest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
)

func emotionalAI(text string) (string, error) {
	data := struct {
		Model            string  `json:"model"`
		Prompt           string  `json:"prompt"`
		Temperature      float32 `json:"temperature"`
		MaxTokens        int     `json:"max_tokens"`
		TopP             float32 `json:"top_p"`
		FrequencyPenalty float32 `json:"frequency_penalty"`
		PresencePenalty  float32 `json:"presence_penalty"`
	}{
		Model:            "text-davinci-003",
		Prompt:           "Classify the sentiment in these text:\nBased On These Emotions Determining Voice output:\nDefault\nChat\nAngry\nCheerful\nExcited\nFriendly\nHopeful\nSad\nShouting\nTerrified\nUnfriendly\nWhispering\n\"" + text + "\"\n\nEmotion voice to use:\n",
		Temperature:      0,
		MaxTokens:        60,
		TopP:             1,
		FrequencyPenalty: 0,
		PresencePenalty:  0,
	}

	payload, err := json.Marshal(data)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(payload))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")
	apiKey := os.Getenv("OPENAI")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var result struct {
		Choices []struct {
			Text string `json:"text"`
		} `json:"choices"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return "", err
	}

	if len(result.Choices) > 0 {
		return result.Choices[0].Text, nil
	}

	return "", nil
}
