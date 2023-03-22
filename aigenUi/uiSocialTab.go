package aigenUi

import (
	_ "aigen/aigeUi"
	"aigen/aigenRest"
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

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

// SocialTab is the tab that contains social apps
func SocialTab() *container.TabItem {
	var tweetResponse TweetResponse
	getTweets := aigenRest.TwitterHome()
	err := json.Unmarshal(getTweets, &tweetResponse)
	if err != nil {
		panic(err)
	}

	// Loop through tweets and create a card for each tweet
	var tweetCards []*widget.Card

	for _, tweet := range tweetResponse.Data {
		cardContent := widget.NewCard("", "", widget.NewLabel(tweet.Text))
		tweetCard := widget.NewCard("", "", cardContent)
		tweetCard.Content = cardContent
		tweetCard.SetTitle(tweet.ID)
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		cardContent.ExtendBaseWidget(cardContent)
		tweetCards = append(tweetCards, tweetCard)
		tweetCards = append(tweetCards, widget.NewCard("", "", widget.NewLabel("")))
		tweetCard = fyne.CanvasObject(tweetCard).(*widget.Card)
		tweetCard.MinSize()
		tweetCard.Resize(fyne.NewSize(100, 200))
	}

	inputBox := widget.NewMultiLineEntry()
	inputBox.Wrapping = fyne.TextWrapWord
	inputBox.OnChanged = func(text string) {
		log.Printf("Text changed to: %s", text)
	}

	//Button to post to twitter
	postToTwitterButton := widget.NewButton("Post To Twitter", func() {
		aigenRest.SendTweet(inputBox.Text)
		inputBox.SetText("")
		aigenRest.SendNotificationNow("Tweet Sent Successfully")
	})
	postToTwitterButton.MinSize()
	postToTwitterButton.Resize(fyne.NewSize(100, 200))
	postToTwitterButton.ExtendBaseWidget(postToTwitterButton)

	// Post To Twitter Card
	postToTwitterCard := widget.NewCard("Post To Twitter", "", inputBox)
	page := container.NewVBox(postToTwitterCard, postToTwitterButton, tweetCards[0])

	//Facebook REST API Response
	var facebookResponse facebookFeed
	getFacebookFeed := aigenRest.FacebookPosts()
	err = json.Unmarshal(getFacebookFeed, &facebookResponse)
	if err != nil {
		panic(err)
	}

	// Loop through facebook posts and create a card for each post
	var facebookCards []*widget.Card

	for _, facebookPost := range facebookResponse.Data {
		cardContent := widget.NewCard("", "", widget.NewLabel(facebookPost.Message))
		facebookCard := widget.NewCard("", "", cardContent)
		facebookCard.Content = cardContent
		facebookCard.SetTitle(facebookPost.Id)
		fyne.CurrentApp().Settings().SetTheme(theme.LightTheme())
		cardContent.ExtendBaseWidget(cardContent)
		facebookCards = append(facebookCards, facebookCard)
		facebookCards = append(facebookCards, widget.NewCard("", "", widget.NewLabel("")))
		facebookCard = fyne.CanvasObject(facebookCard).(*widget.Card)
		facebookCard.MinSize()
		facebookCard.Resize(fyne.NewSize(100, 200))
	}

	//Button to post to facebook
	postToFacebookButton := widget.NewButton("Post To Facebook", func() {
		//aigenRest.SendFacebookPost(inputBox.Text)
		inputBox.SetText("")
		aigenRest.SendNotificationNow("Facebook Post Sent Successfully")

	})
	postToFacebookButton.MinSize()
	postToFacebookButton.Resize(fyne.NewSize(100, 200))
	postToFacebookButton.ExtendBaseWidget(postToFacebookButton)

	// Post To Facebook Card
	postToFacebookCard := widget.NewCard("Post To Facebook", "", inputBox)
	facebookPages := container.NewVBox(postToFacebookCard, postToFacebookButton, facebookCards[0])

	socialTabCon := container.NewTabItem("Social", container.NewAppTabs(
		container.NewTabItem("Twitter", page),
		container.NewTabItem("Facebook", facebookPages),
		//container.NewTabItem("Discord", widget.NewAccordion()),
		//container.NewTabItem("Telegram", widget.NewAccordion()),
		//container.NewTabItem("WhatsApp", widget.NewAccordion()),
	))
	socialTabCon.Icon = theme.GridIcon()

	return socialTabCon
}
