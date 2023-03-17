package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

// Create a new input box and a send button
// The input box is a multi-line entry
// The send button is a button with an icon
// The send button will send the message in the input box to the chat window
// The send button will also clear the input box
func bottomInputBox(chat *fyne.Container, tabs *container.AppTabs, aiGen *container.TabItem) *container.Split {
	inputBox := widget.NewMultiLineEntry()
	inputBox.Wrapping = fyne.TextWrapWord
	inputBox.OnChanged = func(s string) {
		log.Printf("Input changed to: %s", s)
		kitchenLog(s)
	}
	//TODO: Edit InputBox Container
	inputBox.PlaceHolder = "Enter your message here..."
	sendButton := sendButton(inputBox, chat)
	voiceNoteButton := voiceChatButton(inputBox, chat)
	inputBoxContainer := container.NewVSplit(inputBox, sendButton)
	inputBoxContainer = container.NewHSplit(inputBoxContainer, voiceNoteButton)
	//Hide the input box container if it is not chat tab
	//Use `AppTabs.Selected() *TabItem` instead.
	tabs.OnSelected = func(tab *container.TabItem) {
		if tab == aiGen {
			inputBoxContainer.Show()
		} else {
			inputBoxContainer.Hide()
		}
	}
	return inputBoxContainer
}
