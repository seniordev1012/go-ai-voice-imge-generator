package aigenUi

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// UserMedia TODO: Use this as a template for creating new tabs
func UserMedia() *container.TabItem {

	return container.NewTabItem("Media", widget.NewLabel("News Tab Content yr"))
}
