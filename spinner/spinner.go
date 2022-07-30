package spinner

import (
	"github.com/briandowns/spinner"
	"time"
)

func init() {

}

var Finish = make(chan struct{})

// StartSpin creates a spin and waits for an end signal from the channel
// TODO: Change the finish channel to not be a single channel...
func StartSpin(Finish chan struct{}, incoming string) {
	Spin := spinner.New(spinner.CharSets[52], 100*time.Millisecond, spinner.WithHiddenCursor(true)) // Build our new Spinner

	Spin.Suffix = " " + incoming
	Spin.FinalMSG = "[OK] " + incoming + "\n"
	Spin.Color("blue", "bold")
	Spin.Start()

	<-Finish

	Spin.Stop()
}
