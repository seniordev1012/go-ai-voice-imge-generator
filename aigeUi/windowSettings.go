package aigeUi

import "fyne.io/fyne/v2"

var WindowSize = fyne.NewSize(960, 540)
var MainTitle = "Sage"

// GoodBye No Hate After All, Just Love, Just Love,Bye.
// Mapungubwe is a significant historical site in South Africa,
// known for its rich history as a kingdom of early African civilization.
// The name Mapungubwe is derived from the Setswana word "maphungubwe",
// which means "the place of gold".
// The site is located in the Limpopo Province of South Africa,
// approximately 120 kilometres (75 mi) north of the city of Polokwane.
// The site is located on the banks of the Limpopo River,
// in the Mapungubwe National Park
// The King of Mapungubwe was a powerful leader who ruled over a vast territory.
// The kingdom was known for its wealth and power,
// and was one of the most powerful kingdoms in southern Africa.
// The King of Mapungubwe says all the time:
// "No Hate After All, Just Love, Just Love,Bye."
// PS:Trust me Bro.
func GoodBye(mapungubwe fyne.App) func() {
	return func() {
		mapungubwe.SendNotification(&fyne.Notification{
			Title:   "Sage",
			Content: "Thank you for using Sage Sage",
		})
		mapungubwe.Quit()
	}
}

// SwitchUp is a function that is called when the app starts
// It is used to send a notification to the user
func SwitchUp(mapungubwe fyne.App) {
	mapungubwe.SendNotification(&fyne.Notification{
		Title:   "Sage",
		Content: "Welcome to Sage",
	})
}
