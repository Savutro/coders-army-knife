package main

import (
	"gitlab.com/savutro/coders-army-knife/services/client"
	"gitlab.com/savutro/coders-army-knife/services/server"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Coders Army Knife")
	myWindow.Resize(fyne.NewSize(600, 400))

	// Create the initial content with four buttons
	var initialContent fyne.CanvasObject
	initialContent = container.NewVBox(
		container.NewHBox(
			container.NewVBox(
				createRow("cube.png", "Keylogger", "App Collection to log Keys.", "GO", nil, func() {
					// Create a new content view with three additional buttons
					newContent := container.NewVBox(
						createRow("cube.png", "Client", "Sender", "Start Client", nil, func() {
							go client.StartHttpClient()
						}),
						createRow("cube.png", "Server", "Receiver", "Start Server", nil, func() {
							go server.StartHttpServer()
						}),
						widget.NewButton("Go Back", func() {
							// Go back to the initial content
							myWindow.SetContent(initialContent)
						}),
					)
					myWindow.SetContent(newContent)
				}),
				createRow("cube.png", "Calculator", "Perform calculations.", "GO", nil, nil),
				createRow("cube.png", "Encoder", "Encode data.", "GO", nil, nil),
				createRow("cube.png", "Decoder", "Decode data.", "GO", nil, nil),
			),
		),
	)

	myWindow.SetContent(initialContent)
	myWindow.ShowAndRun()
}

func createRow(imageSrc, cardTitle, cardDescription string, buttonLabel string, buttonImage fyne.Resource, onPressed func()) *fyne.Container {
	image := canvas.NewImageFromFile("./assets/" + imageSrc)
	image.FillMode = canvas.ImageFillStretch
	image.Resize(fyne.NewSquareSize(100))

	card := widget.NewCard(cardTitle, cardDescription, nil)
	card.Resize(fyne.NewSize(300, 200))

	button := widget.NewButtonWithIcon(buttonLabel, buttonImage, onPressed)
	button.Importance = widget.MediumImportance
	button.Resize(fyne.NewSize(100, 30))

	row := container.NewHBox(
		container.NewVBox(image),
		container.NewVBox(card),
		container.NewVBox(widget.NewLabel(" "), button),
	)

	return row
}
