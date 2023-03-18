package main

import (
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

// bubbleCardContent is a function that sets the content of the message card
// The message card is the card that contains the message
// cardImage is the image of the sender or the receiver
func bubbleCardContent(messageCard *widget.Card, cardImage *canvas.Image) {
	messageCard.Image = cardImage
	messageCard.SetTitle("You")
	messageCard.SetSubTitle("Today")
}
