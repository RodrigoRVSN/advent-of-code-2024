package main

import (
	"testing"
)

func TestIsSequenceSafe(t *testing.T) {
	tests := []struct {
		name          string
		isDecreasing  bool
		currentLevel  int
		previousLevel int
		want          bool
	}{
		{"Valid Decrease Within 3", true, 5, 8, true},
		{"Valid Decrease With 1", true, 5, 6, true},
		{"Invalid Decrease Beyond 3", true, 4, 8, false},
		{"Valid Increase Within 3", false, 8, 5, true},
		{"Valid Increase With 1", false, 8, 7, true},
		{"Invalid Increase With beyond 3", false, 8, 4, false},
		{"Invalid No Change", false, 5, 5, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := isSequenceSafe(tt.isDecreasing, tt.currentLevel, tt.previousLevel)
			if got != tt.want {
				t.Errorf("isSequenceSafe() = %v, want %v", got, tt.want)
			}
		})
	}
}
