package aigenUi

import "fyne.io/fyne/v2/container"

type aigenUi interface {
	container.AppTabs
	container.TabItem
	container.ScrollDirection
	container.Scroll
}
