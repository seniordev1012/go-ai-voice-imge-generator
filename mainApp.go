package main

import (
	"aigen/aigenUi"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

type Users struct {
	Id       int
	Email    string
	Password string
}

// Create the main app tabs
// The tabs are the main content of the app
// This is where All the fun happens :)
func mainApp(mapungubwe fyne.App) (*container.AppTabs, *container.Split) {
	//socialTabCon := aigenUi.SocialTab()
	//financeTab := aigenUi.FinancialTab(mapungubwe)
	//planPage := aigenUi.PersonalTab(mapungubwe)
	//newsTab := aigenUi.NewsTab()
	chat, aiGen := ChatTab()
	settingsTab := aigenUi.GenSettings(mapungubwe)
	extendAI := aigenUi.Extensions(mapungubwe)

	//Create the tabs container and add the tabs to it
	tabs := container.NewAppTabs(
		aiGen,
		//financeTab,
		//planPage,
		//newsTab,
		//socialTabCon,
		extendAI,
		settingsTab,
	)

	inputBoxContainer := SignInHandler(chat, tabs, aiGen)

	return tabs, inputBoxContainer
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
