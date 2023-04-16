package main

import (
	"aigen/aigenRest"
	"fyne.io/fyne/v2"
	"log"
)

// Default logic to handle communication with GPT
func defaultCallConverseLogic(message string, tab1 *fyne.Container) {

	//message = longTermMemory(message)

	messageCall, err := aigenRest.MakeApiCall(message)
	limit := 120
	notificationMessage := messageCall

	if len(notificationMessage) > limit {
		notificationMessage = notificationMessage[:limit]
	}
	aigenRest.SendNotificationNow(notificationMessage)

	log.Printf("Message call: %v", messageCall)

	if err != nil {
		log.Printf("Error making API call: %v", err)
	}

	botMessages(messageCall, err, tab1, "text")

	addBotMessages := addMessage("Bot", messageCall)
	if addBotMessages != nil {
		log.Printf("Error adding bot message: %v", addBotMessages)
	}
}

// Long Term Memory Logic to add last conversation to GPT
func longTermMemory(message string) string {
	lastConvo, lastMessagesSuccess := getLastMessages()

	if lastMessagesSuccess != nil {
		log.Println(lastMessagesSuccess)
	}

	//Loop Through Messages From DB and Append
	for _, lastConversation := range lastConvo {
		//Append text and feed to GPT for long term memory
		message = message + " " + lastConversation.Content
		log.Printf("Message: %v", message)
	}
	//Add Prompt to GPT before the lastConversation
	message = "This is a conversation betwee  us, you are Bot and I am you. " +
		"sometimes don't even use this last conversation, just use the last message. , it is for context. " +
		"" + message
	return message
}

// Display Image and Generate Image Logic
func imageGenerationLogic(message string, tab1 *fyne.Container) {
	//generate image
	messageCall, err := aigenRest.ImageGenerationCall(message)
	limit := 120
	notificationMessage := message

	if len(notificationMessage) > limit {
		notificationMessage = notificationMessage[:limit]
	}

	aigenRest.SendNotificationNow("Image Generated Successfully For:" + notificationMessage)

	log.Printf("Message call: %v", messageCall)
	if err != nil {
		log.Printf("Error making API call: %v", err)
	}
	botMessages(messageCall, err, tab1, "image")

	//TODO: Add image to database properly

	addBotMessages := addMessageWithMedia("Bot", message, "none", messageCall)
	if addBotMessages != nil {
		log.Printf("Error adding bot message: %v", addBotMessages)
	}
}

// Handle Twitter Push Logic
func twitterPushLogic(message string, tab1 *fyne.Container) {
	messageCall, err := aigenRest.MakeApiCall(message)
	limit := 280
	sendToTwitter, sentSuccess := aigenRest.SendTweet(messageCall)
	if sentSuccess != nil {
		log.Println(sentSuccess)
		log.Println(sendToTwitter)
	}

	notificationMessage := "Sent Tweet:" + messageCall

	if len(notificationMessage) > limit {
		notificationMessage = notificationMessage[:limit]
	}
	aigenRest.SendNotificationNow(notificationMessage)
	log.Printf("Message call: %v", messageCall)
	if err != nil {
		log.Printf("Error making API call: %v", err)
	}

	botMessages("Tweet Sent:"+messageCall, err, tab1, "text")

	addBotMessages := addMessage("Bot", messageCall)
	if addBotMessages != nil {
		log.Printf("Error adding bot message: %v", addBotMessages)
	}
}
