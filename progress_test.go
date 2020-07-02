package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProgressString(t *testing.T) {
	tests := []struct {
		name        string
		value       int
		expectedStr string
	}{
		{"Zero", 0, "░░░░░░░░░░░░░░░ ०%"},
		{"Twenty One", 21, "████░░░░░░░░░░░ २१%"},
		{"Thirty Five", 35, "██████░░░░░░░░░ ३५%"},
		{"Sixty", 60, "█████████░░░░░░ ६०%"},
		{"Hundred", 100, "███████████████ १००%"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			progress := ProgressString(test.value)

			assert.Equal(t, test.expectedStr, progress)
		})
	}
}

// The total number of seconds in a Bikram Sambat year is either
// 31536000 (365 days) or 31622400 (366 days).
func TestTotalSecondsInYear(t *testing.T) {
	tests := []struct {
		name                 string
		time                 time.Time
		expectedTotalSeconds int
	}{
		{"2075 B.S.", time.Date(2018, 4, 28, 0, 0, 0, 0, time.UTC), 31536000},
		{"2076 B.S.", time.Date(2019, 4, 28, 0, 0, 0, 0, time.UTC), 31536000},
		{"2077 B.S.", time.Date(2020, 4, 28, 0, 0, 0, 0, time.UTC), 31622400},
		{"2078 B.S.", time.Date(2021, 4, 28, 0, 0, 0, 0, time.UTC), 31536000},
		{"2079 B.S.", time.Date(2022, 4, 28, 0, 0, 0, 0, time.UTC), 31536000},
		{"2080 B.S.", time.Date(2023, 4, 28, 0, 0, 0, 0, time.UTC), 31536000},
		{"2081 B.S.", time.Date(2024, 4, 28, 0, 0, 0, 0, time.UTC), 31622400},
		{"2089 B.S.", time.Date(2032, 4, 28, 0, 0, 0, 0, time.UTC), 31536000},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			totalSeconds := totalSecondsInYear(test.time)

			assert.Equal(t, test.expectedTotalSeconds, totalSeconds)
		})
	}
}

// I'm sorry, these tests are unreadable at best.
func TestTotalSecondSpanned(t *testing.T) {
	tests := []struct {
		name                 string
		time                 time.Time
		expectedTotalSeconds int
	}{
		{"1_Min_Before_Chaitra_30_2075", time.Date(2019, 4, 12, 18, 14, 0, 0, time.UTC), 31449540},
		{"Chaitra_30_2075", time.Date(2019, 4, 12, 18, 15, 0, 0, time.UTC), 31449600},
		{"1_Min_After_Chaitra_30_2075", time.Date(2019, 4, 12, 18, 16, 0, 0, time.UTC), 31449660},
		{"5_Hours_44_Minutes_After_Chaitra_30_2075", time.Date(2019, 4, 12, 23, 59, 0, 0, time.UTC), 31470240},
		{"5_Hours_45_Minutes_After_Chaitra_30_2075", time.Date(2019, 4, 13, 0, 0, 0, 0, time.UTC), 31470300},
		{"5_Hours_46_Minutes_After_Chaitra_30_2075", time.Date(2019, 4, 13, 0, 1, 0, 0, time.UTC), 31470360},
		{"1_Min_Before_Baisakh_1_2076", time.Date(2019, 4, 13, 18, 14, 0, 0, time.UTC), 31535940},
		{"Baisakh_1_2076", time.Date(2019, 4, 13, 18, 15, 0, 0, time.UTC), 31536000},
		{"1_Min_After_Baisakh_1_2076", time.Date(2019, 4, 13, 18, 16, 0, 0, time.UTC), 60},
		{"5_Hours_44_Minutes_After_Baisakh_1_2076", time.Date(2019, 4, 13, 23, 59, 0, 0, time.UTC), 20640},
		{"5_Hours_45_Minutes_After_Baisakh_1_2076", time.Date(2019, 4, 14, 0, 0, 0, 0, time.UTC), 20700},
		{"5_Hours_46_Minutes_After_Baisakh_1_2076", time.Date(2019, 4, 14, 0, 1, 0, 0, time.UTC), 20760},
		{"18_Hours_15_Minutes_After_Baisakh_1_2076", time.Date(2019, 4, 14, 12, 30, 0, 0, time.UTC), 65700},
		{"Jestha_1_2076", time.Date(2019, 5, 15, 18, 15, 0, 0, time.UTC), 2764800},
		{"Baisakh_1_2077", time.Date(2020, 4, 12, 18, 15, 0, 0, time.UTC), 31536000},
		{"Baisakh_1_2088", time.Date(2021, 4, 13, 18, 15, 0, 0, time.UTC), 31622400},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			totalSeconds := totalSecondsSpanned(test.time)

			assert.Equal(t, test.expectedTotalSeconds, totalSeconds)
		})
	}
}
