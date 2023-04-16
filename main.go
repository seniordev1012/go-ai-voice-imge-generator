package main

import (
	"aigen/aigeUi"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	_ "github.com/mattn/go-sqlite3"
)

//func updateTime(clock *widget.Label) {
//	formatted := time.Now().Format("Time: 03:04:05")
//	clock.SetText(formatted)
//}

func main() {
	MigrationAssist()
	setup()

	// Stock Market Trader App
	//PlayNotificationSound()
	//Whisper()
	//pathFinder, _ := VoiceRecorder()
	//if pathFinder != "" {
	//	log.Println(pathFinder)
	//}

	mapungubwe := app.New()
	//mapungubwe.Settings().SetTheme(theme.LightTheme())
	aigeUi.SwitchUp(mapungubwe)
	mapungubwe.SetIcon(theme.MailAttachmentIcon())

	tabs, inputBoxContainer := mainApp(mapungubwe)

	//Main Window
	window := mapungubwe.NewWindow(aigeUi.MainTitle)
	window.SetIcon(theme.MailAttachmentIcon())
	window.SetFixedSize(true)

	window.CenterOnScreen()
	window.Resize(aigeUi.WindowSize)
	window.SetPadded(false)
	scrollApp := container.NewScroll(tabs)
	// Create a background image object

	window.SetContent(container.NewBorder(nil, inputBoxContainer, nil, nil, scrollApp))
	window.ShowAndRun()
	window.SetOnClosed(aigeUi.GoodBye(mapungubwe))
}
