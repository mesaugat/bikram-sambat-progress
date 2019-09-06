package main

import (
	"log"
	"strconv"
	"time"

	"github.com/srishanbhattarai/nepcal/dateconv"
	"github.com/tj/go-progress"
)

var nepaliNumerals = map[string]string{
	"0": "०",
	"1": "१",
	"2": "२",
	"3": "३",
	"4": "४",
	"5": "५",
	"6": "६",
	"7": "७",
	"8": "८",
	"9": "९",
}

// ProgressString returns a progress bar made of Unicode symbols
// for the progress value.
func ProgressString(value int) string {
	b := progress.NewInt(100)

	b.Width = 15
	b.StartDelimiter = ""
	b.EndDelimiter = ""

	text := convToNepaliNumeral(value)

	b.Template(`{{.Bar}} {{.Text}}%`)
	b.ValueInt(value)
	b.Text(text)

	return b.String()
}

// TotalProgress calculates the floating progress of BS year up until now.
func TotalProgress() float32 {
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
// spanned up until now.
func totalSecondsSpanned(t time.Time) int {
	hour := t.Hour()
	min := t.Minute()
	sec := t.Second()

	days, _ := dateconv.TotalDaysSpanned()

	return (days-1)*24*60*60 + hour*60*60 + min*60 + sec
}

// totalSecondsInYear calculates the total number of seconds in the
// current BS year.
func totalSecondsInYear(t time.Time) int {
	year, _, _ := dateconv.ToBS(t).Date()
	totalDays, _ := dateconv.TotalDaysInBSYear(year)

	return totalDays * 24 * 60 * 60
}

// convToNepaliNumeral converts arabic numeral to it's nepali counterpart.
func convToNepaliNumeral(value int) string {
	text := ""
	for _, c := range strconv.Itoa(value) {
		text = text + nepaliNumerals[string(c)]
	}

	return text
}
