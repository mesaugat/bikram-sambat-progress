package main

import (
	"testing"

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
