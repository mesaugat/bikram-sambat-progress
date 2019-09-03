package main

import (
	"time"
)

// checkBSProgress checks for total BS progress every second. Whenever
// the initial progress and the running progress has a difference of
// one, it tweets out a progress string.
func checkBSProgress() {
	currentProgress := TotalProgress()
	ticker := time.NewTicker(time.Second * 1)

	for range ticker.C {
		runningProgress := TotalProgress()
		diff := int(runningProgress) - int(currentProgress)

		if diff >= 1 {
			currentProgress = runningProgress
			progress := ProgressString(int(currentProgress))

			Tweet(progress)
		}
	}
}

func main() {
	go checkBSProgress()

	select {}
}
