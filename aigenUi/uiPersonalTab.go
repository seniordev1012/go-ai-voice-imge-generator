package aigenUi

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// PersonalTab is the tab that contains the stock market
func PersonalTab(mapungubwe fyne.App) *container.TabItem {
	planPage := container.NewTabItem("Plan", widget.NewAccordion(

		widget.NewAccordionItem("Plan", widget.NewButtonWithIcon("Plan", theme.MailAttachmentIcon(), func() {
			//pop up window
			window := mapungubwe.NewWindow("Stock Market")

			container.NewAdaptiveGrid(1, container.NewVBox(
				widget.NewCard("Plan", "Plan", widget.NewLabel("Stock Market Tab Content")),
			),
				widget.NewCard("Plan", "Stock Market", widget.NewProgressBarInfinite()))

			// <a href="https://docs.google.com/spreadsheets/d/e/2PACX-1vQzV37XPSMjYDi17SoskSvZbp2k3Iu4rAAp6RkU667Hbnd8Z3jO89VywBjYYhkubgMVWxHEmhwtYCS9/pubhtml?gid=0&single=true"><p style="color:#FFF">Link to the Google Sheet</p></a>
			window.Resize(fyne.NewSize(400, 300))
			window.Show()
		})),

		widget.NewAccordionItem("Add Watchlist", widget.NewLabel("Add a new stock to the watchlist")),

		widget.NewAccordionItem("Create Watchlist", widget.NewLabel("Create a new watchlist")),

		widget.NewAccordionItem("Automated Trading", widget.NewLabel("Automated Trading")),
	),
	)
	return planPage
}
