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

	// Because dateconv.TotalDaysSpanned() is based on UTC (if system
	// time is in UTC), we have a skewed result when calculating
	// total seconds spanned. For example 2019/09/10 18:15:00 UTC
	// is 2019/09/11 00:00:00 NST. On the next second, the total
	// seconds spanned drops to a lower value as total days
	// spanned still remains the same because of the
	// +05:45 difference.
	days, _ := dateconv.TotalDaysSpanned()

	// The workaround is to increase the day count by 1
	// until UTC time reaches a new day (i.e. +5:45)
	// so that total seconds spanned doesn't drop off
	// inconsistently.
	skew := 0

	// If current time is less than 5:45
	if hour*60*60+min*60 < 5*60*60+45*60 {
		skew = 24 * 60 * 60
	}

	totalSeconds := (days-1)*24*60*60 + hour*60*60 + min*60 + sec + skew

	return totalSeconds
}

// totalSecondsInYear calculates the total number of seconds in the
// current BS year.
func totalSecondsInYear(t time.Time) int {
	year, _, _ := dateconv.ToBS(t).Date()
	totalDays, _ := dateconv.TotalDaysInBSYear(year)
	totalSeconds := totalDays * 24 * 60 * 60

	return totalSeconds
}

// convToNepaliNumeral converts arabic numeral to it's nepali counterpart.
func convToNepaliNumeral(value int) string {
	text := ""
	for _, c := range strconv.Itoa(value) {
		text = text + nepaliNumerals[string(c)]
	}

	return text
}
