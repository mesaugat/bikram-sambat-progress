package main

import (
	"log"
	"time"

	"github.com/srishanbhattarai/nepcal/nepcal"
	"github.com/tj/go-progress"
)

// CheckBSProgress checks for total BS progress every second. Whenever
// the initial progress and the running progress has a difference of
// one or smaller than zero it tweets out a progress string. The idea
// is to only care about jumps from one integer value to another
// consecutive integer value upto 100.
func CheckBSProgress() {
	currentProgress := int(TotalProgress())
	ticker := time.NewTicker(time.Second * 1)

	for range ticker.C {
		runningProgress := int(TotalProgress())
		diff := runningProgress - currentProgress

		// Increase the current progress value by one if the
		// difference between running progress and current
		// progress is greater than or equal to one.
		if diff >= 1 {
			currentProgress = currentProgress + 1
			progress := ProgressString(currentProgress)

			Tweet(progress)
		}

		// Set the current progress value to zero if the difference
		// between running progress and current progress is negative.
		// This is particularly true if running progress comes back
		// down to zero after reaching 100. In this case, tweet the
		// progress as 100 and then as 0.
		if diff < 0 && int(runningProgress) == 0 {
			currentProgress = 0
			hundred := ProgressString(100)
			zero := ProgressString(0)

			Tweet(hundred)
			Tweet(zero)
		}
	}
}

// ProgressString returns a progress bar made of Unicode symbols
// for the progress value.
func ProgressString(value int) string {
	b := progress.NewInt(100)

	// Use dark shade unicode character instead of a filled block
	// because it increases the total size of the progress bar when
	// using fonts that don't support the block element. MacOS uses
	// Helvetica Neue which is not an unicode font.
	//
	// Check the length of the progress bar of these two tweets in MacOS:
	// -->	https://twitter.com/bikram_sambat/status/1382034603372908550
	// -->  https://twitter.com/bikram_sambat/status/1382034607676219392
	b.Filled = "▓"
	b.Width = 15
	b.StartDelimiter = ""
	b.EndDelimiter = ""

	text := ConvToNepaliNumeral(value)

	b.Template(`{{.Bar}} {{.Text}}%`)
	b.ValueInt(value)
	b.Text(text)

	return b.String()
}

// TotalProgress calculates the floating progress of BS year up until now.
func TotalProgress() float64 {
	t := time.Now().In(time.UTC)

	totalSecondsSpanned := totalSecondsSpanned(t)
	totalSecondsInYear := totalSecondsInYear(t)
	value := float64(totalSecondsSpanned) / float64(totalSecondsInYear)
	percent := value * 100

	log.Printf("Bikram Sambat Progress: %f", percent)

	return percent
}

// totalSecondsSpanned calculates the number of seconds that has been
// spanned up until now.
func totalSecondsSpanned(t time.Time) int {
	hour := t.Hour()
	min := t.Minute()
	sec := t.Second()

	bsDate := nepcal.FromGregorianUnchecked(t)
	days := bsDate.NumDaysSpanned()
	totalSecondsInYear := totalSecondsInYear(t)

	skew := 5*3600 + 45*60

	totalSeconds := (days-1)*24*3600 + hour*3600 + min*60 + sec + skew

	// Special case to handle total seconds spanned for the 1st day
	// of a new year. The total number of seconds will overflow the
	// total number of seconds in a year for a period of 18:15 - 24:00
	// i.e. 5 hour 45 minutes.
	if totalSeconds > totalSecondsInYear {
		return (hour*3600 + min*60 + sec) - (18*3600 + 15*60)
	}

	return totalSeconds
}

// totalSecondsInYear calculates the total number of seconds in the
// current BS year.
func totalSecondsInYear(t time.Time) int {
	totalDays := nepcal.FromGregorianUnchecked(t).NumDaysInYear()
	totalSeconds := totalDays * 24 * 3600

	return totalSeconds
}
