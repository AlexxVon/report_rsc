package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/theme"
)

func main() {
	load_data()
	getEventToday()
	getEventAll()
	// for i, k := range diagnosis {
	// 	fmt.Println(i, " ", k)
	// }
	a = app.New()
	w = a.NewWindow("Обезьяна")
	w.Resize(fyne.NewSize(600, 600))
	w.CenterOnScreen()
	w.SetMaster()
	w.SetContent(setContent())
	w.SetMainMenu(setMenu())
	a.Settings().SetTheme(theme.DarkTheme())

	w.Show()
	a.Run()

}
