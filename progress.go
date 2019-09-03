package main

import (
	"log"
	"time"

	"github.com/srishanbhattarai/nepcal/dateconv"
	"github.com/tj/go-progress"
)

// BSProgressString returns a progress bar made of Unicode symbols
// for the current progress value of Bikram Sambat
func BSProgressString() string {
	b := progress.NewInt(100)

	b.Width = 15
	b.StartDelimiter = ""
	b.EndDelimiter = ""

	value := int(totalProgress())

	b.Template(`{{.Bar}} {{.Percent}}%`)
	b.ValueInt(value)

	return b.String()
}

// totalProgress calculates the floating progress of BS year up until now
func totalProgress() float32 {
	loc, _ := time.LoadLocation("Asia/Kathmandu")
	t := time.Now().In(loc)

	totalSecondsSpanned := totalSecondsSpanned(t)
	totalSecondsInYear := totalSecondsInYear(t)
	value := float32(totalSecondsSpanned) / float32(totalSecondsInYear)
	percent := value * 100

	log.Printf("Total BS Progress is %f", percent)

	return percent
}

// totalSecondsSpanned calculates the number of seconds that has been
// spanned up until now
func totalSecondsSpanned(t time.Time) int {
	hour := t.Hour()
	min := t.Minute()
	sec := t.Second()

	days, _ := dateconv.TotalDaysSpanned()

	return (days-1)*24*60*60 + hour*60*60 + min*60 + sec
}

// totalSecondsInYear calculates the total number of seconds in the
// current BS year
func totalSecondsInYear(t time.Time) int {
	year, _, _ := dateconv.ToBS(t).Date()
	totalDays, _ := dateconv.TotalDaysInBSYear(year)

	return totalDays * 24 * 60 * 60
}
