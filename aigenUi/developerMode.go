package aigenUi

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Developer UserMedia TODO: Use this as a template for creating new tabs
func Developer() *container.TabItem {

	return container.NewTabItem("Developer Mode", widget.NewLabel("Enable Developer Mode"))
}
