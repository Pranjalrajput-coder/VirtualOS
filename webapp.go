package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var myApp fyne.App = app.New()
var myWindow fyne.Window = myApp.NewWindow("WEB APP")

var btn1 fyne.Widget
var btn2 fyne.Widget
var btn3 fyne.Widget
var btn4 fyne.Widget

// var btn4 fyne.Widget

var img fyne.CanvasObject
var DeskBtn fyne.Widget

var panelContent *fyne.Container

func main() {

	img = canvas.NewImageFromFile("D:\\PEPCODING DEV\\HACKATHON 2\\Wallpaper\\Wallpaper 7.jpg")

	btn1 = widget.NewButtonWithIcon("Weather App", theme.InfoIcon(), func() {
		weatherapp(myWindow)
	})

	btn2 = widget.NewButtonWithIcon("Calculator App", theme.ContentAddIcon(), func() {
		calculatorapp()
	})

	btn3 = widget.NewButtonWithIcon("Gallery App", theme.StorageIcon(), func() {
		galleryapp(myWindow)
	})

	btn4 = widget.NewButtonWithIcon("MoviesInfo App",theme.FileVideoIcon() , func(){
		movieinfo(myWindow);
	})
	DeskBtn = widget.NewButtonWithIcon("Home", theme.HomeIcon(), func() {
		myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))
	})

	panelContent = container.NewVBox(container.NewGridWithColumns(5, DeskBtn, btn1, btn2, btn3,btn4)) // for menu on top panel like home,calculator,etc


	myWindow.Resize(fyne.NewSize(1000, 700))
	myWindow.SetPadded(true)
	myWindow.CenterOnScreen()
	myWindow.SetContent(container.NewBorder(panelContent, nil, nil, nil, img))

	myWindow.ShowAndRun()

}
