package aigenUi

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// TODO: Use this as a template for creating new tabs
func NewsTab() *container.TabItem {
	return container.NewTabItem("Sage Tab", widget.NewLabel("News Tab Content"))
}
