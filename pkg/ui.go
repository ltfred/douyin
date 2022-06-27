package pkg

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

var path string

func NewGui() {
	a := app.New()
	w := a.NewWindow("DouYin Video Download")
	w.Resize(fyne.NewSize(800, 500))
	url := widget.NewLabel("Urls:")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the video url. Separate multiple url with commas (,).")
	label, label1, button := openFolders(w)
	status := widget.NewLabel("")
	w.SetContent(container.NewVBox(
		url,
		input,
		label,
		label1,
		button,
		widget.NewButton("Download", func() {
			WaterMark(input.Text, path, status)
		}),
		status,
	))

	w.ShowAndRun()
}

func openFolders(w fyne.Window) (*widget.Label, *widget.Label, *widget.Button) {
	label := widget.NewLabel("Folder:")
	label1 := widget.NewLabel("")
	folderOpen := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
		if uri == nil {
			return
		}
		label1.SetText(uri.Path())
		path = uri.Path()
	}, w)

	luri, _ := storage.ListerForURI(storage.NewFileURI("."))
	folderOpen.SetLocation(luri)
	button1 := widget.NewButton("Select Folder", func() {
		folderOpen.Show()
	})

	return label, label1, button1
}
