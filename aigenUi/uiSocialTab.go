package aigenUi

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// SocialTab is the tab that contains the stock market
func SocialTab() *container.TabItem {
	socialTabCon := container.NewTabItem("Social", widget.NewAccordion(
		widget.NewAccordionItem("Social", widget.NewLabel("Social Tab Content")),
		widget.NewAccordionItem("Social", widget.NewLabel("Social Tab Content")),
		widget.NewAccordionItem("Social", widget.NewLabel("Social Tab Content")),
		widget.NewAccordionItem("Social", widget.NewLabel("Social Tab Content")),
	),
	)
	return socialTabCon
}
