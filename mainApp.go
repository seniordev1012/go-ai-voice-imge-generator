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
	chat, aiGen := ChatTab()
	settingsTab := aigenUi.GenSettings(mapungubwe)
	extendAI := aigenUi.Extensions(mapungubwe)
	chatMediaTab := aigenUi.UserMedia()
	developerMode := aigenUi.Developer()

	//Create the tabs container and add the tabs to it
	tabs := container.NewAppTabs(aiGen, //financeTab,
		extendAI, chatMediaTab, developerMode, settingsTab)

	inputBoxContainer := SignInHandler(chat, tabs, aiGen)

	return tabs, inputBoxContainer
}
