package main

import (
	"aigen/aigenUi"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// Create the main app tabs
// The tabs are the main content of the app
// This is where All the fun happens :)
func mainApp(mapungubwe fyne.App) (*container.AppTabs, *container.Split) {
	chat, aiGen := ChatTab()
	settingsTab := aigenUi.GenSettings(mapungubwe)
	extendAI := aigenUi.Extensions(mapungubwe)
	chatMediaTab := aigenUi.UserMedia()
	//developerMode := aigenUi.Developer()
	audioSettingsTab := aigenUi.AudioSettingsTab()
	DualVoiceSettings := aigenUi.DualVoiceSettings()
	//Create the tabs container and add the tabs to it
	tabs := container.NewAppTabs(aiGen, //financeTab,
		extendAI, chatMediaTab, settingsTab, audioSettingsTab, DualVoiceSettings)

	inputBoxContainer := SignInHandler(chat, tabs, aiGen)

	return tabs, inputBoxContainer
}
