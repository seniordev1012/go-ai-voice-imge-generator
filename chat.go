package main

import (
	"database/sql"
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

		messageCall, err := makeApiCall()
		if err != nil {
			log.Printf("Error making API call: %v", err)
		}

		botMessages(messageCall, err, tab1)
		addMessage := addMessage("Bot", messageCall)
		if addMessage != nil {
			log.Printf("Error adding bot message: %v", addMessage)

		}

	}
}

// botMessages function to display messages from the bot
// This function is used to split messages into multiple chat bubbles if the message is too long
// This function is also used to send voice notes if the message is too long
func botMessages(messageCall string, err error, tab1 *fyne.Container) {
	if len(messageCall) > 60 {
		//Send voice note if message is more than 120 characters
		if len(messageCall) > 90 {
			playAiVoice := voiceNote(messageCall, err)
			if playAiVoice != nil {
				log.Printf("Error playing voice note: %v", playAiVoice)
			}
		}

		log.Printf("Message is too long. Splitting message into multiple chat bubbles")

		var messageArray []string
		//Split message into multiple chat bubbles
		for i := 0; i < len(messageCall); i += 60 {
			end := i + 60
			//If end is greater than the length of the message,
			//set end to the length of the message
			if end > len(messageCall) {
				end = len(messageCall)
			}
			//Append message to messageArray
			messageArray = append(messageArray, messageCall[i:end])

		}
		//Add chat bubbles to the chat window
		for _, message := range messageArray {
			addChatBubble(tab1, "Bot: "+message, true)
		}
	} else {
		addChatBubble(tab1, "Bot: "+messageCall, true)
	}
}

func userMessages(message string, tab1 *fyne.Container) {

	if len(message) > 2 {
		voiceNote(message, nil)
	}

	var messageArray []string
	//Split message into multiple chat bubbles
	for i := 0; i < len(message); i += 60 {
		end := i + 60
		//If end is greater than the length of the message,
		//set end to the length of the message
		if end > len(message) {
			end = len(message)
		}
		//Append message to messageArray
		messageArray = append(messageArray, message[i:end])
	}
	for _, message := range messageArray {
		if len(message) > 60 {
			//Send voice note if message is more than 120 characters

			log.Printf("Message is too long. Splitting message into multiple chat bubbles")

			var messageArray []string
			//Split message into multiple chat bubbles
			for i := 0; i < len(message); i += 60 {
				end := i + 60
				//If end is greater than the length of the message,
				//set end to the length of the message
				if end > len(message) {
					end = len(message)
				}
				//Append message to messageArray
				messageArray = append(messageArray, message[i:end])

			}
			//Add chat bubbles to the chat window
			for _, message := range messageArray {
				addChatBubble(tab1, "Bot: "+message, true)
			}
		} else {
			addChatBubble(tab1, "You: "+message, true)
		}
	}
}

// addMessage adds a message to the database
func addMessage(sender string, content string) error {
	// Open a connection to the database
	db, err := sql.Open("sqlite3", "DB/messages.db")
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Printf("Error closing database: %v", err)
		}
	}(db)

	// Prepare a SQL statement to insert the message into the database
	stmt, err := db.Prepare("INSERT INTO messages (sender, content) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			log.Printf("Error closing statement: %v", err)
		}
	}(stmt)

	// Execute the prepared statement with the message as parameters
	_, err = stmt.Exec(sender, content)
	if err != nil {
		return err
	}

	return nil
}
