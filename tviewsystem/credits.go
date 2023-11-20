package tviewsystem

import "github.com/rivo/tview"

func ShowCredits() {
	modal := tview.NewModal().
		SetText("Made by Medzernik ©2023").
		AddButtons([]string{"OK"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonLabel == "OK" {
				app.SetRoot(mainView, false)
			}
		})

	app.SetRoot(modal, true).SetFocus(modal)
}
