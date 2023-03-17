package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func genSettings(mapungubwe fyne.App) *container.TabItem {
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

// MarketTab is the tab that contains the stock market
func marketTab(mapungubwe fyne.App) *container.TabItem {
	forexMarketTab := container.NewTabItem("Stock Market", widget.NewAccordion(

		widget.NewAccordionItem("Stock Market", widget.NewButtonWithIcon("Stock Market", theme.MailAttachmentIcon(), func() {
			//pop up window
			window := mapungubwe.NewWindow("Stock Market")
			//Add content to the window and show it to the user image
			// <a href="https://docs.google.com/spreadsheets/d/e/2PACX-1vQzV37XPSMjYDi17SoskSvZbp2k3Iu4rAAp6RkU667Hbnd8Z3jO89VywBjYYhkubgMVWxHEmhwtYCS9/pubhtml?gid=0&single=true"><p style="color:#FFF">Link to the Google Sheet</p></a>
			//embed the google sheet in the app

			container.NewAdaptiveGrid(1, container.NewVBox(
				widget.NewCard("Stock Market", "Stock Market", widget.NewLabel("Stock Market Tab Content")),
			),
				widget.NewCard("Stock Market", "Stock Market", widget.NewProgressBarInfinite()))

			// <a href="https://docs.google.com/spreadsheets/d/e/2PACX-1vQzV37XPSMjYDi17SoskSvZbp2k3Iu4rAAp6RkU667Hbnd8Z3jO89VywBjYYhkubgMVWxHEmhwtYCS9/pubhtml?gid=0&single=true"><p style="color:#FFF">Link to the Google Sheet</p></a>
			window.Resize(fyne.NewSize(400, 300))
			window.Show()
		})),

		widget.NewAccordionItem("Add Watchlist", widget.NewLabel("Add a new stock to the watchlist")),

		widget.NewAccordionItem("Create Watchlist", widget.NewLabel("Create a new watchlist")),

		widget.NewAccordionItem("Automated Trading", widget.NewLabel("Automated Trading")),
	),
	)
	return forexMarketTab
}

// StockTab is the tab that contains the stock market
func StockTab() *container.TabItem {
	stockMarketTab := container.NewTabItem("Stock Market", widget.NewAccordion(
		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
	),
	)
	return stockMarketTab
}

// CryptoTab is the tab that contains the stock market
func cTab(mapungubwe fyne.App) *container.TabItem {

	container.NewBorder(nil, nil, nil, nil, widget.NewLabel("Crypto Market Tab Content"))

	cryptoMarketTab := container.NewTabItem("Crypto Market",
		&widget.Button{OnTapped: func() {
			//pop up window
			window := mapungubwe.NewWindow("Crypto Market")
			//Add content to the window and show it to the user image
			container.NewAdaptiveGrid(1, container.NewVBox(
				widget.NewCard("Crypto Market", "Crypto Market", widget.NewLabel("Crypto Market Tab Content")),
			))
			window.SetContent(widget.NewLabel("Crypto Market Tab Content"))
			window.Resize(fyne.NewSize(400, 300))
			window.Show()
		}})
	return cryptoMarketTab
}
