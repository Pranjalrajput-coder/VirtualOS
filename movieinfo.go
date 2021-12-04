package main

import (
	"encoding/json"
	"os"

	"io/ioutil"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/widget"
)

type MovieResults struct {
	Result []Movie `json:"results"`
}
type Movie struct {
	Title    string `json:"title"`
	Overview string `json:"overview"`
}

func loadMovies() (MovieResults, error) {
	data, err := ioutil.ReadFile("./data.json")
	if err != nil {
		return MovieResults{}, err
	}
	var movieResults MovieResults
	err = json.Unmarshal(data, &movieResults)
	if err != nil {
		return MovieResults{}, err
	}
	return movieResults, nil
}

func movieinfo(w fyne.Window) {
	os.Setenv("FYNE_THEME", "dark")

	movieResults, err := loadMovies()
	if err != nil {
		panic(err)
	}

	//
	// w.Resize(fyne.NewSize(500, 500))

	w.CenterOnScreen()
	w.SetFixedSize(true)
	// w.Resize(fyne.NewSize(500, 500))

	listView := widget.NewList(
		func() int {
			return len(movieResults.Result)

		}, func() fyne.CanvasObject {

			return widget.NewLabel("")

		}, func(lii widget.ListItemID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(movieResults.Result[lii].Title)
		})

	content := widget.NewLabel("Select the movie ")
	content.Wrapping = fyne.TextWrapWord

	listView.OnSelected = func(id widget.ListItemID) {
		content.Text = movieResults.Result[id].Overview
		content.Refresh()
	}
	movieContainer := (container.NewHSplit(listView,
		container.NewMax(content),
	))

	w.SetContent(container.NewBorder(panelContent, nil, nil, nil, movieContainer))
	w.Show()
}
