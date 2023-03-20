package aigenUi

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewsTab() *container.TabItem {
	return container.NewTabItem("News", widget.NewLabel("News Tab Content"))
}
