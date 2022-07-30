// Package tviewsystem that makes the terminal GUI possible.
package tviewsystem

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
	"os"
)

//Initiate the main application var
var app = tview.NewApplication()

//Testing
var statusTextView = tview.NewTextView().
	SetChangedFunc(func() { app.Draw() }).
	SetTextAlign(tview.AlignCenter).
	SetMaxLines(10).
	SetScrollable(true).
	SetDynamicColors(true).
	SetWrap(true).
	SetWordWrap(true).
	SetText("Please enter your name:").
	SetBorder(true)

//Initiate the main view
var mainView = tview.NewFlex().
	AddItem(tview.NewBox().SetBorder(true).SetTitle("Left (1/2 x width of Top)"), 0, 1, false).
	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(tview.NewBox().SetBorder(true).SetTitle("CleverFox 2 Go Edition"), 0, 1, false).
		AddItem(statusTextView, 5, 1, false), 0, 2, true).
	AddItem(tview.NewBox().SetBorder(true).SetTitle("Right (20 cols)"), 20, 1, false)

//Initiate the quit dialog
var quitDialog = tview.NewModal().
	SetText("Are you sure you want to exit?").
	AddButtons([]string{"Cancel", "Quit"}).
	SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonIndex == 1 {
			app.Stop()
			os.Exit(0)
		} else {
			app.SetRoot(mainView, true)
		}
	})

// StartGUI Starts and sets up the main GUI
func StartGUI() error {
	app.SetRoot(mainView, true).EnableMouse(true)

	//Capture ESC key for a dialog to quit the bot.
	app.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		// Anything handled here will be executed on the main thread
		switch event.Key() {
		case tcell.KeyEsc:
			// Exit the application
			app.SetRoot(quitDialog, true)
			return nil
		case tcell.KeyDelete:
			//statusTextVi

		}

		return event

	})

	//Start the GUI. Fail if cannot be started (TODO: Make a non-GUI version)
	if err := app.Run(); err != nil {
		return err
	}

	return nil
}

// DrawStatus - function to draw a new status?
func DrawStatus() {

}
