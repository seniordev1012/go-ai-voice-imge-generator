package aigenUi

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"log"
)

func Extensions(mapungubwe fyne.App) *container.TabItem {
	//content := container.NewAdaptiveGrid(4, container.NewVScroll(container.NewAppTabs()))
	tokensForm := container.NewAdaptiveGrid(4, container.NewVBox(widget.NewTextGridFromString("Enter Tokens")))
	loginBtn := widget.NewButton("Login", func() {
		log.Println("Login")
	})
	// Create the login form
	usernameField := widget.NewEntry()
	passwordField := widget.NewPasswordEntry()

	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "Email", Widget: usernameField},
			{Text: "Password", Widget: passwordField},
		},
		OnSubmit: loginBtn.OnTapped,
		OnCancel: func() {
			log.Println("Cancelled")
		},
	}
	loginTab := container.NewTabItem("Login", form)
	loginTab.Icon = theme.LoginIcon()
	// Add the login tab to the tabs container
	cardin := widget.NewCard("", "", tokensForm)
	extensionsTab := container.NewTabItem("Extensions", widget.NewAccordion(
		widget.NewAccordionItem("Sage", widget.NewCheck("Web Search", func(OnandOff bool) {

		})),

		widget.NewAccordionItem("Calendar", cardin),
		widget.NewAccordionItem("EmotionalAI", widget.NewRichTextFromMarkdown("Enable Emotions")),
		widget.NewAccordionItem("Whatsapp", &widget.Button{
			OnTapped: func() {
				fmt.Print("Hello")
			},
		}),
		widget.NewAccordionItem("Twitter", widget.NewEntry()),
		widget.NewAccordionItem("OpenAI", widget.NewEntry()),
		widget.NewAccordionItem("Humor", widget.NewEntry()),
		widget.NewAccordionItem("Web Search", widget.NewEntry()),
	),
	)
	extensionsTab.Icon = theme.ListIcon()
	return extensionsTab
}
