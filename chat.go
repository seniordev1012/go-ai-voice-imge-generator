package main

import (
	"aigen/aigenAudioAutoPlay"
	"aigen/aigenRecorder"
	"aigen/aigenRest"
	"aigen/aigenUi"
	"aigen/textHandler"
	"fmt"
	"log"
	"os"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/atotto/clipboard"
)

const WeatherReport = "\nHere's a detailed weather breakdown for '  \"Location name\": \"\",':\n" +
	"- Temperature: 'use Degree Celsius'\n" +
	"- Feels Like: Degree\n" +
	"- Description: \n" +
	"- Humidity: %\n" +
	"- Wind Speed:  km/h\n" +
	"- Cloud Coverage: %\n\n" +
	"Based on this weather data, here are some recommendations:\n\n" +
	". **What to Wear**: , \n\n" +
	". **Things You Can Do**:\n\n" +
	"**How to Stay Productive**: \n\n" +
	"**Amount of Water to Stay Hydrated**:" +
	"**Funny Joke about current Weather**: \" \"\n\n"

// addChatBubble adds a chat bubble to the chat box
// The message is the text to be displayed
// The isUser parameter determines if the message is from the user or the bot
// If the message is from the user, the bubble will be on the right side of the chat window
func addChatBubble(box *fyne.Container, message string, isUser bool) {
	label := widget.NewButton(textHandler.SeparateLines(message), saveForLater(message))
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
		box.Add(container.NewHBox(layout.NewSpacer(), messageCard)) //avatarImg,

	} else {
		box.Add(container.NewHBox(
			//botAvatarImg,
			messageCard, layout.NewSpacer()))
	}
	container.NewScroll(box).SetMinSize(fyne.NewSize(100, 100))
}

func saveForLater(message string) func() {
	return func() {
		log.Printf("content to copy to clipboard is %s", message)
		err := clipboard.WriteAll(message)
		if err != nil {
			log.Println(err)
		}
		aigenRest.SendNotificationNow("Content Copied To Clipboard")
	}
}

// displayConvo displays the conversation in the chat box
// The message is the text to be displayed
func addMediaChatBubble(box *fyne.Container, message string, isUser bool) {
	label := widget.NewButton("Sage: AI Generated Image......From Prompt", nil)
	//avatarImg, _ := chatAvatars()
	//Check for message in db
	//If message is in db, display the message
	label.OnTapped = func() {
		saveForLater(message)
		newImageViewWindow(message)

	}
	var messageCard *widget.Card
	image := canvas.NewImageFromFile(message)
	image.FillMode = canvas.ImageFillStretch
	image.Move(fyne.NewPos(0, 0))
	image.Resize(fyne.NewSize(512, 512))
	image.Refresh()

	messageCard = widget.NewCard("", "", label)
	messageCard.Image = image
	messageCard.SetImage(image)
	messageCard.Size()
	messageCard.CreateRenderer()
	messageCard.Refresh()

	if isUser {
		box.Add(container.NewHBox(layout.NewSpacer(), messageCard)) //avatarImg,
	} else {
		box.Add(container.NewHBox(
			//botAvatarImg,
			messageCard, layout.NewSpacer()))
	}
	container.NewScroll(box).SetMinSize(fyne.NewSize(100, 100))
}

func newImageViewWindow(message string) {
	myApp := fyne.CurrentApp()
	imageDialog := myApp.NewWindow("View Image")

	// Load your image into an image object
	image := canvas.NewImageFromFile(message)

	// Create a content container for the image
	content := widget.NewCard(message, message, image)
	image.FillMode = canvas.ImageFillStretch
	imageDialog.Resize(fyne.NewSize(1200, 1200))
	imageDialog.CenterOnScreen()
	imageDialog.Icon()
	imageDialog.SetTitle("AI Generated Image")
	content.Resize(fyne.NewSize(960, 540))
	imageDialog.Show()
	imageDialog.FullScreen()
	// Set the window content to the image container
	imageDialog.SetContent(content)
}

func sendButton(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create a send button for sending messages
	sendButton := widget.NewButtonWithIcon("", theme.MailSendIcon(), func() {

	})

	sendButton.OnTapped = func() {
		messageNotificationSent()
		sendButton.Importance = widget.HighImportance
		sendButton.Refresh()
		message := inputBox.Text
		//Separate each line with new line /n
		message = textHandler.SeparateLines(message)
		//TODO::Remove Logger
		fmt.Println(message)
		//DISPLAY MESSAGE
		displayConvo(message, tab1, inputBox, "none")

	}

	return sendButton
}

func messageNotificationSent() {
	aigenRest.SendNotificationNow("Message sent")
}

func voiceChatButton(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create a voice chat button for sending voice messages
	voiceChatButton := widget.NewButtonWithIcon("PRESS TO TALK ", theme.MediaRecordIcon(), func() {
		//Fill in body for this part right here
	})
	voiceChatButton.Importance = widget.HighImportance

	voiceChatButton.OnTapped = func() {
		aigenRest.SendNotificationNow("Voice chat started, 20 seconds for message")
		voiceChatButton.Importance = widget.DangerImportance

		voiceChatButton.SetText("Recording...Message")
		voiceChatButton.SetIcon(theme.MediaStopIcon())

		recorder, err := aigenRecorder.VoiceRecorder()
		if recordingError(err) {
			return
		}
		//show count for seconds of recording
		log.Printf("Voice recorder __ started: %v", recorder)
		//Catch Words Spoken Through Whisper
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

/*
Upload image to platform , this is an attempt to have the ability
to send images to a model and have data brought back to the chat ui and have user interact with it
*/
func imageUploadInput(inputBox *widget.Entry, tab1 *fyne.Container) *widget.Button {
	// Create an "Upload Image" button
	uploadImageButton := widget.NewButtonWithIcon("Upload Image", theme.MailAttachmentIcon(), func() {
		// Implement the image upload functionality here
		filePath := OpenFile()
		if filePath != "" {
			// You can perform further actions here, like displaying the image or saving the file path
			// For example:
			// image := canvas.NewImageFromFile(filePath)
			// tab1.Add(widget.NewLabel("Image path: " + filePath))
			// tab1.Add(container.NewCenter(image))
		}
	})

	uploadImageButton.Importance = widget.WarningImportance
	uploadImageButton.Alignment = widget.ButtonAlignCenter

	uploadImageButton.Refresh()
	uploadImageButton.Show()
	uploadImageButton.OnTapped = func() {
		// Add any additional actions you want to perform when the button is tapped
		inputBox.Keyboard()
		log.Println("Upload Image button was pressed. Checking what is happening...")
	}
	inputBox.Show()

	return uploadImageButton
}

func OpenFile() string {
	fileDialog := widget.NewEntry()
	fileDialog.MultiLine = false
	fileDialog.Text = "Select an image file..."
	fileDialog.OnChanged = func(text string) {
		// Placeholder text, no action needed
	}
	widget.NewForm(widget.NewFormItem("Select Image", fileDialog))

	// Wait for file selection
	filePath := fileDialog.Text
	if filePath == "Select an image file..." {
		return ""
	}
	return filePath
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
		//TODO:Make AI Smarter
		//APIS!
		//TODO::Calendar
		//TODO::Email
		//TODO::Web Search
		//TODO:: Documents (PDF, CSV, TXT)
		//TODO:: Top News.
		//TODO:: Online Purchase
		//TODO:: Long Term Memory
		if triggerPublishToX(message) {
			log.Printf("Trying to send tweet to twitter")
			twitterPushLogic(message, tab1)
		} else if triggerImageGeneration(message) {
			imageGenerationLogic(message, tab1)
		} else if triggerWeatherInfo(message) {
			weatherDetailsFineTune := fineTuneWeather()
			defaultCallConverseLogic(weatherDetailsFineTune, tab1)
		} else if triggerSettings(message) {
			//UNCOMMENT THIS IF YOU WANT TO ALLOW THE MODEL TO RUN SYSTEM COMMANDS, TAKE CAUTION :)
			//} else if triggerProgramRun(message) {
			//	message = strings.Replace(message, "run program", "", -1)
			//	RunProgram(message)
		} else {
			defaultCallConverseLogic(message, tab1)
		}
	}
}

func triggerSettings(message string) bool {
	if strings.Contains(message, "triggerSettings") {
		aigenUi.AudioSettings(1)
		return true
	}
	if strings.Contains(message, "TurnVoiceOff") {
		aigenUi.AudioSettings(0)
		return true
	}
	if strings.Contains(message, "TurnVoiceOn") {
		aigenUi.AudioSettings(1)
		return true
	}
	if strings.Contains(message, "Shutdown Platform") ||
		strings.Contains(message, "SHUTDOWN") {
		os.Exit(0)
		return true
	}
	return false
}

// botMessages function to display messages from the bot
// This function is used to split messages into multiple chat bubbles if the message is too long
// This function is also used to send voice notes if the message is too long
func botMessages(messageCall string, err error, tab1 *fyne.Container, contentType string) {
	//Send voice note if message is more than 120 characters
	if contentType == "text" {

		//Check if message length is "platform" playable
		if len(messageCall) > 0 {
			if getAudioSettings() {
				log.Println(getAudioSettings())
				sendAudio, _ := pressPlayAudio(messageCall)

				if sendAudio != true {
					log.Printf("Error sending audio: %v", sendAudio)
				}
			}
			addChatBubble(tab1, "Sage:"+messageCall, false)
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

func fineTuneWeather() string {
	var weatherInfo string
	weatherInfo = aigenRest.GetCurrentWeather()
	weatherDetailsFineTune := WeatherReport + weatherInfo
	return weatherDetailsFineTune
}

func triggerWeatherInfo(message string) bool {
	return strings.Contains(message, "What is the weather") || strings.Contains(message, "check weather") ||
		strings.Contains(message, "Hey Sage, what is the weather") ||
		strings.Contains(message, "what's the weather") ||
		strings.Contains(message, "what is the weather") ||
		strings.Contains(message, "What is the weather like") ||
		strings.Contains(message, "What do you think about today's weather") ||
		strings.Contains(message, "What is the weather like today") ||
		strings.Contains(message, "Weather tips") ||
		strings.Contains(message, "Tell me about the weather today") ||
		strings.Contains(message, "Today weather") ||
		strings.Contains(message, "Today's weather conditions")
}

func triggerPublishToX(message string) bool {
	return strings.Contains(message, "Send Tweet") || strings.Contains(message, "Tweet") ||
		strings.Contains(message, "Twitter") || strings.Contains(message, "send tweet") ||
		strings.Contains(message, "Send tweet") || strings.Contains(message, "Create a tweet") ||
		strings.Contains(message, "Post content")
}

func triggerImageGeneration(message string) bool {
	return strings.Contains(message, "Image") || strings.Contains(message, "image") ||
		strings.Contains(message, "photo") || strings.Contains(message, "Photo") ||
		strings.Contains(message, "Generate") || strings.Contains(message, "Generate Image")
}

func triggerProgramRun(message string) bool {
	return strings.Contains(message, "run program") ||
		strings.Contains(message, "open program")
}
