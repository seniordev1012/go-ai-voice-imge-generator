package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	MigrationAssist()
	// Stock Market Trader App
	//PlayNotificationSound()
	//Whisper()
	//pathFinder, _ := VoiceRecorder()
	//if pathFinder != "" {
	//	log.Println(pathFinder)
	//}

	mapungubwe := app.New()
	//mapungubwe.Settings().SetTheme(theme.LightTheme())
	switchUp(mapungubwe)
	mapungubwe.SetIcon(theme.MailAttachmentIcon())

	tabs, inputBoxContainer := mainApp(mapungubwe)

	//Main Window
	window := mapungubwe.NewWindow(mainTitle)
	window.SetIcon(theme.MailAttachmentIcon())
	window.SetFixedSize(true)

	window.CenterOnScreen()
	window.Resize(windowSize)
	window.SetPadded(false)
	scrollApp := container.NewScroll(tabs)
	// Create a background image object

	window.SetContent(container.NewBorder(nil, inputBoxContainer, nil, nil, scrollApp))
	window.ShowAndRun()
	window.SetOnClosed(goodBye(mapungubwe))
}
