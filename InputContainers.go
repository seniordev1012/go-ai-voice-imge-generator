package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func bottomInputBox(chat *fyne.Container, tabs *container.AppTabs, aiGen *container.TabItem) *container.Split {
	inputBox := widget.NewMultiLineEntry()
	inputBox.Wrapping = fyne.TextWrapWord
	//TODO: Edit InputBox OnChanged
	inputBox.OnChanged = func(s string) {
		log.Printf("Input changed to: %s", s)
		kitchenLog(s)
	}
	//TODO: Edit InputBox Container
	inputBox.PlaceHolder = "Enter your message here..."
	sendButton := sendButton(inputBox, chat)
	voiceNoteButton := voiceChatButton(inputBox, chat)
	mediaUploadButton := imageUploadInput(inputBox, chat)
	voiceNoteButton.Resize(fyne.NewSize(50, 100))
	inputBoxContainer := container.NewHSplit(inputBox, sendButton)
	inputBoxContainer = container.NewVSplit(inputBoxContainer, voiceNoteButton)
	inputBoxContainer = container.NewVSplit(inputBoxContainer, mediaUploadButton)
	tabs.OnSelected = func(tab *container.TabItem) {
		if tab == aiGen {
			//Show if tab is home
			inputBoxContainer.Show()
		} else {
			inputBoxContainer.Hide()
		}
	}
	return inputBoxContainer
}
