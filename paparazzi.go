package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

func chatAvatars() (*canvas.Image, *canvas.Image) {
	//Add image card
	image := canvas.NewImageFromFile("source/avatar.jpg")
	image.SetMinSize(fyne.NewSize(100, 100))
	imageCard := widget.NewCard("", "", image)
	imageCard.Resize(fyne.NewSize(100, 100))
	//Create a new image widget with the avatar URL
	avatarImg := canvas.NewImageFromFile("source/avatar.jpg")
	avatarImg.ScaleMode = canvas.ImageScaleSmooth
	avatarImg.Refresh()
	avatarImg.SetMinSize(fyne.NewSize(100, 100))
	avatarImg.Resize(fyne.NewSize(64, 64))

	botAvatarImg := canvas.NewImageFromFile("source/botAvatar.png")
	botAvatarImg.SetMinSize(fyne.NewSize(64, 64))
	botAvatarImg.Move(fyne.NewPos(-5, -5))
	return avatarImg, botAvatarImg
}
