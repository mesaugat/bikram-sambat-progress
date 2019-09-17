package main

import (
	"time"
)

// checkBSProgress checks for total BS progress every second. Whenever
// the initial progress and the running progress has a difference of
// one or smaller than zero it tweets out a progress string. The idea
// is to only care about jumps from one integer value to another
// consecutive integer value upto 100.
//
// TODO: Questionable if this function should be here. Also, this
// might be a noob approach.
func checkBSProgress() {
	currentProgress := int(TotalProgress())
	ticker := time.NewTicker(time.Second * 1)

	var progress string

	for range ticker.C {
		runningProgress := int(TotalProgress())
		diff := runningProgress - currentProgress

		// Increase the current progress value by one if the
		// difference between running progress and current
		// progress is greater than or equal to one
		if diff >= 1 {
			currentProgress = currentProgress + 1
			progress = ProgressString(currentProgress)

			Tweet(progress)
		}

		// Set the current progress value to zero if the difference
		// between running progress and current progress is negative.
		// This is particularly true if running progress comes back
		// down to zero after reaching 100. In this case, tweet the
		// progress as 100 and then as 0.
		if diff < 0 && int(currentProgress) == 0 {
			currentProgress = 0
			hundred := ProgressString(100)
			zero := ProgressString(0)

			Tweet(hundred)
			Tweet(zero)
		}
	}
}

func main() {
	go checkBSProgress()

	select {}
}
