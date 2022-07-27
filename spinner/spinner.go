package spinner

import (
	"github.com/briandowns/spinner"
	"time"
)

func init() {

}

var Finish = make(chan struct{})

var Finished = make(chan struct{})

// StartSpin creates a spin and waits for an end signal from the channel
// TODO: Change the finish channel to not be a single channel...
func StartSpin(Finish chan struct{}, incoming string) {
	spin := spinner.New(spinner.CharSets[78], 100*time.Millisecond, spinner.WithHiddenCursor(true)) // Build our new Spinner
	spin.Suffix = " " + incoming
	spin.FinalMSG = "[OK] " + incoming + "\n"
	spin.Color("blue", "bold")
	spin.Start()

	<-Finish

	spin.Stop()
	//fmt.Println("")
	Finished <- struct{}{}

}
