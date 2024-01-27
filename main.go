package main

import (
	"aigen/aigenUi"
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
	//Run Setup Scripts
	setup()
	//Start Application
	mapungubwe := app.New()

	mapungubwe.Settings().SetTheme(theme.DarkTheme())
	aigenUi.SwitchUp(mapungubwe)
	playWelcomeSound()

	tabs, inputBoxContainer := mainApp(mapungubwe)
	window := mapungubwe.NewWindow(aigenUi.MainTitle)
	window.SetIcon(theme.MailAttachmentIcon())
	window.SetFixedSize(false)

	window.CenterOnScreen()
	window.Resize(aigenUi.WindowSize)
	window.SetPadded(true)
	scrollApp := container.NewScroll(tabs)

	window.SetContent(container.NewBorder(nil, inputBoxContainer, nil, nil, scrollApp))
	window.ShowAndRun()
	window.SetOnClosed(aigenUi.GoodBye(mapungubwe))
}
