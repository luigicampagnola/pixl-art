package ui

import "fyne.io/fyne/v2/container"

func Setup(app *AppInit) {
	swatchesContainer := BuildSwatches(app)
	colorPicker := setupColorPicker(app)

	appLayout := container.NewBorder(nil, swatchesContainer, nil, colorPicker)
	app.PixlWindow.SetContent(appLayout)
}
