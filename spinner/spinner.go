package spinner

import (
	"fmt"
	"github.com/briandowns/spinner"
	"time"
)

func init() {

}

var Finish = make(chan struct{})

// SpinnerFun creates a spin and waits for an end signal from the channel
// TODO: Change the finish channel to not be a single channel...
func SpinnerFun(Finish chan struct{}, incoming string) {

	spin := spinner.New(spinner.CharSets[78], 100*time.Millisecond, spinner.WithHiddenCursor(true), spinner.WithColor("blue"), spinner.WithSuffix(incoming)) // Build our new Spinner
	spin.Start()
	<-Finish
	spin.Stop()
	fmt.Println("")

}
