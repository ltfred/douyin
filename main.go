package main

import (
	"douyin-gui/pkg"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"os"
)

var path string

func init() {
	path, _ = os.Getwd()
}

func main() {
	a := app.New()
	w := a.NewWindow("抖音视频去水印下载")
	w.Resize(fyne.NewSize(800, 500))
	url := widget.NewLabel("Urls:")
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter the video url. Separate multiple url with commas (,).")
	label, label1, button := openFolders(w)
	w.SetContent(container.NewVBox(
		url,
		input,
		label,
		label1,
		button,
		widget.NewButton("download", func() {
			pkg.WaterMark(input.Text, path)
		}),
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
	button1 := widget.NewButton("Show folderOpen", func() {
		folderOpen.Show()
	})

	return label, label1, button1
}
