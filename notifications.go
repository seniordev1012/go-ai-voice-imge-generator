package main

import "fyne.io/fyne/v2"

func sendNotificationNow(message string) {
	n := fyne.NewNotification("Ai Gen", message)
	myApp := fyne.CurrentApp()
	myApp.SendNotification(n)
}
