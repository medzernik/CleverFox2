package tviewsystem

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"log"
)

//var modal = tview.NewModal()

func StartGUI() {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("CleverFox 2 Go Edition"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Status"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	app.SetRoot(flex, true).EnableMouse(true)
	modal := tview.NewModal().
		SetText("Are you sure you want to exit?").
		AddButtons([]string{"Cancel", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 1 {
				app.Stop()
			} else {
				app.SetRoot(flex, true)
			}
		})
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyEsc:
			// Exit the application
			app.SetRoot(modal, true)
			return nil
		}

		return event

	})

	if err := app.Run(); err != nil {
		log.Fatalf("critical GUI error: %s", err)
	}

}
