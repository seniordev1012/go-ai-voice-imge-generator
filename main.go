package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	// Stock Market Trader App
	mapungubwe := app.New()
	mapungubwe.Settings().SetTheme(theme.LightTheme())
	mapungubwe.SendNotification(&fyne.Notification{
		Title:   "Sage Chatbot",
		Content: "Welcome to Sage Chatbot",
	})
	mapungubwe.SetIcon(theme.MailAttachmentIcon())

	//Create 3 Tabs, one for each of the 3 main functions of the app
	//Home Tab(Automated Trading Inputs)
	//1. Stock Market
	//2. Crypto Market
	//3. Forex Market
	//4. News
	//5. Settings
	//Create 3 Tabs, one for each of the 3 main functions of the app
	//Home Tab(Automated Trading Inputs)
	//1. Stock Market
	//2. Crypto Market
	//3. Forex Market
	//4. News
	//5. Settings

	//Create 3 tabs

	stockMarketTab := container.NewTabItem("Stock Market", widget.NewAccordion(
		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
	),
	)
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

		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),

		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),

		widget.NewAccordionItem("Stock Market", widget.NewLabel("Stock Market Tab Content")),
	),
	)

	newsTab := container.NewTabItem("News", widget.NewLabel("News Tab Content"))
	settingsTab := container.NewTabItem("Settings", widget.NewLabel("Settings Tab Content"))
	//	  <a href="https://docs.google.com/spreadsheets/d/e/2PACX-1vQzV37XPSMjYDi17SoskSvZbp2k3Iu4rAAp6RkU667Hbnd8Z3jO89VywBjYYhkubgMVWxHEmhwtYCS9/pubhtml?gid=0&single=true"><p style="color:#FFF">Link to the Google Sheet</p></a>
	//
	//Create Tab Container
	tabs := container.NewAppTabs(
		stockMarketTab,
		cryptoMarketTab,
		forexMarketTab,
		newsTab,
		settingsTab,
	)

	//Create the main window and set the content to the tabs container
	window := mapungubwe.NewWindow("Stock Market Trader App")

	contentSize := fyne.NewSize(800, 600)
	window.Resize(contentSize)

	window.SetContent(tabs)

	//Show the main window and run the application
	window.ShowAndRun()

	//Show the main window and run the application
	window.ShowAndRun()
}
