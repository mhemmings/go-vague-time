package vague

import (
	"fmt"
	"time"
)

func ExampleMonth() {
	d := Month(time.Date(2017, time.July, 5, 0, 0, 0, 0, time.UTC))
	fmt.Println(d)
}

func ExampleMonthWithYear() {
	d := MonthWithYear(time.Date(2020, time.January, 15, 0, 0, 0, 0, time.UTC))
	fmt.Println(d)
}
