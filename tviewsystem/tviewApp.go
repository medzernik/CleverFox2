package tviewsystem

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
)

func StartGUI() error {
	app := tview.NewApplication()
	flex := tview.NewFlex().
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
		AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("CleverFox 2 Go Edition"), 0, 1, false).
			AddItem(tview.NewBox().SetBorder(true).SetTitle("Status"), 5, 1, false), 0, 2, false).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)
	app.SetRoot(flex, true).EnableMouse(true)

	//Define the various dialogs, this one is for the quit dialog.
	quitDialog := tview.NewModal().
		SetText("Are you sure you want to exit?").
		AddButtons([]string{"Cancel", "Quit"}).
		SetDoneFunc(func(buttonIndex int, buttonLabel string) {
			if buttonIndex == 1 {
				app.Stop()
				os.Exit(0)
			} else {
				app.SetRoot(flex, true)
			}
		})
	//Capture ESC key for a dialog to quit the bot.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyEsc:
			// Exit the application
			app.SetRoot(quitDialog, true)
			return nil
		}

		return event

	})

	//Start the GUI. Fail if cannot be started (TODO: Make a non-GUI version)
	if err := app.Run(); err != nil {
		return err
	}

	return nil
}
