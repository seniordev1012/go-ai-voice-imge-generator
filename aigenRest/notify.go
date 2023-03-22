package aigenRest

import "fyne.io/fyne/v2"

func SendNotificationNow(message string) {
	n := fyne.NewNotification("Sage", message)
	myApp := fyne.CurrentApp()
	myApp.SendNotification(n)
}
