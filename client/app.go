package client

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func InitApp() fyne.App {
	return app.New()
}

func InitWindow(a fyne.App, width, height float32) fyne.Window {
	w := a.NewWindow("cqupt-grabber")
	if width <= 0 || height <= 0 {
		w.Resize(fyne.NewSize(720, 520))
	} else {
		w.Resize(fyne.NewSize(width, height))
	}
	return w
}
