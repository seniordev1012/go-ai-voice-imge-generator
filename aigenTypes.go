package main

import (
	"time"
)

// Message represents a message in the chat
type Message struct {
	ID        int
	Sender    string
	Media     string
	Content   string
	CreatedAt time.Time
}

type User struct {
	ID       int
	Username string
	Password string
	Token    string
}

type twitterContent struct {
	Data []struct {
		EditHistoryTweetIds []string `json:"edit_history_tweet_ids"`
		Id                  string   `json:"id"`
		Text                string   `json:"text"`
		Withheld            struct {
			Copyright    bool     `json:"copyright"`
			CountryCodes []string `json:"country_codes"`
		} `json:"withheld,omitempty"`
	} `json:"data"`
	Meta struct {
		NextToken   string `json:"next_token"`
		ResultCount int    `json:"result_count"`
		NewestId    string `json:"newest_id"`
		OldestId    string `json:"oldest_id"`
	} `json:"meta"`
}

type Users struct {
	Id       int
	Email    string
	Password string
}
