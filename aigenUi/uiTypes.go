package aigenUi

type Tweet struct {
	EditHistoryTweetIDs []string `json:"edit_history_tweet_ids"`
	ID                  string   `json:"id"`
	Text                string   `json:"text"`
}

type TweetResponse struct {
	Data []Tweet `json:"data"`
}
type facebookFeed struct {
	Data []struct {
		CreatedTime string `json:"created_time"`
		Id          string `json:"id"`
		Message     string `json:"message,omitempty"`
	} `json:"data"`
	Paging struct {
		Previous string `json:"previous"`
		Next     string `json:"next"`
	} `json:"paging"`
}
