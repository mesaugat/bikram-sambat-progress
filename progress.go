package main

import (
	"log"
	"time"

	"github.com/srishanbhattarai/nepcal/nepcal"
	"github.com/tj/go-progress"
)

// ProgressString returns a progress bar made of Unicode symbols
// for the progress value.
func ProgressString(value int) string {
	b := progress.NewInt(100)

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
func TotalProgress() float32 {
	t := time.Now().In(time.UTC)

	totalSecondsSpanned := totalSecondsSpanned(t)
	totalSecondsInYear := totalSecondsInYear(t)
	value := float32(totalSecondsSpanned) / float32(totalSecondsInYear)
	percent := value * 100

	log.Printf("Total BS Progress is %f", percent)

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
		return (hour*60*60 + min*60 + sec) - (18*3600 + 15*60)
	}

	return totalSeconds
}

// totalSecondsInYear calculates the total number of seconds in the
// current BS year.
func totalSecondsInYear(t time.Time) int {
	totalDays := nepcal.FromGregorianUnchecked(t).NumDaysInYear()
	totalSeconds := totalDays * 24 * 60 * 60

	return totalSeconds
}
