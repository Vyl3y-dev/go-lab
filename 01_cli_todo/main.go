package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("To-Do App")
	myWindow.SetFullScreen(true)

	label := widget.NewLabel("YOANY IT IS WORKING 💙")
	myWindow.SetContent(container.NewVBox(label))

	myWindow.ShowAndRun()
}
