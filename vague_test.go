package vague

import (
	"fmt"
	"testing"
	"time"
)

type vagueTest struct {
	Date     time.Time
	Expected string
}

type daysTest struct {
	Date     time.Time
	Expected int
}

var vagueTests = []vagueTest{
	{newDate(2017, time.January, 1), "at the beginning of January"},
	{newDate(2017, time.January, 10), "towards the beginning of January"},
	{newDate(2017, time.January, 11), "in the middle of January"},
	{newDate(2017, time.January, 21), "in the middle of January"},
	{newDate(2017, time.January, 22), "towards the end of January"},
	{newDate(2017, time.January, 31), "at the end of January"},

	{newDate(2017, time.February, 1), "at the beginning of February"},
	{newDate(2017, time.February, 9), "towards the beginning of February"},
	{newDate(2017, time.February, 10), "in the middle of February"},
	{newDate(2017, time.February, 19), "in the middle of February"},
	{newDate(2017, time.February, 20), "towards the end of February"},
	{newDate(2017, time.February, 28), "at the end of February"},

	{newDate(2017, time.April, 1), "at the beginning of April"},
	{newDate(2017, time.April, 10), "towards the beginning of April"},
	{newDate(2017, time.April, 11), "in the middle of April"},
	{newDate(2017, time.April, 20), "in the middle of April"},
	{newDate(2017, time.April, 21), "towards the end of April"},
	{newDate(2017, time.April, 30), "at the end of April"},

	{newDate(2020, time.February, 1), "at the beginning of February"},
	{newDate(2020, time.February, 9), "towards the beginning of February"},
	{newDate(2020, time.February, 10), "in the middle of February"},
	{newDate(2020, time.February, 20), "in the middle of February"},
	{newDate(2020, time.February, 21), "towards the end of February"},
	{newDate(2020, time.February, 29), "at the end of February"},
}

var daysTests = []daysTest{
	{newDate(2017, time.January, 1), 31},
	{newDate(2017, time.February, 1), 28},
	{newDate(2017, time.March, 1), 31},
	{newDate(2017, time.April, 1), 30},
	{newDate(2020, time.January, 1), 31},
	{newDate(2020, time.February, 1), 29},
	{newDate(2020, time.March, 1), 31},
	{newDate(2020, time.April, 1), 30},
}

func TestDate(t *testing.T) {
	// Test zero date
	empty := time.Time{}
	if Month(empty) != "never" {
		t.Errorf(`expected: never got: "%s" for: zero date`, Month(empty))
	}

	// Test dates table
	for _, v := range vagueTests {
		// Test without Year
		dateRes := Month(v.Date)
		if dateRes != v.Expected {
			t.Errorf(`expected: "%s" got: "%s" for: %s`, v.Expected, dateRes, v.Date.Format("2006-01-02"))
		}

		// Test with year
		dateYearRes := MonthWithYear(v.Date)
		expected := fmt.Sprintf("%s %d", v.Expected, v.Date.Year())
		if dateYearRes != expected {
			t.Errorf(`expected: "%s" got: "%s" for: %s`, expected, dateYearRes, v.Date.Format("2006-01-02"))
		}
	}
}

func TestTotalDaysInMonth(t *testing.T) {
	for _, v := range daysTests {
		res := totalDaysInMonth(v.Date)
		if res != v.Expected {
			t.Errorf(`expected: %d got: %d`, v.Expected, res)
		}
	}
}

func newDate(year int, month time.Month, day int) time.Time {
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}
