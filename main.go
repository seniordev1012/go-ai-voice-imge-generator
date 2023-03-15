package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"

	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// Stock Market Trader App
	dbInit()
	mapungubwe := app.New()
	mapungubwe.Settings().SetTheme(theme.LightTheme())
	mapungubwe.SendNotification(&fyne.Notification{
		Title:   "AiGenie",
		Content: "Welcome to Sage AiGenie",
	})
	mapungubwe.SetIcon(theme.MailAttachmentIcon())

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

		widget.NewAccordionItem("Add Watchlist", widget.NewLabel("Add a new stock to the watchlist")),

		widget.NewAccordionItem("Create Watchlist", widget.NewLabel("Create a new watchlist")),

		widget.NewAccordionItem("Automated Trading", widget.NewLabel("Automated Trading")),
	),
	)

	newsTab := container.NewTabItem("News", widget.NewLabel("News Tab Content"))
	//aiGen := container.NewTabItem("AiGen-Chat", widget.NewLabel("Chat With AiGen"))
	chat := container.NewVBox()

	aiGen := container.NewTabItem("AiGen-Chat", chat)

	// Add chat bubbles to the message box
	messages1, err := getMessages()
	if err != nil {
		log.Printf("Error getting messages: %v", err)
	}

	for _, message := range messages1 {
		addChatBubble(chat, message.Sender+": "+message.Content, message.Sender == "Bot")
	}

	messageCall, checkError := makeApiCall()

	if checkError != nil {
		log.Printf("Error making API call: %v", checkError)
	}

	addChatBubble(chat, "YOU: I am looking for a quote", false)
	addChatBubble(chat, "Bot: "+messageCall, true)

	settingsTab := container.NewTabItem("Settings", widget.NewLabel("Settings Tab Content"))
	//	  <a href="https://docs.google.com/spreadsheets/d/e/2PACX-1vQzV37XPSMjYDi17SoskSvZbp2k3Iu4rAAp6RkU667Hbnd8Z3jO89VywBjYYhkubgMVWxHEmhwtYCS9/pubhtml?gid=0&single=true"><p style="color:#FFF">Link to the Google Sheet</p></a>
	//
	//Create Tab Container
	tabs := container.NewAppTabs(
		aiGen,
		cryptoMarketTab,
		forexMarketTab,
		newsTab,
		stockMarketTab,
		settingsTab,
	)

	inputBox := widget.NewMultiLineEntry()
	inputBox.Wrapping = fyne.TextWrapWord
	inputBox.OnChanged = func(s string) {
		log.Printf("Input changed to: %s", s)
		kitchenLog(s)
	}
	inputBox.PlaceHolder = "Enter your message here..."
	sendButton := sendButton(inputBox, chat)
	inputBoxContainer := container.NewVSplit(inputBox, sendButton)
	//Hide the input box container if it is not chat tab
	//Use `AppTabs.Selected() *TabItem` instead.
	tabs.OnSelected = func(tab *container.TabItem) {
		if tab == aiGen {
			inputBoxContainer.Show()
		} else {
			inputBoxContainer.Hide()
		}
	}

	//Create the main window and set the content to the tabs container
	window := mapungubwe.NewWindow("AiGenie")

	contentSize := fyne.NewSize(600, 400)
	window.Resize(contentSize)
	//Scroll the tabs container
	scrollApp := container.NewVScroll(tabs)
	window.SetContent(container.NewBorder(nil, inputBoxContainer, nil, nil, scrollApp))

	//Show the main window and run the application
	window.ShowAndRun()
}
