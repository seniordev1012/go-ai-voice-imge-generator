package aigenUi

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"os"
	"path/filepath"
	"strings"
)

// UserMedia TODO: Use this as a template for creating new tabs
func UserMedia() *container.TabItem {
	imageDir := "dalleAssets"

	// Create a list to hold the image previews
	imageList := make([]fyne.CanvasObject, 0)

	// List files in the specified directory
	err := filepath.Walk(imageDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isImageFile(path) {
			// Load and display the image
			img := loadImage(path)
			imageList = append(imageList, img)
		}
		return nil
	})

	if err != nil {
		widget.NewLabel(err.Error())
	}

	// Create a scrollable container to hold the images
	scrollable := container.NewVScroll(container.New(layout.NewGridLayout(3), imageList...))

	return container.NewTabItem("Media", scrollable)
}

// Check if a file has an image extension
func isImageFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".png" || ext == ".jpg" || ext == ".jpeg" || ext == ".gif" || ext == ".bmp" || ext == ".webp"
}

// Load an image from file and create a Fyne canvas.Image
func loadImage(filePath string) *canvas.Image {
	img := canvas.NewImageFromFile(filePath)
	img.SetMinSize(fyne.NewSize(200, 200))
	img.FillMode = canvas.ImageFillContain
	return img
}
