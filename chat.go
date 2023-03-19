package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

type Console struct {
	Content *fyne.Container
	Buttons []*widget.Button
	Focus   func()
	wait    chan string
	view    *fyne.Container
	row     int // current row
	rowMax  int
	entry   *widget.Entry
	prompt  string
	font    fyne.TextStyle
}

// addChatBubble adds a chat bubble to the chat box
// The message is the text to be displayed
// The isUser parameter determines if the message is from the user or the bot
// If the message is from the user, the bubble will be on the right side of the chat window
func addChatBubble(box *fyne.Container, message string, isUser bool) {

	label := widget.NewLabel(separateLines(message))
	//avatarImg, _ := chatAvatars()

	messageCard := widget.NewCard("", "", label)
	//bubbleCardContent(messageCard, avatarImg)

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
		message := inputBox.Text
		//Separate each line with new line /n
		message = separateLines(message)
		fmt.Println(message)
		displayConvo(message, tab1, inputBox)
	})
	return sendButton
}

func voiceChatButton(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create a voice chat button for sending voice messages
	voiceChatButton := widget.NewButtonWithIcon("", theme.MediaRecordIcon(), func() {
		// Start recording voice
		recorder, err := VoiceRecorder()
		if recordingError(err) {
			return
		}
		log.Printf("Voice recorder started: %v", recorder)
	})

	// Set the button to stop recording if held down
	voiceChatButton.ExtendBaseWidget(voiceChatButton)
	voiceChatButton.OnTapped = func() {
		// Start recording voice
		recorder, err := VoiceRecorder()
		if recordingError(err) {
			return
		}
		log.Printf("Voice recorder started: %v", recorder)
		return
	}

	//voiceChatButton.OnPointerUp = func(event *fyne.PointEvent) {
	//	// Stop recording voice
	//	StopVoiceRecorder()
	//}

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
func displayConvo(message string, tab1 *fyne.Container, inputBox *widget.Entry) {
	if message != "" {
		userMessages(message, tab1)
		addMessages := addMessage("YOU", message)

		if addMessages != nil {
			log.Printf("Error adding user message: %v", addMessages)
		}
		// Clear input box
		inputBox.SetText("")
		//TODO: Make API call to get response from bot
		messageCall, err := makeApiCall(message)
		//messageCall, err := ronSwan()
		log.Printf("Message call: %v", messageCall)
		if err != nil {
			log.Printf("Error making API call: %v", err)
		}

		botMessages(messageCall, err, tab1)
		addBotMessages := addMessage("Bot", messageCall)

		if addBotMessages != nil {
			log.Printf("Error adding bot message: %v", addBotMessages)
		} else {

			log.Printf("Bot message added successfully")
		}
	}
}

// botMessages function to display messages from the bot
// This function is used to split messages into multiple chat bubbles if the message is too long
// This function is also used to send voice notes if the message is too long
func botMessages(messageCall string, err error, tab1 *fyne.Container) {
	//Send voice note if message is more than 120 characters
	if len(messageCall) > 3 {
		sendAudio, audioFilePathFinder := pressPlayAudio(messageCall)
		if sendAudio != true {
			log.Printf("Error sending audio: %v", sendAudio)
		}

		messageCall = "audio: " + audioFilePathFinder + messageCall
		addChatBubble(tab1, "Bot: "+messageCall, false)
	}
}

func userMessages(message string, tab1 *fyne.Container) {

	addChatBubble(tab1, "YOU: "+message, true)
}
