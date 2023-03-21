package main

import (
	"aigen/aigenAudioAutoPlay"
	"aigen/aigenRecorder"
	"aigen/aigenRest"
	"aigen/textHandler"
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"strings"
)

// addChatBubble adds a chat bubble to the chat box
// The message is the text to be displayed
// The isUser parameter determines if the message is from the user or the bot
// If the message is from the user, the bubble will be on the right side of the chat window
func addChatBubble(box *fyne.Container, message string, isUser bool) {

	label := widget.NewLabel(textHandler.SeparateLines(message))
	//avatarImg, _ := chatAvatars()
	//Check for message in db
	//If message is in db, display the message
	audio, err := getAudio(message)
	if err != nil {
		log.Printf("Error: %s", err)
	}
	log.Printf("Audio: %s", audio)
	var messageCard *widget.Card

	if audio != "" {
		audioFile := widget.NewButtonWithIcon("Audio", theme.MediaPlayIcon(), func() {
			//Play audio
			log.Printf("Playing audio file %s", audio)

		})
		audioFile.BaseWidget.ExtendBaseWidget(audioFile)

		audioFile.OnTapped = func() {
			audioFile.SetIcon(nil)
			//TODO:Set counter to count seconds of audio playing
			audioFile.SetText("Playing")
			audioFile.Importance = widget.HighImportance
			audioFile.ExtendBaseWidget(audioFile)

			audioFile.Refresh()
			//Play audio
			defer func(filename string) {
				err := aigenAudioAutoPlay.PlayAudioPlayback(filename)
				if err != nil {
					log.Printf("Error: %s", err)
				}
				log.Printf("Audio file %s played successfully", filename)
				audioFile.SetIcon(theme.MediaPlayIcon())

			}(audio)

			log.Printf("Playing audio file %s", audio)
			return
		}

		messageCard = widget.NewCard("", "", label)
		messageCard.Content = container.NewHBox(audioFile, messageCard.Content)

	} else {
		messageCard = widget.NewCard("", "", label)
	}

	if isUser {
		box.Add(container.NewHBox(
			layout.NewSpacer(),
			messageCard,
			//avatarImg,
		))
	} else {
		box.Add(container.NewHBox(
			//botAvatarImg,
			messageCard,
			layout.NewSpacer(),
		))
	}
	container.NewScroll(box).SetMinSize(fyne.NewSize(100, 100))
}

func addMediaChatBubble(box *fyne.Container, message string, isUser bool) {
	label := widget.NewLabel("AiGen: Ai Generated Image......From Prompt")
	//avatarImg, _ := chatAvatars()
	//Check for message in db
	//If message is in db, display the message
	var messageCard *widget.Card
	log.Println("loading image from " + message)
	image := canvas.NewImageFromFile(message)
	image.FillMode = canvas.ImageFillStretch
	image.Move(fyne.NewPos(0, 0))
	//image.SetMinSize(fyne.NewSize(100, 100))
	image.Resize(fyne.NewSize(512, 512))
	image.Refresh()

	messageCard = widget.NewCard("", "", label)
	messageCard.Image = image
	messageCard.Size()
	messageCard.CreateRenderer()
	messageCard.Refresh()

	if isUser {

		box.Add(container.NewHBox(
			layout.NewSpacer(),
			messageCard,
			//avatarImg,
		))

	} else {

		box.Add(container.NewHBox(
			//botAvatarImg,
			messageCard,
			layout.NewSpacer(),
		))

	}
	container.NewScroll(box).SetMinSize(fyne.NewSize(100, 100))

}

func sendButton(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create a send button for sending messages
	sendButton := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {

	})

	sendButton.OnTapped = func() {
		sendNotificationNow("Message sent")
		sendButton.Importance = widget.HighImportance
		sendButton.Refresh()
		message := inputBox.Text
		//Separate each line with new line /n
		message = textHandler.SeparateLines(message)
		fmt.Println(message)
		//DISPLAY MESSAGE
		displayConvo(message, tab1, inputBox, "none")

	}

	return sendButton
}

func voiceChatButton(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create a voice chat button for sending voice messages
	voiceChatButton := widget.NewButtonWithIcon("PRESS TO TALK ", theme.MediaRecordIcon(), func() {

	})
	voiceChatButton.Importance = widget.HighImportance

	voiceChatButton.OnTapped = func() {
		sendNotificationNow("Voice chat started")
		voiceChatButton.Importance = widget.DangerImportance
		voiceChatButton.SetText("Recording...")
		voiceChatButton.SetIcon(theme.MediaStopIcon())

		recorder, err := aigenRecorder.VoiceRecorder()
		if recordingError(err) {
			return
		}

		//show count for seconds of recording

		log.Printf("Voice recorder __ started: %v", recorder)

		message := aigenRest.Whisper(recorder)

		message = textHandler.SeparateLines(message)

		fmt.Println(message)
		displayConvo(message, tab1, inputBox, recorder)
		voiceChatButton.Importance = widget.HighImportance
		voiceChatButton.SetText("PRESS TO TALK")
		voiceChatButton.SetIcon(theme.MediaRecordIcon())
		return
	}

	return voiceChatButton
}

func recordingError(err error) bool {
	if err != nil {
		log.Printf("Error starting voice recorder: %v", err)
		return true
	}
	return false
}

// Display Conversation with the bot
// displayConvo function to display messages from the user and the bot
func displayConvo(message string, tab1 *fyne.Container, inputBox *widget.Entry, mediaInputPath string) {
	if message != "" {
		userMessages(message, tab1)
		if mediaInputPath != "none" {

			addMessages := addMessageWithMedia("YOU", message, mediaInputPath, mediaInputPath)

			if addMessages != nil {
				log.Printf("Error adding user message: %v", addMessages)
			}

		} else {
			addMessages := addMessage("YOU", message)
			if addMessages != nil {
				log.Printf("Error adding user message: %v", addMessages)
			}
		}

		// Message Input Box
		inputBox.SetText("")
		//check if message contains words: Image, Generate or Generate Image, or each word separately and then call the function to generate an image

		if strings.Contains(message, "Image") || strings.Contains(message, "image") || strings.Contains(message, "photo") || strings.Contains(message, "Photo") || strings.Contains(message, "Generate") || strings.Contains(message, "Generate Image") {
			//generate image
			messageCall, err := aigenRest.ImageGenerationCall(message)
			limit := 120
			notificationMessage := message

			if len(notificationMessage) > limit {
				notificationMessage = notificationMessage[:limit]
			}

			sendNotificationNow("Image Generated Successfully For:" + notificationMessage)

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

		} else {
			messageCall, err := aigenRest.MakeApiCall(message)
			limit := 120
			notificationMessage := messageCall

			if len(notificationMessage) > limit {
				notificationMessage = notificationMessage[:limit]
			}
			sendNotificationNow(notificationMessage)

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
		// Switch API Provider

		//messageCall, err := ronSwan()

	}
}

// botMessages function to display messages from the bot
// This function is used to split messages into multiple chat bubbles if the message is too long
// This function is also used to send voice notes if the message is too long
func botMessages(messageCall string, err error, tab1 *fyne.Container, contentType string) {
	//Send voice note if message is more than 120 characters
	if contentType == "text" {
		if len(messageCall) > 0 {
			sendAudio, _ := pressPlayAudio(messageCall)
			if sendAudio != true {
				log.Printf("Error sending audio: %v", sendAudio)
			}
			addChatBubble(tab1, "Bot: "+messageCall, false)
		}
	}

	if contentType == "image" {
		if len(messageCall) > 0 {
			addMediaChatBubble(tab1, messageCall, false)
		}
	}

}

func userMessages(message string, tab1 *fyne.Container) {

	addChatBubble(tab1, "YOU: "+message, true)
}
