package utils

import (
	"errors"
	"strings"
	"time"
)

// ParseEntryDate
// This function will probably need a lot of tweaking in the future to handle
// different requirements and functionalities
func ParseEntryDate(input string) (time.Time, error) {
	input = strings.TrimSpace(strings.ToLower(input))
	now := time.Now()

	if input == "" {
		return time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location()), nil
	}

	if input == "yesterday" {
		d := now.AddDate(0, 0, -1)
		return time.Date(d.Year(), d.Month(), d.Day(), 0, 0, 0, 0, d.Location()), nil
	}

	if t, err := time.Parse("2/1/2006", input); err == nil {
		return t, nil
	}

	if t, err := time.Parse("2/1", input); err == nil {
		return t.AddDate(now.Year(), 0, 0), nil
	}

	return time.Time{}, errors.New("invalid date format: use 'dd/mm/yyyy', 'dd/mm', or 'Yesterday'")
}
