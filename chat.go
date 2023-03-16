package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
	"os"
	"time"
)

// Create a new label with the message and add it to the chat window.
// The isUser parameter determines if the message is from the user or the bot
// If the message is from the user, the bubble will be on the right side of the chat window
// If the message is from the bot, the bubble will be on the left side of the chat window
// Create a new label with the message and add it to the chat window
// The isUser parameter determines if the message is from the user or the bot
// If the message is from the user, the bubble will be on the right side of the chat window
// If the message is from the bot, the bubble will be on the left side of the chat window
func addChatBubble(box *fyne.Container, message string, isUser bool) {

	// Create a new label with the message
	label := widget.NewLabel(message)
	label.Resize(fyne.NewSize(100, 0))
	label.TextStyle = fyne.TextStyle{Bold: false, Italic: false, Monospace: false}
	//Add image card
	image := canvas.NewImageFromFile("source/avatar.jpg")
	image.SetMinSize(fyne.NewSize(100, 100))
	imageCard := widget.NewCard("", "", image)
	imageCard.Resize(fyne.NewSize(100, 100))
	//Add image card

	// Create a new chat bubble with the label

	bubble := container.NewHBox(label)
	bubble.Layout = layout.NewVBoxLayout()

	container.NewScroll(bubble)

	// Create a new image widget with the avatar URL
	avatarImg := canvas.NewImageFromFile("source/avatar.jpg")
	avatarImg.SetMinSize(fyne.NewSize(64, 64))
	avatarImg.Resize(fyne.NewSize(64, 64))

	botAvatarImg := canvas.NewImageFromFile("source/botAvatar.png")
	botAvatarImg.SetMinSize(fyne.NewSize(64, 64))
	botAvatarImg.Move(fyne.NewPos(-5, -5))

	// Add the chat bubble to the card
	if isUser {
		// If the message is from the user, add the bubble to the right side of the card
		box.Add(container.NewHBox(
			layout.NewSpacer(),
			widget.NewCard("", "", bubble),
			avatarImg,
		))
	} else {
		// If the message is from someone else, add the bubble to the left side of the card
		box.Add(container.NewHBox(
			botAvatarImg,
			widget.NewCard("", "", bubble),
			layout.NewSpacer(),
		))
	}
	container.NewScroll(box)
}

func sendButton(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create a send button for sending messages
	sendButton := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {
		message := inputBox.Text
		//Increase width of input box
		fmt.Println(message)
		//Display Conversation with the bot
		displayConvo(message, tab1, inputBox)
	})
	return sendButton
}

func voiceChatButton(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create a voice chat button for sending voice messages
	voiceChatButton := widget.NewButtonWithIcon("", theme.MediaRecordIcon(), func() {
		// Start recording voice
		VoiceRecorder()

		// Wait for 15 seconds before stopping the recording
		go func() {
			time.Sleep(15 * time.Second)
			// Stop recording voice
			//StopVoiceRecorder()
			os.Exit(0)
		}()
	})

	// Set the button to stop recording if held down
	voiceChatButton.ExtendBaseWidget(voiceChatButton)
	voiceChatButton.OnTapped = func() {
		// Start recording voice
		VoiceRecorder()
	}

	//voiceChatButton.OnPointerUp = func(event *fyne.PointEvent) {
	//	// Stop recording voice
	//	StopVoiceRecorder()
	//}

	return voiceChatButton
}

func StopVoiceRecorder() {
	// Stop recording voice
	//StopVoiceRecorder()
	os.Exit(0)
}

func displayConvo(message string, tab1 *fyne.Container, inputBox *widget.Entry) {
	if message != "" {

		userMessages(message, tab1)
		addMessages := addMessage("YOU", message)
		if addMessages != nil {
			log.Printf("Error adding user message: %v", addMessages)
		}

		// Clear input box
		inputBox.SetText("")

		messageCall, err := makeApiCall(message)
		if err != nil {
			log.Printf("Error making API call: %v", err)
		}

		botMessages(messageCall, err, tab1)
		addBotMessages := addMessage("Bot", messageCall)
		if addBotMessages != nil {
			log.Printf("Error adding bot message: %v", addBotMessages)

		}

	}
}

// botMessages function to display messages from the bot
// This function is used to split messages into multiple chat bubbles if the message is too long
// This function is also used to send voice notes if the message is too long
func botMessages(messageCall string, err error, tab1 *fyne.Container) {
	//Send voice note if message is more than 120 characters
	if len(messageCall) > 3 {
		playAiVoice := voiceNote(messageCall, err)
		if playAiVoice != nil {
			log.Printf("Error playing voice note: %v", playAiVoice)
		}
	}
	addChatBubble(tab1, "Bot: "+messageCall, false)

}

func userMessages(message string, tab1 *fyne.Container) {

	//if len(message) > 100 {
	//	voiceNote(message, nil)
	//}

	addChatBubble(tab1, "YOU: "+message, true)
}
