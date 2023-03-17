package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"log"
)

// ChatTab Chat Bubble is a container that contains a label and an avatar
// The label contains the message
// The avatar contains the avatar of the sender
// The isUser parameter determines if the message is from the user or the bot
// If the message is from the user, the bubble will be on the right side of the chat window
// If the message is from the bot, the bubble will be on the left side of the chat window
func ChatTab() (*fyne.Container, *container.TabItem) {
	//Create the chat tab
	chat := container.NewHBox()
	container.NewAdaptiveGrid(2, chat)
	chat.Layout = layout.NewVBoxLayout()
	aiGen := container.NewTabItem("AiGen-Chat", chat)

	messages1, err := getMessages()

	if err != nil {
		log.Printf("Error getting messages: %v", err)
	}

	for _, message := range messages1 {
		if message.Sender == "YOU" {
			addChatBubble(chat, message.Content, true)
		} else {
			addChatBubble(chat, message.Content, false)
		}
	}
	startUpCall(chat) //This is a function that is called when the chat tab is opened
	return chat, aiGen
}