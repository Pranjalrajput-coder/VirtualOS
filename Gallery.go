package main

import (
	// "fmt"
	"io/ioutil"
	"log"
	
	"strings"

	"fyne.io/fyne/v2"

	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/canvas"
)
func galleryapp(w fyne.Window) {

	

	w.CenterOnScreen()
	w.SetFixedSize(true)
	// w.Resize(fyne.NewSize(1000, 600))


	root_src:="D:\\PEPCODING DEV\\HACKATHON 2\\Wallpaper";

	files, err := ioutil.ReadDir(root_src)
    if err != nil {
        log.Fatal(err)
    }

	// var appendarray [] string
	tabs := container.NewAppTabs()
	
    for _, file := range files {
        // fmt.Println(file.Name(), file.IsDir())

		if !file.IsDir(){
			extension:= strings.Split(file.Name(), ".")[1]
			if extension =="jpg" || extension=="png"{
				// appendarray = append(appendarray, root_src +"\\"+file.Name());
				image:= canvas.NewImageFromFile(root_src +"\\"+file.Name())
			tabs.Append(container.NewTabItem(file.Name(), image))
			}
		}
	}


	tabs.SetTabLocation(container.TabLocationLeading)


	galleryContainer := (tabs)

	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, galleryContainer),)
	
	w.Show()
}