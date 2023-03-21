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

package jalali

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
	"time"
)

func TestWeekdayString(t *testing.T) {
	testCases := []struct {
		weekday Weekday
		wantEn  string
		wantFa  string
	}{
		{Yekshanbe, "1Shanbeh", "یکشنبه"},
		{Doshanbe, "2Shanbeh", "دوشنبه"},
		{Seshanbe, "3Shanbeh", "سه‌شنبه"},
		{Chaharshanbe, "4Shanbeh", "چهارشنبه"},
		{Panjshanbe, "5Shanbeh", "پنج‌شنبه"},
		{Joomeh, "Joomeh", "جمعه"},
		{Shanbe, "Shanbeh", "شنبه"},
	}

	for _, tc := range testCases {
		t.Run(tc.weekday.String(), func(t *testing.T) {
			gotEn := tc.weekday.String()
			gotFa := tc.weekday.FaString()
			if gotEn != tc.wantEn {
				t.Errorf("Weekday(%v).String() = %v, want %v", tc.weekday, gotEn, tc.wantEn)
			}
			if gotFa != tc.wantFa {
				t.Errorf("Weekday(%v).FaString() = %v, want %v", tc.weekday, gotFa, tc.wantFa)
			}
		})
	}
}

func TestMonthString(t *testing.T) {
	testCases := []struct {
		month  Month
		wantEn string
		wantFa string
	}{
		{Farvardin, "Farvardin", "فروردین"},
		{Ordibehesht, "Ordibehesht", "اردیبهشت"},
		{Khordad, "Khordad", "خرداد"},
		{Tir, "Tir", "تیر"},
		{Mordad, "Mordad", "مرداد"},
		{Shahrivar, "Shahrivar", "شهریور"},
		{Mehr, "Mehr", "مهر"},
		{Aban, "Aban", "آبان"},
		{Azar, "Azar", "آذر"},
		{Dey, "Dey", "دی"},
		{Bahman, "Bahman", "بهمن"},
		{Esfand, "Esfand", "اسفند"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%v", tc.month), func(t *testing.T) {
			gotEn := tc.month.String()
			gotFa := tc.month.FaString()
			if gotEn != tc.wantEn {
				t.Errorf("Month(%v).String() = %v, want %v", tc.month, gotEn, tc.wantEn)
			}
			if gotFa != tc.wantFa {
				t.Errorf("Month(%v).FaString() = %v, want %v", tc.month, gotFa, tc.wantFa)
			}
		})
	}
}

func TestJalaliTime(t *testing.T) {
	// Test values for March 20, 2023 12:34:56 UTC+03:30 (8 Esfand 1401 12:04:56)
	testCases := []struct {
		jalaliTime JalaliTime
		year       int
		month      Month
		day        int
		hour       int
		minute     int
		second     int
		yearDay    int
	}{
		{
			JalaliTime{1401, Esfand, 28, 12, 34, 56, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			1401, Esfand, 28, 12, 34, 56, 364,
		},
		{
			JalaliTime{1397, Bahman, 5, 15, 30, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			1397, Bahman, 5, 15, 30, 0, 311,
		},
		{
			JalaliTime{1399, Ordibehesht, 15, 23, 0, 30, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			1399, Ordibehesht, 15, 23, 0, 30, 46,
		},
		{
			JalaliTime{1400, Tir, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			1400, Tir, 1, 0, 0, 0, 94,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.jalaliTime.loc.String(), func(t *testing.T) {
			// Test Year
			if got := tc.jalaliTime.Year(); got != tc.year {
				t.Errorf("Year() = %v, want %v", got, tc.year)
			}

			// Test Month
			if got := tc.jalaliTime.Month(); got != tc.month {
				t.Errorf("Month() = %v, want %v", got, tc.month)
			}

			// Test Day
			if got := tc.jalaliTime.Day(); got != tc.day {
				t.Errorf("Day() = %v, want %v", got, tc.day)
			}

			// Test Hour
			if got := tc.jalaliTime.Hour(); got != tc.hour {
				t.Errorf("Hour() = %v, want %v", got, tc.hour)
			}

			// Test Minute
			if got := tc.jalaliTime.Minute(); got != tc.minute {
				t.Errorf("Minute() = %v, want %v", got, tc.minute)
			}

			// Test Second
			if got := tc.jalaliTime.Second(); got != tc.second {
				t.Errorf("Second() = %v, want %v", got, tc.second)
			}

			// Test YearDay
			if got := tc.jalaliTime.YearDay(); got != tc.yearDay {
				t.Errorf("YearDay() = %v, want %v", got, tc.yearDay)
			}
		})
	}
}

func TestJalaliTime_Weekday(t *testing.T) {
	testCases := []struct {
		jalaliTime JalaliTime
		weekday    Weekday
	}{
		{
			JalaliTime{1401, Esfand, 28, 12, 34, 56, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			Yekshanbe,
		},
		{
			JalaliTime{1397, Bahman, 5, 15, 30, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			Joomeh,
		},
		{
			JalaliTime{1399, Ordibehesht, 15, 23, 0, 30, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			Doshanbe,
		},
		{
			JalaliTime{1400, Tir, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			Seshanbe,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.jalaliTime.loc.String(), func(t *testing.T) {
			// Test Weekday
			if got := tc.jalaliTime.Weekday(); got != tc.weekday {
				t.Errorf("Weekday() = %v, want %v", got, tc.weekday)
			}
		})
	}
}

func TestDate(t *testing.T) {
	// Test values for March 20, 2023 12:34:56 UTC+03:30 (8 Esfand 1401 12:04:56)
	testCases := []struct {
		year  int
		month Month
		day   int
		hour  int
		min   int
		sec   int
		nsec  int
		loc   *time.Location
	}{
		{
			1401, Esfand, 28, 12, 34, 56, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second))),
		},
		{
			1397, Bahman, 5, 15, 30, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second))),
		},
		{
			1399, Ordibehesht, 15, 23, 0, 30, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second))),
		},
		{
			1400, Tir, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second))),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.loc.String(), func(t *testing.T) {
			// Create a new JalaliTime value
			jTime := Date(tc.year, tc.month, tc.day, tc.hour, tc.min, tc.sec, tc.nsec, tc.loc)

			// Check the fields of the JalaliTime value
			if jTime.year != tc.year {
				t.Errorf("Year = %d, want %d", jTime.year, tc.year)
			}
			if jTime.month != tc.month {
				t.Errorf("Month = %v, want %v", jTime.month, tc.month)
			}
			if jTime.day != tc.day {
				t.Errorf("Day = %d, want %d", jTime.day, tc.day)
			}
			if jTime.hour != tc.hour {
				t.Errorf("Hour = %d, want %d", jTime.hour, tc.hour)
			}
			if jTime.min != tc.min {
				t.Errorf("Minute = %d, want %d", jTime.min, tc.min)
			}
			if jTime.sec != tc.sec {
				t.Errorf("Second = %d, want %d", jTime.sec, tc.sec)
			}
			if jTime.nsec != tc.nsec {
				t.Errorf("Nanosecond = %d, want %d", jTime.nsec, tc.nsec)
			}
			if jTime.loc != tc.loc {
				t.Errorf("Location = %v, want %v", jTime.loc, tc.loc)
			}
		})
	}
}

func TestJalaliTime_DaysInMonth(t *testing.T) {
	// Test values for March 20, 2023 12:34:56 UTC+03:30 (8 Esfand 1401 12:04:56)
	testCases := []struct {
		jalaliTime  JalaliTime
		daysInMonth int
	}{
		{
			JalaliTime{1401, Esfand, 28, 12, 34, 56, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			29,
		},
		{
			JalaliTime{1399, Bahman, 5, 15, 30, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			30,
		},
		{
			JalaliTime{1397, Ordibehesht, 15, 23, 0, 30, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			31,
		},
		{
			JalaliTime{1400, Tir, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))},
			31,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.jalaliTime.loc.String(), func(t *testing.T) {
			// Test DaysInMonth
			if got := tc.jalaliTime.DaysInMonth(); got != tc.daysInMonth {
				t.Errorf("DaysInMonth() = %v, want %v", got, tc.daysInMonth)
			}
		})
	}
}

func TestNow(t *testing.T) {
	// Get the current time in UTC
	nowUTC := time.Now()

	// Calculate the expected Jalali date and time
	jYear, jMonth, jDay := gregorianToJalali(nowUTC.Year(), nowUTC.Month(), nowUTC.Day())
	jHour, jMin, jSec := nowUTC.Hour(), nowUTC.Minute(), nowUTC.Second()

	// Set the expected JalaliTime value
	expectedJalaliTime := JalaliTime{
		year:  jYear,
		month: jMonth,
		day:   jDay,
		hour:  jHour,
		min:   jMin,
		sec:   jSec,
		nsec:  nowUTC.Nanosecond(),
		loc:   time.Local,
	}

	// Call the Now function to get the actual JalaliTime value
	actualJalaliTime := Now()

	// Test that the actual and expected JalaliTime values match to within one second
	if actualJalaliTime.year != expectedJalaliTime.year ||
		actualJalaliTime.month != expectedJalaliTime.month ||
		actualJalaliTime.day != expectedJalaliTime.day ||
		actualJalaliTime.hour != expectedJalaliTime.hour ||
		actualJalaliTime.min != expectedJalaliTime.min ||
		actualJalaliTime.sec != expectedJalaliTime.sec {
		t.Errorf("Now() = %v, want %v", actualJalaliTime, expectedJalaliTime)
	}

	// Test that the actual and expected time zones match
	if actualJalaliTime.loc.String() != expectedJalaliTime.loc.String() {
		t.Errorf("Now() timezone = %v, want %v", actualJalaliTime.loc, expectedJalaliTime.loc)
	}
}

func TestJalaliTime_UTC(t *testing.T) {
	// Test values for March 20, 2023 12:34:56 UTC+03:30 (8 Esfand 1401 12:04:56)
	jalaliTime := JalaliTime{1401, Esfand, 28, 3, 30, 56, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	utcJalaliTime := JalaliTime{1401, Esfand, 28, 0, 0, 56, 0, time.UTC}

	// Test UTC
	if got := jalaliTime.UTC(); got != utcJalaliTime {
		t.Errorf("UTC() = %v, want %v", got, utcJalaliTime)
	}
}

func TestToJalali(t *testing.T) {
	// Test with a specific Gregorian time
	gregorianTime := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	jalaliTime := ToJalali(gregorianTime)
	if jalaliTime.year != 1399 || jalaliTime.month != 10 || jalaliTime.day != 12 {
		t.Errorf("ToJalali(%v) = %v; want %d/%d/%d", gregorianTime, jalaliTime, 1399, 10, 12)
	}

	// Test with the current time
	currentTime := time.Now()
	jalaliTime = ToJalali(currentTime)

	// Check that the year, month, and day are valid
	if !isValidJalaliDate(jalaliTime.year, int(jalaliTime.month), jalaliTime.day) {
		t.Errorf("ToJalali(%v) produced an invalid Jalali date %d-%d-%d", currentTime, jalaliTime.year, jalaliTime.month, jalaliTime.day)
	}

	// Check that the hour, minute, second, and nanosecond are the same
	if jalaliTime.hour != currentTime.Hour() || jalaliTime.min != currentTime.Minute() || jalaliTime.sec != currentTime.Second() || jalaliTime.nsec != currentTime.Nanosecond() {
		t.Errorf("ToJalali(%v) produced an incorrect time %d:%d:%d.%d", currentTime, jalaliTime.hour, jalaliTime.min, jalaliTime.sec, jalaliTime.nsec)
	}

	// Check that the location is the same as the input location
	if jalaliTime.loc.String() != currentTime.Location().String() {
		t.Errorf("ToJalali(%v) produced an incorrect location %v; want %v", currentTime, jalaliTime.loc, currentTime.Location())
	}
}

func TestJalaliTime_ToGregorian(t *testing.T) {
	// Expected conversion result
	expected := time.Date(2023, time.January, 1, 0, 0, 0, 0, time.UTC)

	// Create a JalaliTime instance
	j := JalaliTime{year: 1401, month: 10, day: 11, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC}

	// Convert it to a Gregorian time
	actual := j.ToGregorian()

	// Compare with the expected result
	if actual != expected {
		t.Errorf("JalaliTime.ToGregorian() failed, expected %v, got %v", expected, actual)
	}
}

func TestJalaliTime_ToTime(t *testing.T) {
	j := Date(1399, 7, 13, 10, 22, 45, 0, time.UTC)
	expected := time.Date(2020, time.October, 4, 10, 22, 45, 0, time.UTC)
	result := j.ToTime()
	if !result.Equal(expected) {
		t.Errorf("ToTime() = %v; expected %v", result, expected)
	}
}

func TestJalaliTime_Local(t *testing.T) {
	jalaliTime := JalaliTime{year: 1400, month: 10, day: 30, hour: 12, min: 0, sec: 0, nsec: 0, loc: time.UTC}
	localJalaliTime := jalaliTime.Local()

	// check if the location is changed to local
	if localJalaliTime.loc.String() != time.Local.String() {
		t.Errorf("Local() returned wrong time zone. Expected: %v, got: %v.", time.Local, localJalaliTime.loc)
	}

	// check if other values stay the same
	if localJalaliTime.year != jalaliTime.year || localJalaliTime.month != jalaliTime.month || localJalaliTime.day != jalaliTime.day || localJalaliTime.hour != jalaliTime.hour || localJalaliTime.min != jalaliTime.min || localJalaliTime.sec != jalaliTime.sec || localJalaliTime.nsec != jalaliTime.nsec {
		t.Errorf("Local() changed some values. Expected: %v, got: %v.", jalaliTime, localJalaliTime)
	}
}

func TestJalaliTime_Location(t *testing.T) {
	location, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		t.Errorf("Failed to load location: %v", err)
	}

	j := JalaliTime{
		year:  1400,
		month: 4,
		day:   25,
		hour:  13,
		min:   30,
		sec:   0,
		nsec:  0,
		loc:   location,
	}

	expectedLocation := location
	actualLocation := j.Location()

	if expectedLocation != actualLocation {
		t.Errorf("Expected timezone location %v, but got %v", expectedLocation, actualLocation)
	}
}

func TestJalaliTime_Unix(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Tehran")
	j := JalaliTime{
		year:  1399,
		month: 4,
		day:   12,
		hour:  15,
		min:   30,
		sec:   0,
		nsec:  0,
		loc:   loc,
	}
	expectedUnix := int64(1593687600)
	actualUnix := j.Unix()
	if actualUnix != expectedUnix {
		t.Errorf("Incorrect timestamp. Expected: %d, Actual: %d", expectedUnix, actualUnix)
	}
}

func TestJalaliTime_Zone(t *testing.T) {
	j := JalaliTime{
		year:  1399,
		month: 10,
		day:   22,
		hour:  11,
		min:   0,
		sec:   0,
		nsec:  0,
		loc:   time.FixedZone("Tehran Time", 12600),
	}

	name, offset := j.Zone()

	expectedName := "Tehran Time"
	if name != expectedName {
		t.Errorf("got %v, expected %v for time zone name", name, expectedName)
	}

	expectedOffset := 12600
	if offset != expectedOffset {
		t.Errorf("got %v, expected %v for time zone offset", offset, expectedOffset)
	}
}

func TestJalaliTime_UnixNano(t *testing.T) {
	// Test Timezone: Iran Standard Time (UTC+3:30)
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		t.Fatalf("Failed to load timezone: %v", err)
	}

	// Test Jalali Time
	j := JalaliTime{
		year:  1400,
		month: 6,
		day:   12,
		hour:  8,
		min:   30,
		sec:   0,
		nsec:  0,
		loc:   loc,
	}

	// Expected Unix Nano Time: 1623503400000000000 (2021-06-12 04:00:00 UTC)
	expected := int64(1630641600000000000)

	// Actual Unix Nano Time
	actual := j.UnixNano()

	// Compare expected and actual Unix Nano Time
	if actual != expected {
		t.Errorf("Expected UnixNano() to return %v, but got %v", expected, actual)
	}
}

func TestJalaliTime_Format(t *testing.T) {
	j := JalaliTime{year: 1400, month: 6, day: 20, hour: 10, min: 30, sec: 15, loc: time.FixedZone("Asia/Tehran", 12600)}
	expectedOutput := "1400-06-20 10:30:15 صبح"
	actualOutput := j.Format("%Y-%m-%d %T %p")
	if actualOutput != expectedOutput {
		t.Errorf("Expected '%s', but got '%s'", expectedOutput, actualOutput)
	}
}

func TestFormat(t *testing.T) {
	// Initialize a JalaliTime struct for testing
	j := JalaliTime{1380, 7, 25, 10, 25, 30, 0, time.UTC}

	// Define the test cases as input-output pairs
	tests := []struct {
		layout   string
		expected string
	}{
		{"%Y/%m/%d", "1380/07/25"},
		{"%y/%B/%d", "80/مهر/25"},
		{"%w%n%R", "چهارشنبه\n10:25"},
		{"%T %p", "10:25:30 صبح"},
		{"%z %Z", "+0000 UTC"},
	}

	// Loop through the test cases and compare the output with the expected value
	for _, tc := range tests {
		result := j.Format(tc.layout)
		if result != tc.expected {
			t.Errorf("Format(%v) = %v, want %v", tc.layout, result, tc.expected)
		}
	}
}

func TestFormatShort(t *testing.T) {
	// Test the format for YYYY/MM/DD
	jt := JalaliTime{1398, 2, 20, 23, 59, 59, 0, time.Local}
	expected := "1398/02/20"
	result := jt.FormatShort()
	if result != expected {
		t.Errorf("FormatShort - expected: %s, but got: %s", expected, result)
	}
}

func TestFormatLong(t *testing.T) {
	// Test the format for "DD MonthName YYYY"
	jt := JalaliTime{1398, 2, 20, 23, 59, 59, 0, time.Local}
	expected := "20 اردیبهشت 1398"
	result := jt.FormatLong()
	if result != expected {
		t.Errorf("FormatLong - expected: %s, but got: %s", expected, result)
	}
}

func TestString(t *testing.T) {
	// Test the format for YYYY/MM/DD HH:MM:SS
	jt := JalaliTime{1398, 2, 20, 23, 59, 59, 0, time.Local}
	expected := "1398/02/20 23:59:59"
	result := jt.String()
	if result != expected {
		t.Errorf("String - expected: %s, but got: %s", expected, result)
	}
}

func TestDaysBetween(t *testing.T) {
	testCases := []struct {
		name     string
		j1       JalaliTime
		j2       JalaliTime
		expected int
	}{
		{
			name:     "Same day",
			j1:       JalaliTime{year: 1400, month: Tir, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			j2:       JalaliTime{year: 1400, month: Tir, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			expected: 0,
		},
		{
			name:     "One day apart",
			j1:       JalaliTime{year: 1400, month: Tir, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			j2:       JalaliTime{year: 1400, month: Tir, day: 2, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			expected: 1,
		},
		{
			name:     "One month apart",
			j1:       JalaliTime{year: 1400, month: Farvardin, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			j2:       JalaliTime{year: 1400, month: Ordibehesht, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			expected: 31,
		},
		{
			name:     "One year apart",
			j1:       JalaliTime{year: 1400, month: Farvardin, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			j2:       JalaliTime{year: 1401, month: Farvardin, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			expected: 365,
		},
		{
			name:     "One year apart with leap",
			j1:       JalaliTime{year: 1399, month: Farvardin, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			j2:       JalaliTime{year: 1400, month: Farvardin, day: 1, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			expected: 366,
		},
		{
			name:     "Different years and months",
			j1:       JalaliTime{year: 1401, month: Khordad, day: 15, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			j2:       JalaliTime{year: 1402, month: Mehr, day: 20, hour: 0, min: 0, sec: 0, nsec: 0, loc: time.UTC},
			expected: 494,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := tc.j1.DaysBetween(tc.j2)
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func TestAfter(t *testing.T) {
	time1 := JalaliTime{1399, Mordad, 27, 10, 40, 0, 0, time.UTC}
	time2 := JalaliTime{1399, Mordad, 27, 10, 39, 0, 0, time.UTC}
	result := time1.After(time2)
	expected := true
	if result != expected {
		t.Errorf("After() failed, expected %v but got %v", expected, result)
	}
}

func TestBefore(t *testing.T) {
	time1 := JalaliTime{1399, Shahrivar, 27, 10, 39, 0, 0, time.UTC}
	time2 := JalaliTime{1399, Shahrivar, 27, 10, 40, 0, 0, time.UTC}
	result := time1.Before(time2)
	expected := true
	if result != expected {
		t.Errorf("Before() failed, expected %v but got %v", expected, result)
	}
}

func TestEqual(t *testing.T) {
	time1 := JalaliTime{1399, Mehr, 27, 10, 39, 0, 0, time.UTC}
	time2 := JalaliTime{1399, Mehr, 27, 10, 39, 0, 0, time.UTC}
	result := time1.Equal(time2)
	expected := true
	if result != expected {
		t.Errorf("Equal() failed, expected %v but got %v", expected, result)
	}
}

func TestIsZero(t *testing.T) {
	time1 := JalaliTime{}
	result := time1.IsZero()
	expected := true
	if result != expected {
		t.Errorf("IsZero() failed, expected %v but got %v", expected, result)
	}
}

func TestIsLeapJalaliYear(t *testing.T) {
	// Test a non-leap year
	j1 := JalaliTime{year: 1400}
	got := j1.IsLeapJalaliYear()
	want := false
	if got != want {
		t.Errorf("IsLeapJalaliYear() = %v; want %v", got, want)
	}

	// Test a leap year
	j2 := JalaliTime{year: 1399}
	got = j2.IsLeapJalaliYear()
	want = true
	if got != want {
		t.Errorf("IsLeapJalaliYear() = %v; want %v", got, want)
	}

	// Test a negative year
	j3 := JalaliTime{year: -1}
	got = j3.IsLeapJalaliYear()
	want = false
	if got != want {
		t.Errorf("IsLeapJalaliYear() = %v; want %v", got, want)
	}
}

func TestJalaliTime_JulianDate(t *testing.T) {
	// Sample JalaliTime
	jalali := JalaliTime{
		year:  1399,
		month: 1,
		day:   1,
		hour:  0,
		min:   0,
		sec:   0,
		nsec:  0,
		loc:   time.Local,
	}

	// Expected Julian date value
	expected := 2458928.5

	// Actual Julian date value using JalaliTime's JulianDate() method
	actual := jalali.JulianDate()

	// Test for correct Julian date value
	if actual != expected {
		t.Errorf("TestJalaliTime_JulianDate: expected %v, but got %v", expected, actual)
	}

	// Test for valid Julian date range
	if actual < 0 || actual > 9999999999 {
		t.Errorf("TestJalaliTime_JulianDate: expected Julian date within valid range, but got %v", actual)
	}
}

func TestJulianDayNumber(t *testing.T) {
	testCases := []struct {
		year        int
		month       time.Month
		day         int
		hour        int
		minute      int
		second      int
		nanosecond  int
		expectedJD  float64
		expectedErr error
	}{
		// Valid input case
		{
			year:        2022,
			month:       time.January,
			day:         1,
			hour:        0,
			minute:      0,
			second:      0,
			nanosecond:  0,
			expectedJD:  2459580.5,
			expectedErr: nil,
		},
		// Invalid year case
		{
			year:        -4713,
			month:       time.January,
			day:         1,
			hour:        0,
			minute:      0,
			second:      0,
			nanosecond:  0,
			expectedJD:  0,
			expectedErr: fmt.Errorf("year -4713 is too early (minimum is -4712)"),
		},
		// Invalid month case
		{
			year:        2022,
			month:       time.December + 1,
			day:         1,
			hour:        0,
			minute:      0,
			second:      0,
			nanosecond:  0,
			expectedJD:  0,
			expectedErr: fmt.Errorf("invalid month %%!Month(13)"),
		},
		// Invalid day case
		{
			year:        2022,
			month:       time.January,
			day:         32,
			hour:        0,
			minute:      0,
			second:      0,
			nanosecond:  0,
			expectedJD:  0,
			expectedErr: fmt.Errorf("invalid day 32 (must be between 1 and 31)"),
		},
		// Invalid hour case
		{
			year:        2022,
			month:       time.January,
			day:         1,
			hour:        24,
			minute:      0,
			second:      0,
			nanosecond:  0,
			expectedJD:  0,
			expectedErr: fmt.Errorf("invalid hour 24 (must be between 0 and 23)"),
		},
		// Invalid minute case
		{
			year:        2022,
			month:       time.January,
			day:         1,
			hour:        0,
			minute:      60,
			second:      0,
			nanosecond:  0,
			expectedJD:  0,
			expectedErr: fmt.Errorf("invalid minute 60 (must be between 0 and 59)"),
		},
		// Invalid second case
		{
			year:        2022,
			month:       time.January,
			day:         1,
			hour:        0,
			minute:      0,
			second:      60,
			nanosecond:  0,
			expectedJD:  0,
			expectedErr: fmt.Errorf("invalid second 60 (must be between 0 and 59)"),
		},
		// Invalid nanosecond case
		{
			year:        2022,
			month:       time.January,
			day:         1,
			hour:        0,
			minute:      0,
			second:      0,
			nanosecond:  1000000000,
			expectedJD:  0,
			expectedErr: fmt.Errorf("invalid nanosecond 1000000000 (must be between 0 and 999999999)"),
		},
		// Leap year case
		{
			year:        2024,
			month:       time.February,
			day:         29,
			hour:        0,
			minute:      0,
			second:      0,
			nanosecond:  0,
			expectedJD:  2460369.5,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		jd, err := julianDayNumber(tc.year, tc.month, tc.day, tc.hour, tc.minute, tc.second, tc.nanosecond)

		if jd != tc.expectedJD {
			t.Errorf("Expected JD %f, but got %f", tc.expectedJD, jd)
		}

		if err != nil && err.Error() != tc.expectedErr.Error() {
			t.Errorf("Expected error %v, but got %v", tc.expectedErr, err)
		}
	}
}

func TestJalaliTime_Add(t *testing.T) {
	// Sample JalaliTime value
	jt := JalaliTime{
		year:  1400,
		month: 1,
		day:   10,
		hour:  12,
		min:   0,
		sec:   0,
		nsec:  0,
		loc:   time.UTC,
	}

	// Test Cases
	tests := []struct {
		name      string
		duration  time.Duration
		wantYear  int
		wantMonth Month
		wantDay   int
		wantHour  int
		wantMin   int
		wantSec   int
		wantNsec  int
		wantLoc   *time.Location
	}{
		{
			name:      "Add 10 seconds",
			duration:  10 * time.Second,
			wantYear:  1400,
			wantMonth: 1,
			wantDay:   10,
			wantHour:  12,
			wantMin:   0,
			wantSec:   10,
			wantNsec:  0,
			wantLoc:   time.UTC,
		},
		{
			name:      "Add 1 hour",
			duration:  1 * time.Hour,
			wantYear:  1400,
			wantMonth: 1,
			wantDay:   10,
			wantHour:  13,
			wantMin:   0,
			wantSec:   0,
			wantNsec:  0,
			wantLoc:   time.UTC,
		},
		{
			name:      "Add 2 days",
			duration:  2 * 24 * time.Hour,
			wantYear:  1400,
			wantMonth: 1,
			wantDay:   12,
			wantHour:  12,
			wantMin:   0,
			wantSec:   0,
			wantNsec:  0,
			wantLoc:   time.UTC,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Invoke Add() method with the given duration
			newJt := jt.Add(tt.duration)

			// Validate the new JalaliTime value
			if newJt.year != tt.wantYear {
				t.Errorf("year is %v, want %v", newJt.year, tt.wantYear)
			}
			if newJt.month != tt.wantMonth {
				t.Errorf("month is %v, want %v", newJt.month, tt.wantMonth)
			}
			if newJt.day != tt.wantDay {
				t.Errorf("day is %v, want %v", newJt.day, tt.wantDay)
			}
			if newJt.hour != tt.wantHour {
				t.Errorf("hour is %v, want %v", newJt.hour, tt.wantHour)
			}
			if newJt.min != tt.wantMin {
				t.Errorf("min is %v, want %v", newJt.min, tt.wantMin)
			}
			if newJt.sec != tt.wantSec {
				t.Errorf("sec is %v, want %v", newJt.sec, tt.wantSec)
			}
			if newJt.nsec != tt.wantNsec {
				t.Errorf("nsec is %v, want %v", newJt.nsec, tt.wantNsec)
			}
			if newJt.loc != tt.wantLoc {
				t.Errorf("loc is %v, want %v", newJt.loc, tt.wantLoc)
			}
		})
	}
}

func TestAddMonths(t *testing.T) {
	j := JalaliTime{year: 1400, month: 9, day: 20, hour: 17, min: 30, sec: 0, nsec: 0, loc: time.Local}
	result := j.AddMonths(5)
	expected := JalaliTime{year: 1401, month: 2, day: 20, hour: 17, min: 30, sec: 0, nsec: 0, loc: time.Local}
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestAddDays(t *testing.T) {
	j := JalaliTime{year: 1400, month: 9, day: 20, hour: 17, min: 30, sec: 0, nsec: 0, loc: time.Local}
	result := j.AddDays(7)
	expected := JalaliTime{year: 1400, month: 9, day: 27, hour: 17, min: 30, sec: 0, nsec: 0, loc: time.Local}
	if result != expected {
		t.Errorf("Expected %v but got %v", expected, result)
	}
}

func TestAddYears(t *testing.T) {
	// Test case 1: Check that the function returns the correct year
	j := JalaliTime{year: 1399, month: 11, day: 30}
	newJ := j.AddYears(1)
	if newJ.year != 1400 {
		t.Errorf("Expected year to be 1400, but got %d", newJ.year)
	}

	// Test case 2: Check that the function returns the zero value if the new year is before year 1
	j = JalaliTime{year: 1399, month: 11, day: 30}
	newJ = j.AddYears(-1400)
	if newJ.year != 0 {
		t.Errorf("Expected year to be 0, but got %d", newJ.year)
	}

	// Test case 3: Check that the function adjusts the day if the new year is a leap year and the current month is Esfand
	j = JalaliTime{year: 1399, month: 12, day: 30}
	newJ = j.AddYears(4)
	if newJ.day != 29 {
		t.Errorf("Expected day to be 29, but got %d", newJ.day)
	}

	// Test case 4: Check that the function sets all other fields correctly
	j = JalaliTime{
		year:  1399,
		month: 11,
		day:   30,
		hour:  21,
		min:   30,
		sec:   45,
		nsec:  0,
		loc:   time.Local,
	}
	newJ = j.AddYears(1)
	if newJ.hour != 21 || newJ.min != 30 || newJ.sec != 45 || newJ.nsec != 0 || newJ.loc != time.Local {
		t.Errorf("Expected all fields to be the same, but got %+v", newJ)
	}
}

func TestJalaliTime_Sub(t *testing.T) {
	// Use a sample Jalali date to test
	j1 := JalaliTime{1399, 10, 9, 12, 0, 0, 0, time.Local}
	j2 := JalaliTime{1399, 10, 9, 11, 30, 0, 0, time.Local}

	// Expected duration between the two dates in seconds
	expected := time.Duration(30) * time.Minute

	// Test the Sub method for the JalaliTime struct
	result := j1.Sub(j2)

	// Assert that the result matches the expected value within tolerance
	tolerance := time.Second // Some tolerance in comparing durations
	if result < (expected-tolerance) || result > (expected+tolerance) {
		t.Errorf("Sub test failed: expected %v, but got %v", expected, result)
	}
}

func TestRecurringEvent_Occurrences(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name                string
		event               RecurringEvent
		startDate           JalaliTime
		endDate             JalaliTime
		expectedOccurrences []JalaliTime
	}{
		{
			name: "Basic daily event",
			event: RecurringEvent{
				StartTime: JalaliTime{2023, 3, 1, 0, 0, 0, 0, time.UTC},
				EndTime:   JalaliTime{2023, 3, 31, 0, 0, 0, 0, time.UTC},
				Frequency: 24 * time.Hour,
			},
			startDate: JalaliTime{2023, 3, 1, 0, 0, 0, 0, time.UTC},
			endDate:   JalaliTime{2023, 3, 5, 0, 0, 0, 0, time.UTC},
			expectedOccurrences: []JalaliTime{
				{2023, 3, 1, 0, 0, 0, 0, time.UTC},
				{2023, 3, 2, 0, 0, 0, 0, time.UTC},
				{2023, 3, 3, 0, 0, 0, 0, time.UTC},
				{2023, 3, 4, 0, 0, 0, 0, time.UTC},
				{2023, 3, 5, 0, 0, 0, 0, time.UTC},
			},
		},
		// Add more test cases as needed
	}

	// Run test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			occurrences := tc.event.Occurrences(tc.startDate, tc.endDate)

			if len(occurrences) != len(tc.expectedOccurrences) {
				t.Fatalf("expected %d occurrences, got %d", len(tc.expectedOccurrences), len(occurrences))
			}

			for i, occurrence := range occurrences {
				if !occurrence.Equal(tc.expectedOccurrences[i]) {
					t.Errorf("expected occurrence at index %d to be %v, got %v", i, tc.expectedOccurrences[i], occurrence)
				}
			}
		})
	}
}

func TestJalaliTime_DaysUntil(t *testing.T) {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		t.Fatalf("error loading time zone: %v", err)
	}
	j1 := JalaliTime{1399, Mordad, 15, 12, 0, 0, 0, loc}
	j2 := JalaliTime{1399, Mordad, 20, 12, 0, 0, 0, loc}
	want := 5
	got := j1.DaysUntil(j2)
	if got != want {
		t.Errorf("DaysUntil(%v, %v) = %v; want %v", j1, j2, got, want)
	}
}

func TestJalaliFromTime(t *testing.T) {
	gregorianDate := time.Date(2021, 3, 20, 13, 30, 0, 0, time.UTC)
	expectedJalaliDate := JalaliTime{
		year:  1399,
		month: 12,
		day:   30,
		hour:  13,
		min:   30,
		sec:   0,
		nsec:  0,
		loc:   time.UTC,
	}

	// Call JalaliFromTime with the gregorianDate
	jalaliDate := JalaliFromTime(gregorianDate)

	// Check if the returned JalaliTime matches the expected value
	if jalaliDate != expectedJalaliDate {
		t.Errorf("Expected JalaliTime: %v, but got: %v", expectedJalaliDate, jalaliDate)
	}
}

func TestParseJalali(t *testing.T) {
	type args struct {
		layout string
		value  string
	}
	tests := []struct {
		name    string
		args    args
		want    JalaliTime
		wantErr error
	}{
		{
			name: "valid input",
			args: args{
				layout: "%Y-%m-%d %H:%M:%S",
				value:  "1400-01-01 12:00:00",
			},
			want: JalaliTime{
				year:  1400,
				month: Month(1),
				day:   1,
				hour:  12,
				min:   0,
				sec:   0,
				nsec:  0,
				loc:   time.Local,
			},
			wantErr: nil,
		},
		{
			name: "invalid input",
			args: args{
				layout: "%Y-%m-%d %H:%M:%S",
				value:  "1400-13-01 12:00:00",
			},
			want:    JalaliTime{},
			wantErr: errors.New("invalid Jalali date: 1400/13/01"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseJalali(tt.args.layout, tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseJalali() got = %v, want %v", got, tt.want)
			}
			if (err == nil && tt.wantErr != nil) || (err != nil && err.Error() != tt.wantErr.Error()) {
				t.Errorf("ParseJalali() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestJalaliTime_AddDate(t *testing.T) {
	jt := JalaliTime{
		year:  1400,
		month: 1,
		day:   1,
		hour:  0,
		min:   0,
		sec:   0,
		nsec:  0,
		loc:   time.Local,
	}

	tests := []struct {
		name   string
		years  int
		months int
		days   int
		want   JalaliTime
	}{
		{
			name:   "adding 1 year, 1 month and 1 day",
			years:  1,
			months: 1,
			days:   1,
			want: JalaliTime{
				year:  1401,
				month: 2,
				day:   2,
				hour:  0,
				min:   0,
				sec:   0,
				nsec:  0,
				loc:   time.Local,
			},
		},
		{
			name:   "adding 2 years, 6 months and 15 days",
			years:  2,
			months: 6,
			days:   15,
			want: JalaliTime{
				year:  1402,
				month: 7,
				day:   16,
				hour:  0,
				min:   0,
				sec:   0,
				nsec:  0,
				loc:   time.Local,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := jt.AddDate(tt.years, tt.months, tt.days)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddDate() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGregorianToJalali(t *testing.T) {
	cases := []struct {
		gYear  int
		gMonth time.Month
		gDay   int
		jYear  int
		jMonth Month
		jDay   int
	}{
		{2021, 3, 20, 1399, 12, 30},
		{2023, 3, 20, 1401, 12, 29},
		{2023, 3, 21, 1402, 1, 1},
		{2021, 7, 8, 1400, 4, 17},
		{2022, 1, 1, 1400, 10, 11},
		{2021, 1, 1, 1399, 10, 12},
	}

	for _, tc := range cases {
		jYear, jMonth, jDay := gregorianToJalali(tc.gYear, tc.gMonth, tc.gDay)
		if jYear != tc.jYear || jMonth != tc.jMonth || jDay != tc.jDay {
			t.Errorf("gregorianToJalali(%d, %d, %d) = (%d, %d, %d); want (%d, %d, %d)",
				tc.gYear, tc.gMonth, tc.gDay, jYear, jMonth, jDay, tc.jYear, tc.jMonth, tc.jDay)
		}
	}
}

func TestAddJalaliDuration(t *testing.T) {
	// Test adding positive duration to a date
	initialDate := JalaliTime{1399, 2, 28, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	duration := JalaliDuration{1, 2, 3}
	expectedResult := JalaliTime{1400, 4, 31, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	result := initialDate.AddJalaliDuration(duration)
	if result.year != expectedResult.year {
		t.Errorf("Expected year %d, but got %d", expectedResult.year, result.year)
	}
	if result.month != expectedResult.month {
		t.Errorf("Expected month %d, but got %d", expectedResult.month, result.month)
	}
	if result.day != expectedResult.day {
		t.Errorf("Expected day %d, but got %d", expectedResult.day, result.day)
	}
	if result.hour != expectedResult.hour {
		t.Errorf("Expected hour %d, but got %d", expectedResult.hour, result.hour)
	}
	if result.min != expectedResult.min {
		t.Errorf("Expected minute %d, but got %d", expectedResult.min, result.min)
	}
	if result.sec != expectedResult.sec {
		t.Errorf("Expected second %d, but got %d", expectedResult.sec, result.sec)
	}
	if result.nsec != expectedResult.nsec {
		t.Errorf("Expected nanosecond %d, but got %d", expectedResult.nsec, result.nsec)
	}
	if result.loc.String() != expectedResult.loc.String() {
		t.Errorf("Expected location %s, but got %s", expectedResult.loc, result.loc)
	}

	// Test adding negative duration to a date
	initialDate = JalaliTime{1400, 5, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	duration = JalaliDuration{-1, -2, -3}
	expectedResult = JalaliTime{1399, 2, 29, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	result = initialDate.AddJalaliDuration(duration)
	if result.year != expectedResult.year {
		t.Errorf("Expected year %d, but got %d", expectedResult.year, result.year)
	}
	if result.month != expectedResult.month {
		t.Errorf("Expected month %d, but got %d", expectedResult.month, result.month)
	}
	if result.day != expectedResult.day {
		t.Errorf("Expected day %d, but got %d", expectedResult.day, result.day)
	}
	if result.hour != expectedResult.hour {
		t.Errorf("Expected hour %d, but got %d", expectedResult.hour, result.hour)
	}
	if result.min != expectedResult.min {
		t.Errorf("Expected minute %d, but got %d", expectedResult.min, result.min)
	}
	if result.sec != expectedResult.sec {
		t.Errorf("Expected second %d, but got %d", expectedResult.sec, result.sec)
	}
	if result.nsec != expectedResult.nsec {
		t.Errorf("Expected nanosecond %d, but got %d", expectedResult.nsec, result.nsec)
	}
	if result.loc.String() != expectedResult.loc.String() {
		t.Errorf("Expected location %s, but got %s", expectedResult.loc, result.loc)
	}

	// Test adding zero duration to a date
	initialDate = JalaliTime{1400, 5, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	duration = JalaliDuration{0, 0, 0}
	expectedResult = JalaliTime{1400, 5, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	result = initialDate.AddJalaliDuration(duration)
	if result.year != expectedResult.year {
		t.Errorf("Expected year %d, but got %d", expectedResult.year, result.year)
	}
	if result.month != expectedResult.month {
		t.Errorf("Expected month %d, but got %d", expectedResult.month, result.month)
	}
	if result.day != expectedResult.day {
		t.Errorf("Expected day %d, but got %d", expectedResult.day, result.day)
	}
	if result.hour != expectedResult.hour {
		t.Errorf("Expected hour %d, but got %d", expectedResult.hour, result.hour)
	}
	if result.min != expectedResult.min {
		t.Errorf("Expected minute %d, but got %d", expectedResult.min, result.min)
	}
	if result.sec != expectedResult.sec {
		t.Errorf("Expected second %d, but got %d", expectedResult.sec, result.sec)
	}
	if result.nsec != expectedResult.nsec {
		t.Errorf("Expected nanosecond %d, but got %d", expectedResult.nsec, result.nsec)
	}
	if result.loc.String() != expectedResult.loc.String() {
		t.Errorf("Expected location %s, but got %s", expectedResult.loc, result.loc)
	}
}

func TestSubJalaliDuration(t *testing.T) {
	// Test subtracting positive duration from a date
	initialDate := JalaliTime{1399, 2, 28, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	duration := JalaliDuration{1, 2, 3}
	expectedResult := JalaliTime{1397, 12, 25, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	result := initialDate.SubJalaliDuration(duration)
	if result.year != expectedResult.year {
		t.Errorf("Expected year %d, but got %d", expectedResult.year, result.year)
	}
	if result.month != expectedResult.month {
		t.Errorf("Expected month %d, but got %d", expectedResult.month, result.month)
	}
	if result.day != expectedResult.day {
		t.Errorf("Expected day %d, but got %d", expectedResult.day, result.day)
	}
	if result.hour != expectedResult.hour {
		t.Errorf("Expected hour %d, but got %d", expectedResult.hour, result.hour)
	}
	if result.min != expectedResult.min {
		t.Errorf("Expected minute %d, but got %d", expectedResult.min, result.min)
	}
	if result.sec != expectedResult.sec {
		t.Errorf("Expected second %d, but got %d", expectedResult.sec, result.sec)
	}
	if result.nsec != expectedResult.nsec {
		t.Errorf("Expected nanosecond %d, but got %d", expectedResult.nsec, result.nsec)
	}
	if result.loc.String() != expectedResult.loc.String() {
		t.Errorf("Expected location %s, but got %s", expectedResult.loc, result.loc)
	}

	// Test subtracting negative duration from a date
	initialDate = JalaliTime{1400, 5, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	duration = JalaliDuration{-1, -2, -3}
	expectedResult = JalaliTime{1401, 07, 04, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	result = initialDate.SubJalaliDuration(duration)
	if result.year != expectedResult.year {
		t.Errorf("Expected year %d, but got %d", expectedResult.year, result.year)
	}
	if result.month != expectedResult.month {
		t.Errorf("Expected month %d, but got %d", expectedResult.month, result.month)
	}
	if result.day != expectedResult.day {
		t.Errorf("Expected day %d, but got %d", expectedResult.day, result.day)
	}
	if result.hour != expectedResult.hour {
		t.Errorf("Expected hour %d, but got %d", expectedResult.hour, result.hour)
	}
	if result.min != expectedResult.min {
		t.Errorf("Expected minute %d, but got %d", expectedResult.min, result.min)
	}
	if result.sec != expectedResult.sec {
		t.Errorf("Expected second %d, but got %d", expectedResult.sec, result.sec)
	}
	if result.nsec != expectedResult.nsec {
		t.Errorf("Expected nanosecond %d, but got %d", expectedResult.nsec, result.nsec)
	}
	if result.loc.String() != expectedResult.loc.String() {
		t.Errorf("Expected location %s, but got %s", expectedResult.loc, result.loc)
	}

	// Test subtracting zero duration from a date
	initialDate = JalaliTime{1400, 5, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	duration = JalaliDuration{0, 0, 0}
	expectedResult = JalaliTime{1400, 5, 1, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	result = initialDate.SubJalaliDuration(duration)
	if result.year != expectedResult.year {
		t.Errorf("Expected year %d, but got %d", expectedResult.year, result.year)
	}
	if result.month != expectedResult.month {
		t.Errorf("Expected month %d, but got %d", expectedResult.month, result.month)
	}
	if result.day != expectedResult.day {
		t.Errorf("Expected day %d, but got %d", expectedResult.day, result.day)
	}
	if result.hour != expectedResult.hour {
		t.Errorf("Expected hour %d, but got %d", expectedResult.hour, result.hour)
	}
	if result.min != expectedResult.min {
		t.Errorf("Expected minute %d, but got %d", expectedResult.min, result.min)
	}
	if result.sec != expectedResult.sec {
		t.Errorf("Expected second %d, but got %d", expectedResult.sec, result.sec)
	}
	if result.nsec != expectedResult.nsec {
		t.Errorf("Expected nanosecond %d, but got %d", expectedResult.nsec, result.nsec)
	}
	if result.loc.String() != expectedResult.loc.String() {
		t.Errorf("Expected location %s, but got %s", expectedResult.loc, result.loc)
	}

	// Test subtracting duration with days > current month's days
	initialDate = JalaliTime{1399, 2, 28, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	duration = JalaliDuration{0, 0, 32}
	expectedResult = JalaliTime{1399, 01, 27, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	result = initialDate.SubJalaliDuration(duration)
	if result.year != expectedResult.year {
		t.Errorf("Expected year %d, but got %d", expectedResult.year, result.year)
	}
	if result.month != expectedResult.month {
		t.Errorf("Expected month %d, but got %d", expectedResult.month, result.month)
	}
	if result.day != expectedResult.day {
		t.Errorf("Expected day %d, but got %d", expectedResult.day, result.day)
	}
	if result.hour != expectedResult.hour {
		t.Errorf("Expected hour %d, but got %d", expectedResult.hour, result.hour)
	}
	if result.min != expectedResult.min {
		t.Errorf("Expected minute %d, but got %d", expectedResult.min, result.min)
	}
	if result.sec != expectedResult.sec {
		t.Errorf("Expected second %d, but got %d", expectedResult.sec, result.sec)
	}
	if result.nsec != expectedResult.nsec {
		t.Errorf("Expected nanosecond %d, but got %d", expectedResult.nsec, result.nsec)
	}
	if result.loc.String() != expectedResult.loc.String() {
		t.Errorf("Expected location %s, but got %s", expectedResult.loc, result.loc)
	}

	// Test subtracting duration with months > current year's months
	initialDate = JalaliTime{1399, 2, 28, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	duration = JalaliDuration{1, 0, 0}
	expectedResult = JalaliTime{1398, 2, 28, 0, 0, 0, 0, time.FixedZone("IRDT", int(3.5*float64(time.Hour/time.Second)))}
	result = initialDate.SubJalaliDuration(duration)
	if result.year != expectedResult.year {
		t.Errorf("Expected year %d, but got %d", expectedResult.year, result.year)
	}
	if result.month != expectedResult.month {
		t.Errorf("Expected month %d, but got %d", expectedResult.month, result.month)
	}
	if result.day != expectedResult.day {
		t.Errorf("Expected day %d, but got %d", expectedResult.day, result.day)
	}
	if result.hour != expectedResult.hour {
		t.Errorf("Expected hour %d, but got %d", expectedResult.hour, result.hour)
	}
	if result.min != expectedResult.min {
		t.Errorf("Expected minute %d, but got %d", expectedResult.min, result.min)
	}
	if result.sec != expectedResult.sec {
		t.Errorf("Expected second %d, but got %d", expectedResult.sec, result.sec)
	}
	if result.nsec != expectedResult.nsec {
		t.Errorf("Expected nanosecond %d, but got %d", expectedResult.nsec, result.nsec)
	}
	if result.loc.String() != expectedResult.loc.String() {
		t.Errorf("Expected location %s, but got %s", expectedResult.loc, result.loc)
	}
}

func TestJalaliToGregorian(t *testing.T) {
	tests := []struct {
		name   string
		jYear  int
		jMonth Month
		jDay   int
		wantGY int
		wantGM time.Month
		wantGD int
	}{
		{
			name:   "January 1, 1400",
			jYear:  1400,
			jMonth: 1,
			jDay:   1,
			wantGY: 2021,
			wantGM: 3,
			wantGD: 21,
		},
		{
			name:   "February 1, 1400",
			jYear:  1400,
			jMonth: 2,
			jDay:   1,
			wantGY: 2021,
			wantGM: 4,
			wantGD: 21,
		},
		{
			name:   "March 1, 1400",
			jYear:  1400,
			jMonth: 3,
			jDay:   1,
			wantGY: 2021,
			wantGM: 5,
			wantGD: 22,
		},
		{
			name:   "February 29, 1399",
			jYear:  1399,
			jMonth: 12,
			jDay:   10,
			wantGY: 2021,
			wantGM: 2,
			wantGD: 28,
		},
		{
			name:   "December 29, 1399",
			jYear:  1399,
			jMonth: 10,
			jDay:   10,
			wantGY: 2020,
			wantGM: 12,
			wantGD: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gYear, gMonth, gDay := jalaliToGregorian(tt.jYear, tt.jMonth, tt.jDay)
			if gYear != tt.wantGY || gMonth != tt.wantGM || gDay != tt.wantGD {
				t.Errorf("jalaliToGregorian(%d, %d, %d) = (%d, %d, %d); want (%d, %d, %d)",
					tt.jYear, tt.jMonth, tt.jDay, gYear, gMonth, gDay, tt.wantGY, tt.wantGM, tt.wantGD)
			}
		})
	}
}

func TestDaysInMonth(t *testing.T) {
	var tests = []struct {
		year     int
		month    Month
		expected int
	}{
		{1399, Farvardin, 31},
		{1400, Ordibehesht, 31},
		{1401, Khordad, 31},
		{1402, Tir, 31},
		{1403, Mordad, 31},
		{1404, Shahrivar, 31},
		{1405, Mehr, 30},
		{1406, Aban, 30},
		{1407, Azar, 30},
		{1408, Dey, 30},
		{1409, Bahman, 30},
		{1410, Esfand, 29},
	}

	for _, test := range tests {
		if output := daysInMonth(test.year, test.month); output != test.expected {
			t.Errorf("daysInMonth(%d, %v) = %d; expected %d", test.year, test.month, output, test.expected)
		}
	}
}

func TestBoolToInt(t *testing.T) {
	// Test with true input
	result := boolToInt(true)
	expected := 1
	if result != expected {
		t.Errorf("boolToInt(true) returned %d, expected %d", result, expected)
	}

	// Test with false input
	result = boolToInt(false)
	expected = 0
	if result != expected {
		t.Errorf("boolToInt(false) returned %d, expected %d", result, expected)
	}
}

func TestInternalIsLeapJalaliYear(t *testing.T) {
	testCases := []struct {
		year     int
		expected bool
	}{
		{1206, false},
		{1207, false},
		{1208, false},
		{1209, false},
		{1210, true},
		{1211, false},
		{1212, false},
		{1213, false},
		{1214, true},
		{1215, false},
		{1216, false},
		{1217, false},
		{1218, true},
		{1219, false},
		{1220, false},
		{1221, false},
		{1222, true},
		{1223, false},
		{1224, false},
		{1225, false},
		{1226, true},
		{1227, false},
		{1228, false},
		{1229, false},
		{1230, true},
		{1231, false},
		{1232, false},
		{1233, false},
		{1234, true},
		{1235, false},
		{1236, false},
		{1237, false},
		{1238, true},
		{1239, false},
		{1240, false},
		{1241, false},
		{1242, false},
		{1243, true},
		{1244, false},
		{1245, false},
		{1246, false},
		{1247, true},
		{1248, false},
		{1249, false},
		{1250, false},
		{1251, true},
		{1252, false},
		{1253, false},
		{1254, false},
		{1255, true},
		{1256, false},
		{1257, false},
		{1258, false},
		{1259, true},
		{1260, false},
		{1261, false},
		{1262, false},
		{1263, true},
		{1264, false},
		{1265, false},
		{1266, false},
		{1267, true},
		{1268, false},
		{1269, false},
		{1270, false},
		{1271, true},
		{1272, false},
		{1273, false},
		{1274, false},
		{1275, false},
		{1276, true},
		{1277, false},
		{1278, false},
		{1279, false},
		{1280, true},
		{1281, false},
		{1282, false},
		{1283, false},
		{1284, true},
		{1285, false},
		{1286, false},
		{1287, false},
		{1288, true},
		{1289, false},
		{1290, false},
		{1291, false},
		{1292, true},
		{1293, false},
		{1294, false},
		{1295, false},
		{1296, true},
		{1297, false},
		{1298, false},
		{1299, false},
		{1300, true},
		{1301, false},
		{1302, false},
		{1303, false},
		{1304, true},
		{1305, false},
		{1306, false},
		{1307, false},
		{1308, false},
		{1309, true},
		{1310, false},
		{1311, false},
		{1312, false},
		{1313, true},
		{1314, false},
		{1315, false},
		{1316, false},
		{1317, true},
		{1318, false},
		{1319, false},
		{1320, false},
		{1321, true},
		{1322, false},
		{1323, false},
		{1324, false},
		{1325, true},
		{1326, false},
		{1327, false},
		{1328, false},
		{1329, true},
		{1330, false},
		{1331, false},
		{1332, false},
		{1333, true},
		{1334, false},
		{1335, false},
		{1336, false},
		{1337, true},
		{1338, false},
		{1339, false},
		{1340, false},
		{1341, false},
		{1342, true},
		{1343, false},
		{1344, false},
		{1345, false},
		{1346, true},
		{1347, false},
		{1348, false},
		{1349, false},
		{1350, true},
		{1351, false},
		{1352, false},
		{1353, false},
		{1354, true},
		{1355, false},
		{1356, false},
		{1357, false},
		{1358, true},
		{1359, false},
		{1360, false},
		{1361, false},
		{1362, true},
		{1363, false},
		{1364, false},
		{1365, false},
		{1366, true},
		{1367, false},
		{1368, false},
		{1369, false},
		{1370, true},
		{1371, false},
		{1372, false},
		{1373, false},
		{1374, false},
		{1375, true},
		{1376, false},
		{1377, false},
		{1378, false},
		{1379, true},
		{1380, false},
		{1381, false},
		{1382, false},
		{1383, true},
		{1384, false},
		{1385, false},
		{1386, false},
		{1387, true},
		{1388, false},
		{1389, false},
		{1390, false},
		{1391, true},
		{1392, false},
		{1393, false},
		{1394, false},
		{1395, true},
		{1396, false},
		{1397, false},
		{1398, false},
		{1399, true},
		{1400, false},
		{1401, false},
		{1402, false},
		{1403, true},
		{1404, false},
		{1405, false},
		{1406, false},
		{1407, false},
		{1408, true},
		{1409, false},
		{1410, false},
		{1411, false},
		{1412, true},
		{1413, false},
		{1414, false},
		{1415, false},
		{1416, true},
		{1417, false},
		{1418, false},
		{1419, false},
		{1420, true},
		{1421, false},
		{1422, false},
		{1423, false},
		{1424, true},
		{1425, false},
		{1426, false},
		{1427, false},
		{1428, true},
		{1429, false},
		{1430, false},
		{1431, false},
		{1432, true},
		{1433, false},
		{1434, false},
		{1435, false},
		{1436, true},
		{1437, false},
		{1438, false},
		{1439, false},
		{1440, false},
		{1441, true},
		{1442, false},
		{1443, false},
		{1444, false},
		{1445, true},
		{1446, false},
		{1447, false},
		{1448, false},
		{1449, true},
		{1450, false},
		{1451, false},
		{1452, false},
		{1453, true},
		{1454, false},
		{1455, false},
		{1456, false},
		{1457, true},
		{1458, false},
		{1459, false},
		{1460, false},
		{1461, true},
		{1462, false},
		{1463, false},
		{1464, false},
		{1465, true},
		{1466, false},
		{1467, false},
		{1468, false},
		{1469, true},
		{1470, false},
		{1471, false},
		{1472, false},
		{1473, false},
		{1474, true},
		{1475, false},
		{1476, false},
		{1477, false},
		{1478, true},
		{1479, false},
		{1480, false},
		{1481, false},
		{1482, true},
		{1483, false},
		{1484, false},
		{1485, false},
		{1486, true},
		{1487, false},
		{1488, false},
		{1489, false},
		{1490, true},
		{1491, false},
		{1492, false},
		{1493, false},
		{1494, true},
		{1495, false},
	}
	for _, testCase := range testCases {
		result := isLeapJalaliYear(testCase.year)
		if result != testCase.expected {
			t.Errorf("Expected isLeapJalaliYear(%d) to be %v, but got %v", testCase.year, testCase.expected, result)
		}
	}
}
