// Package tviewsystem that makes the terminal GUI possible.
package tviewsystem

import (
	"bytes"
	"fmt"
	"os"
	"time"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

// Initiate the main application var
var app = tview.NewApplication()

// Testing
var StatusTextView = tview.NewTextView()
var MainTextView = tview.NewTextView()
var ActionTextView = tview.NewList()
var UserTextView = tview.NewList()

// For some exemplary reason this must be done like this...
// init updates all the functions that otherwise return stuff, while keeping the vars in global scope of the package.
func init() {
	StatusViewInit(StatusTextView)
	MainViewInit(MainTextView)
	ActionViewInit(ActionTextView)
	UserViewInit(UserTextView)
}

func StatusViewInit(StatusTextView *tview.TextView) {
	StatusTextView.SetBorder(true)
	StatusTextView.SetTitle("Status")
	StatusTextView.SetScrollable(true)
}

func MainViewInit(MainTextView *tview.TextView) {
	MainTextView.SetBorder(true)
	MainTextView.SetTitle("CleverFox 2 Go Edition")
	MainTextView.SetScrollable(true)
}

func ActionViewInit(ActionTextView *tview.List) {
	ActionTextView.SetBorder(true)
	ActionTextView.SetTitle("Actions")
	ActionTextView.AddItem("Credits", "", 0, func() { ShowCredits() })
}

func UserViewInit(UserTextView *tview.List) {
	UserTextView.SetBorder(true)
	UserTextView.SetTitle("Members")

}

// Initiate the main view
var mainView = tview.NewFlex().
	AddItem(ActionTextView, 0, 1, true).
	AddItem(tview.NewFlex().SetDirection(tview.FlexRow).
		AddItem(MainTextView, 0, 1, false).
		AddItem(StatusTextView, 5, 1, false), 0, 2, false).
	AddItem(UserTextView, 0, 1, false)

// Initiate the quit dialog
var quitDialog = tview.NewModal().
	SetText("Are you sure you want to exit?").
	AddButtons([]string{"Cancel", "Quit"}).
	SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonIndex == 1 {
			app.Stop()
			os.Exit(0)
		} else if buttonIndex == 0 {
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
			StatusPush("Hi I am a cool box.")

		}

		return event

	})

	//Start the GUI. Fail if cannot be started (TODO: Make a non-GUI version)
	if err := app.Run(); err != nil {
		return err
	}

	return nil
}

// StatusPush - function to draw a new status?
func StatusPush(update string) {
	//StatusTextView.SetScrollable(true)
	StatusTextView.Write(bytes.NewBufferString("[" + time.Now().Format(time.Kitchen) + "] " + update + "\n").Bytes())
	app.Draw()
}

// StatusPush - function to draw a new status?
func MainViewPush(update string) {
	MainTextView.Write(bytes.NewBufferString("[" + time.Now().Format(time.Kitchen) + "] " + update + "\n").Bytes())
	app.Draw()
}

func MemberListPush(username []string, nick []string) {
	fmt.Println("AHOJ")
	UserTextView.Clear()
	for index, user := range username {
		UserTextView.AddItem(user, fmt.Sprint("  ", nick[index], "\n"), 0, func() { UserAction() })
	}

	app.QueueUpdateDraw(func() {})
}
