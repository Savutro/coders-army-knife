package main

import (
	"gitlab.com/savutro/coders-army-knife/services/client"
	"gitlab.com/savutro/coders-army-knife/services/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Choose Your Tool")

	// Create the initial content with four buttons
	var initialContent fyne.CanvasObject
	initialContent = container.NewVBox(
		widget.NewButton("Keylogger", func() {
			// Create a new content view with three additional buttons
			newContent := container.NewVBox(
				widget.NewButton("Start Client", func() {
					client.StartClient()
				}),
				widget.NewButton("Start Server", func() {
					server.StartServer()
				}),
				widget.NewButton("Go Back", func() {
					// Go back to the initial content
					myWindow.SetContent(initialContent)
				}),
			)
			myWindow.SetContent(newContent)
		}),
		widget.NewButton("Calculator", nil),
		widget.NewButton("Encoder", nil),
		widget.NewButton("Decoder", nil),
	)

	myWindow.SetContent(initialContent)
	myWindow.ShowAndRun()
}
