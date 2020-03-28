package main

import "strconv"

// Nepali representation of arabic numeral
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

// ConvToNepaliNumeral converts arabic numeral to it's nepali counterpart.
func ConvToNepaliNumeral(value int) string {
	text := ""
	for _, c := range strconv.Itoa(value) {
		text = text + nepaliNumerals[string(c)]
	}

	return text
}
