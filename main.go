package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Stock Market Trader App
	MigrationAssist()
	mapungubwe := app.New()
	mapungubwe.Settings().SetTheme(theme.LightTheme())
	mapungubwe.SendNotification(&fyne.Notification{
		Title:   "AiGenie",
		Content: "Welcome to Sage AiGenie",
	})
	mapungubwe.SetIcon(theme.MailAttachmentIcon())

	tabs, inputBoxContainer := mainApp(mapungubwe)

	//Main Window
	window := mapungubwe.NewWindow("AiGenie")
	contentSize := fyne.NewSize(600, 400)
	window.Resize(contentSize)
	scrollApp := container.NewVScroll(tabs)
	window.SetContent(container.NewBorder(nil, inputBoxContainer, nil, nil, scrollApp))
	window.ShowAndRun()
}
