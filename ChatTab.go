package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
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
	chat := container.NewVBox()
	aiGen := container.NewTabItem("AiGen-Chat", chat)

	messages1, err := getMessages()

	if err != nil {
		log.Printf("Error getting messages: %v", err)
	}

	for _, message := range messages1 {
		addChatBubble(chat, message.Sender+": "+message.Content, message.Sender == "Bot")
	}
	startUpCall(chat) //This is a function that is called when the chat tab is opened
	return chat, aiGen
}
