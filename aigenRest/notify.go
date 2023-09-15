package aigenRest

import "fyne.io/fyne/v2"

// SendNotificationNow Handles Platform notifications.
// So it can be used anywhere notification to user is required on platform FE
func SendNotificationNow(message string) {
	n := fyne.NewNotification("Sage", message)
	myApp := fyne.CurrentApp()
	myApp.SendNotification(n)
}
