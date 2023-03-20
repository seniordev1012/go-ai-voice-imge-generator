package aigenUi

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"log"
)

func GenSettings(mapungubwe fyne.App) *container.TabItem {
	settingsTab := container.NewTabItem("Settings", widget.NewAccordion(

		widget.NewAccordionItem("Ai Voice Reply", widget.NewCheck("Audio Replies", func(OnandOff bool) {
			//TODO: Add a function to toggle the audio replies
			if OnandOff {
				log.Printf("Audio Replies: %v", OnandOff)

			} else {
				log.Printf("Audio Replies: %v", OnandOff)
			}

		})),
		widget.NewAccordionItem("Add Watchlist", widget.NewLabel("Add a new stock to the watchlist")),
	),
	)
	return settingsTab
}
