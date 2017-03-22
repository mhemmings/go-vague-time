// Package vague produces vague human readable strings from Golang time.Time
package vague

import (
	"fmt"
	"time"
)

// Month produces a string representation of what part of a month the time occurs
// The possible outcomes are:
// 		at the start of <month>
// 		towards the start of <month>
// 		in the middle of <month>
// 		towards the end of <month>
// 		at the end of <month>
func Month(t time.Time) string {
	return month(t, false)
}

// MonthWithYear is a helper with the same output as Month, but with the year appended
func MonthWithYear(t time.Time) string {
	return month(t, true)
}

func month(t time.Time, includeYear bool) string {
	if t.IsZero() {
		return "never"
	}

	var prefix string

	daysInMonth := totalDaysInMonth(t)
	cut := daysInMonth / 3

	if t.Day() == 1 {
		prefix = "at the beginning of"
	} else if t.Day() == daysInMonth {
		prefix = "at the end of"
	} else if t.Day() <= cut {
		prefix = "towards the beginning of"
	} else if t.Day() > daysInMonth-cut {
		prefix = "towards the end of"
	} else {
		prefix = "in the middle of"
	}

	if includeYear {
		return fmt.Sprintf("%s %s %d", prefix, t.Month(), t.Year())
	}

	return fmt.Sprintf("%s %s", prefix, t.Month())
}

func totalDaysInMonth(t time.Time) int {
	return time.Date(t.Year(), t.Month()+1, 0, 0, 0, 0, 0, t.Location()).Day()
}
