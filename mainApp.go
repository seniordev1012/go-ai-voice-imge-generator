package main

import (
	"aigen/aigenUi"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

// mainApp is the main function that creates the tabs and the input box container
// It returns the tabs and the input box container
// The input box container is the container that contains the input box, send button, and voice chat button
// This is where All the fun happens :)
func mainApp(mapungubwe fyne.App) (*container.AppTabs, *container.Split) {
	chat, aiGen := ChatTab()
	settingsTab := aigenUi.GenSettings(mapungubwe)
	// extendAI := aigenUi.Extensions(mapungubwe)
	chatMediaTab := aigenUi.UserMedia()
	//developerMode := aigenUi.Developer()
	audioSettingsTab := aigenUi.AudioSettingsTab()
	multiModels := aigenUi.MultiModelSettings(mapungubwe)
	//Create the tabs container and add the tabs to it
	//removed extendAI which was meant for extensions, may be restored.
	tabs := container.NewAppTabs(aiGen, //financeTab,
		multiModels, chatMediaTab, settingsTab, audioSettingsTab)

	inputBoxContainer := SignInHandler(chat, tabs, aiGen)

	return tabs, inputBoxContainer
}
