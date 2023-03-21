/*
MIT License

Copyright (c) 2023 Mohammad Shafiee <muhammad.shafiee@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

import (
	"fmt"
	"github.com/mshafiee/jalali"
	"time"
)

func main() {
	// Create a JalaliTime value for the current time
	now := jalali.Now()

	// Print the JalaliTime value
	fmt.Println("Current Jalali Time:", now)

	// Convert the JalaliTime value to a Gregorian time value
	gregorianTime := now.ToGregorian()
	fmt.Println("Current Gregorian Time:", gregorianTime)

	// Format the JalaliTime value
	formatted := now.Format("%Y/%m/%d %T %p")
	fmt.Println("Formatted Jalali Time:", formatted)

	// Calculate the number of days between two JalaliTime values
	j1 := jalali.Date(1399, 1, 1, 0, 0, 0, 0, time.UTC)
	j2 := jalali.Date(1399, 1, 10, 0, 0, 0, 0, time.UTC)
	days := j1.DaysBetween(j2)
	fmt.Println("Days between JalaliTime values:", days)

	layout := "%Y/%m/%d %H:%M:%S"
	value := "1402/05/20 16:30:45"

	jalaliTime, err := jalali.ParseJalali(layout, value)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Parsed Jalali time:", jalaliTime.String())
	}

	jTime := jalali.Date(1402, 5, 20, 16, 30, 45, 0, time.UTC)
	duration := jalali.JalaliDuration{
		Years:  1,
		Months: 2,
		Days:   3,
	}

	newTime := jTime.AddJalaliDuration(duration)
	fmt.Println("After adding duration:", newTime)

	newTime = jTime.SubJalaliDuration(duration)
	fmt.Println("After subtracting duration:", newTime)

	// Create a new recurring event that occurs every two weeks

	// Get a list of occurrences between two dates
	startDate := jalali.Date(1400, jalali.Azar, 1, 0, 0, 0, 0, jalali.Tehran())
	endDate := jalali.Date(1400, jalali.Azar, 30, 0, 0, 0, 0, jalali.Tehran())
	event := jalali.RecurringEvent{
		StartTime: startDate,
		EndTime:   endDate,
		Frequency: 2 * 7 * 24 * time.Hour,
	}

	occurrences := event.Occurrences(startDate, endDate)

	// Loop through the occurrences and do something with each one
	for _, occurrence := range occurrences {
		fmt.Println(occurrence)

	}
}
