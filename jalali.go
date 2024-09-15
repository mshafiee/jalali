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

// Package jalali provides functionality for working with the Jalali calendar, also known as the Persian calendar.
package jalali

import (
	"errors"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// gregorianDaysInMonth contains the number of days in each month in the Gregorian calendar.
var gregorianDaysInMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// jalaliDaysInMonth contains the number of days in each month in the Jalali calendar.
var jalaliDaysInMonth = []int{31, 31, 31, 31, 31, 31, 30, 30, 30, 30, 30, 29}

// EnJalaliMonthName contains the names of the months in the Jalali calendar in English.
var EnJalaliMonthName = []string{
	"",
	"Farvardin", "Ordibehesht", "Khordad",
	"Tir", "Mordad", "Shahrivar",
	"Mehr", "Aban", "Azar",
	"Dey", "Bahman", "Esfand",
}

// FaJalaliMonthName contains the names of the months in the Jalali calendar in Persian.
var FaJalaliMonthName = []string{
	"",
	"فروردین", "اردیبهشت", "خرداد",
	"تیر", "مرداد", "شهریور",
	"مهر", "آبان", "آذر",
	"دی", "بهمن", "اسفند",
}

// JalaliMonthNumber contains the two-digit numbers of the months in the Jalali calendar.
var JalaliMonthNumber = []string{
	"",
	"01", "02", "03",
	"04", "05", "06",
	"07", "08", "09",
	"10", "11", "12",
}

// FaWeekDays contains the names of the weekdays in Persian.
var FaWeekDays = []string{"یکشنبه", "دوشنبه", "سه‌شنبه", "چهارشنبه", "پنج‌شنبه", "جمعه", "شنبه"}

// EnWeekDays contains the names of the weekdays in English.
var EnWeekDays = []string{"1Shanbeh", "2Shanbeh", "3Shanbeh", "4Shanbeh", "5Shanbeh", "Joomeh", "Shanbeh"}

// Month represents a month in the Jalali calendar.
type Month int

const (
	Farvardin Month = 1 + iota
	Ordibehesht
	Khordad
	Tir
	Mordad
	Shahrivar
	Mehr
	Aban
	Azar
	Dey
	Bahman
	Esfand
)

// Weekday represents a day of the week in the Jalali calendar.
type Weekday int

const (
	Yekshanbe Weekday = iota
	Doshanbe
	Seshanbe
	Chaharshanbe
	Panjshanbe
	Joomeh
	Shanbe
)

// String returns the English name of the weekday.
func (w Weekday) String() string {
	if int(w) < 0 || int(w) > len(EnWeekDays)-1 {
		panic(fmt.Sprintf("invalid weekday value: %v", int(w)))
	}
	return EnWeekDays[w]
}

// FaString returns the Persian name of the weekday.
func (w Weekday) FaString() string {
	if int(w) < 0 || int(w) > len(FaWeekDays)-1 {
		panic(fmt.Sprintf("invalid weekday value: %v", int(w)))
	}
	return FaWeekDays[w]
}

// String returns the English name of the month.
func (m Month) String() string {
	if int(m) < 1 || int(m) > len(EnJalaliMonthName)-1 {
		panic(fmt.Sprintf("invalid month value: %v", int(m)))
	}
	return EnJalaliMonthName[m]
}

// FaString returns the Persian name of the month.
func (m Month) FaString() string {
	if int(m) < 1 || int(m) > len(FaJalaliMonthName)-1 {
		panic(fmt.Sprintf("invalid month value: %v", int(m)))
	}
	return FaJalaliMonthName[m]
}

// JalaliTime represents a date and time in the Jalali calendar.
type JalaliTime struct {
	year  int            // Jalali year
	month Month          // Jalali month (1-12)
	day   int            // Jalali day of month (1-31)
	hour  int            // hour (0-23)
	min   int            // Minute (0-59)
	sec   int            // Second (0-59)
	nsec  int            // Nanosecond (0-999999999)
	loc   *time.Location // JalaliTime zone location
}

// Year returns the year of the Jalali date.
func (j JalaliTime) Year() int {
	return j.year
}

// Month returns the month of the Jalali date.
func (j JalaliTime) Month() Month {
	return j.month
}

// Day returns the day of the month of the Jalali date.
func (j JalaliTime) Day() int {
	return j.day
}

// Hour returns the hour of the Jalali time.
func (j JalaliTime) Hour() int {
	return j.hour
}

// Minute returns the minute of the Jalali time.
func (j JalaliTime) Minute() int {
	return j.min
}

// Second returns the second of the Jalali time.
func (j JalaliTime) Second() int {
	return j.sec
}

// YearDay returns the day of the year of the Jalali date.
func (j JalaliTime) YearDay() int {
	// Calculate the number of days from the start of the Jalali year to the date
	yearStart, monthStart, dayStart := jalaliToGregorian(j.year, Farvardin, 1)

	jd := j.JulianDate()
	jdn, _ := julianDayNumber(yearStart, monthStart, dayStart, 0, 0, 0, 0)
	yearDay := int(jd - jdn + 1)
	return yearDay
}

// Weekday returns the day of the week of the Jalali date.
func (j JalaliTime) Weekday() Weekday {
	// Convert the Jalali date to a Gregorian date
	gYear, gMonth, gDay := jalaliToGregorian(j.year, j.month, j.day)

	// Create a time.Time value for the Gregorian date
	gDate := time.Date(gYear, gMonth, gDay, j.hour, j.min, j.sec, j.nsec, j.loc)

	// Get the day of the week as a time.Weekday value
	weekday := gDate.Weekday()

	// Convert the time.Weekday value to our custom Weekday type
	switch weekday {
	case time.Sunday:
		return Yekshanbe
	case time.Monday:
		return Doshanbe
	case time.Tuesday:
		return Seshanbe
	case time.Wednesday:
		return Chaharshanbe
	case time.Thursday:
		return Panjshanbe
	case time.Friday:
		return Joomeh
	case time.Saturday:
		return Shanbe
	default:
		panic(fmt.Sprintf("invalid weekday value: %v", weekday))
	}
}

// Date returns a new JalaliTime value representing the given date and time.
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) JalaliTime {
	// Validate the input values
	if year < 1 || year > 9999 {
		panic(fmt.Sprintf("year out of range: %d", year))
	}
	if month < Farvardin || month > Esfand {
		panic(fmt.Sprintf("invalid month: %v", month))
	}
	if day < 1 || day > daysInMonth(year, month) {
		panic(fmt.Sprintf("day out of range: %d", day))
	}
	if hour < 0 || hour > 23 {
		panic(fmt.Sprintf("hour out of range: %d", hour))
	}
	if min < 0 || min > 59 {
		panic(fmt.Sprintf("minute out of range: %d", min))
	}
	if sec < 0 || sec > 59 {
		panic(fmt.Sprintf("second out of range: %d", sec))
	}
	if nsec < 0 || nsec > 999999999 {
		panic(fmt.Sprintf("nanosecond out of range: %d", nsec))
	}

	// Create a new JalaliTime value
	jTime := JalaliTime{
		year:  year,
		month: month,
		day:   day,
		hour:  hour,
		min:   min,
		sec:   sec,
		nsec:  nsec,
		loc:   loc,
	}

	return jTime
}

// DaysInMonth returns the number of days in the month of the JalaliTime.
func (j JalaliTime) DaysInMonth() int {
	return daysInMonth(j.year, j.month)
}

// Now returns the current JalaliTime.
func Now() JalaliTime {
	// Implement code to convert the current time to Jalali date
	return ToJalali(time.Now())
}

// UTC returns the JalaliTime in UTC time zone.
func (j JalaliTime) UTC() JalaliTime {
	gTime := j.ToGregorian().UTC()
	return ToJalali(gTime)
}

// ToJalali converts a time.Time value to JalaliTime.
func ToJalali(t time.Time) JalaliTime {
	// Convert a Gregorian time to a JalaliTime
	// Calculate the Jalali date
	jy, jm, jd := gregorianToJalali(t.Year(), t.Month(), t.Day())

	// Return the JalaliTime
	return JalaliTime{
		year:  jy,
		month: Month(jm),
		day:   jd,
		hour:  t.Hour(),
		min:   t.Minute(),
		sec:   t.Second(),
		nsec:  t.Nanosecond(),
		loc:   t.Location(),
	}
}

// ToGregorian converts a JalaliTime to time.Time{} value.
func (j JalaliTime) ToGregorian() time.Time {
	// Convert a JalaliTime to a Gregorian time
	gYear, gMonth, gDay := jalaliToGregorian(j.year, j.month, j.day)
	if j.loc == nil {
		j.loc = time.Local
	}
	return time.Date(gYear, gMonth, gDay, j.hour, j.min, j.sec, j.nsec, j.loc)
}

func (j JalaliTime) ToTime() time.Time {
	// Convert a JalaliTime to time.Time{}
	return j.ToGregorian()
}

// Local returns the JalaliTime in local time zone.
func (j JalaliTime) Local() JalaliTime {
	return JalaliTime{
		year:  j.year,
		month: j.month,
		day:   j.day,
		hour:  j.hour,
		min:   j.min,
		sec:   j.sec,
		nsec:  j.nsec,
		loc:   time.Local,
	}
}

// In returns the JalaliTime in the specified time zone.
func (j JalaliTime) In(loc *time.Location) JalaliTime {
	gTime := j.ToGregorian().In(loc)
	return ToJalali(gTime)
}

// Location returns the time zone of the JalaliTime.
func (j JalaliTime) Location() *time.Location {
	if j.loc == nil {
		return time.Local
	}
	return j.loc
}

// Zone returns the time zone abbreviation and offset from UTC in seconds.
func (j JalaliTime) Zone() (name string, offset int) {
	// Convert the Jalali date to a Gregorian date
	gTime := j.ToGregorian()

	// Get the time zone abbreviation and offset in seconds
	name, offset = gTime.Zone()

	return name, offset
}

// Unix returns the number of seconds elapsed since January 1, 1970 UTC to the given JalaliTime
// value. It converts the Jalali date to a Gregorian date and creates a time.Time value for the
// Gregorian date and time, and returns its Unix timestamp as an int64 value.
func (j JalaliTime) Unix() int64 {
	// Convert the Jalali date to a Gregorian date
	gYear, gMonth, gDay := jalaliToGregorian(j.year, j.month, j.day)

	// Create a time.Time value for the Gregorian date and time
	gDate := time.Date(gYear, time.Month(gMonth), gDay, j.hour, j.min, j.sec, j.nsec, j.loc)

	// Get the Unix timestamp as an int64 value
	return gDate.Unix()
}

// UnixNano returns the number of nanoseconds elapsed since January 1, 1970 UTC to the given
// JalaliTime value. It first converts the Jalali time to a UTC time.Time value, and then
// calculates the number of nanoseconds since the Unix epoch (January 1, 1970 UTC) to the UTC time.
func (j JalaliTime) UnixNano() int64 {
	// Convert Jalali time to UTC time
	utc := j.ToGregorian().UTC()

	// Calculate the number of nanoseconds since Unix epoch
	return utc.UnixNano()
}

// Format returns a string representing the Jalali time formatted according to the layout string.
// The layout string uses format specifiers similar to strftime, starting with % followed by a letter.
// Supported specifiers:
//
//	%Y: year as a 4-digit number (e.g., 1402)
//	%y: year as a 2-digit number (e.g., 02)
//	%m: month as a 2-digit number (01-12)
//	%B: full month name in Persian
//	%b: abbreviated month name in Persian
//	%d: day of the month as a 2-digit number (01-31)
//	%H: hour (00-23)
//	%M: minute (00-59)
//	%S: second (00-59)
//	%p: "AM" or "PM" in Persian ("صبح" or "عصر")
//	%w: weekday name in Persian
//	%z: time zone offset as ±hhmm
//	%Z: time zone name
//	%R: 24-hour time in the format "HH:MM"
//	%T: time in the format "HH:MM:SS"
//	%%: percent sign
func (j JalaliTime) Format(layout string) string {
	var builder strings.Builder
	length := len(layout)
	i := 0

	for i < length {
		if layout[i] == '%' && i+1 < length {
			specifier := layout[i : i+2]
			switch specifier {
			case "%n":
				builder.WriteByte('\n')
			case "%%":
				builder.WriteByte('%')
			case "%Y":
				builder.WriteString(fmt.Sprintf("%04d", j.year))
			case "%y":
				builder.WriteString(fmt.Sprintf("%02d", j.year%100))
			case "%m":
				builder.WriteString(fmt.Sprintf("%02d", j.month))
			case "%B":
				builder.WriteString(FaJalaliMonthName[j.month])
			case "%b":
				// Abbreviated month name, take first three characters
				if len(FaJalaliMonthName[j.month]) >= 3 {
					builder.WriteString(FaJalaliMonthName[j.month][:3])
				} else {
					builder.WriteString(FaJalaliMonthName[j.month])
				}
			case "%d":
				builder.WriteString(fmt.Sprintf("%02d", j.day))
			case "%H":
				builder.WriteString(fmt.Sprintf("%02d", j.hour))
			case "%M":
				builder.WriteString(fmt.Sprintf("%02d", j.min))
			case "%S":
				builder.WriteString(fmt.Sprintf("%02d", j.sec))
			case "%p":
				if j.hour < 12 {
					builder.WriteString("صبح") // AM in Persian
				} else {
					builder.WriteString("عصر") // PM in Persian
				}
			case "%w":
				builder.WriteString(FaWeekDays[j.Weekday()])
			case "%z":
				_, offset := j.Zone()
				sign := "+"
				if offset < 0 {
					sign = "-"
					offset = -offset
				}
				hours := offset / 3600
				minutes := (offset % 3600) / 60
				builder.WriteString(fmt.Sprintf("%s%02d%02d", sign, hours, minutes))
			case "%Z":
				if j.loc != nil {
					builder.WriteString(j.loc.String())
				}
			case "%R":
				builder.WriteString(fmt.Sprintf("%02d:%02d", j.hour, j.min))
			case "%T":
				builder.WriteString(fmt.Sprintf("%02d:%02d:%02d", j.hour, j.min, j.sec))
			default:
				// Unknown specifier, write as-is
				builder.WriteString(specifier)
			}
			i += 2
		} else {
			builder.WriteByte(layout[i])
			i++
		}
	}

	return builder.String()
}

// FormatShort returns the JalaliTime formatted as a short string in the format "YYYY/MM/DD".
func (j JalaliTime) FormatShort() string {
	return j.Format("%Y/%m/%d")
}

// FormatLong returns the JalaliTime formatted as a long string in the format "DD MonthName YYYY".
// The month name is written in Persian.
func (j JalaliTime) FormatLong() string {
	return j.Format("%d %B %Y")
}

// String returns the JalaliTime formatted as a string in the format "YYYY/MM/DD HH:MM:SS".
func (j JalaliTime) String() string {
	return j.Format("%Y/%m/%d %T")
}

// DaysBetween returns the number of days between two JalaliTime values.
func (j JalaliTime) DaysBetween(u JalaliTime) int {
	// Swap t1 and t2 if t2 is earlier than t1
	if u.Before(j) {
		j, u = u, j
	}

	unix1 := j.Unix()
	unix2 := u.Unix()
	days := int((unix2 - unix1) / 86400)
	return days
}

// After reports whether the time instant t is after u.
func (j JalaliTime) After(u JalaliTime) bool {
	return j.UnixNano() > u.UnixNano()
}

// Before reports whether the time instant t is before u.
func (j JalaliTime) Before(u JalaliTime) bool {
	return j.UnixNano() < u.UnixNano()
}

// Equal reports whether t and u represent the same time instant.
func (j JalaliTime) Equal(u JalaliTime) bool {
	return j.year == u.year &&
		j.month == u.month &&
		j.day == u.day &&
		j.hour == u.hour &&
		j.min == u.min &&
		j.sec == u.sec &&
		j.nsec == u.nsec &&
		j.loc.String() == u.loc.String()
}

// IsZero returns true if the JalaliTime value is equal to the zero value.
func (j JalaliTime) IsZero() bool {
	return j.year == 0 && j.month == 0 && j.day == 0 && j.hour == 0 && j.min == 0 && j.sec == 0 && j.nsec == 0
}

// IsLeapJalaliYear returns true if the year of the JalaliTime is a leap year in the Jalali calendar,
// and false otherwise.
func (j JalaliTime) IsLeapJalaliYear() bool {
	return isLeapJalaliYear(j.year)
}

// JulianDate returns the Julian date for the current
// Jalali date and time. The Julian day number is a count of the number of days
// elapsed since noon on January 1, 4713 BCE (Julian calendar). This function
// first converts the Jalali date and time to a Gregorian date and time, and then
// performs the necessary calculations to obtain the Julian day number, including
// adding the fraction of the day to the whole number.
func (j JalaliTime) JulianDate() float64 {
	// Convert the Jalali date to a Gregorian date
	gDate := j.ToGregorian()

	jdn, err := julianDayNumber(gDate.Year(), gDate.Month(), gDate.Day(), gDate.Hour(), gDate.Minute(), gDate.Second(), gDate.Nanosecond())
	if err != nil {
		fmt.Println(err)
	}
	// Calculate the Julian day number for the Gregorian date
	return jdn
}

func julianDayNumber(year int, month time.Month, day, hour, minute, second, nanosecond int) (float64, error) {
	if year < -4712 {
		return 0, fmt.Errorf("year %d is too early (minimum is -4712)", year)
	}
	if month < time.January || month > time.December {
		return 0, fmt.Errorf("invalid month %v", month)
	}
	if day < 1 || day > 31 {
		return 0, fmt.Errorf("invalid day %d (must be between 1 and 31)", day)
	}
	if hour < 0 || hour > 23 {
		return 0, fmt.Errorf("invalid hour %d (must be between 0 and 23)", hour)
	}
	if minute < 0 || minute > 59 {
		return 0, fmt.Errorf("invalid minute %d (must be between 0 and 59)", minute)
	}
	if second < 0 || second > 59 {
		return 0, fmt.Errorf("invalid second %d (must be between 0 and 59)", second)
	}
	if nanosecond < 0 || nanosecond > 999999999 {
		return 0, fmt.Errorf("invalid nanosecond %d (must be between 0 and 999999999)", nanosecond)
	}

	if month <= 2 {
		year -= 1
		month += 12
	}
	a := math.Floor(float64(year) / 100)
	b := 2 - a + math.Floor(a/4)
	jdn := math.Floor(365.25*(float64(year)+4716)) + math.Floor(30.6001*(float64(month)+1)) + float64(day) + b - 1524.5

	fraction := float64(hour)/24 + float64(minute)/(24*60) + float64(second)/(24*60*60) + float64(nanosecond)/(24*60*60*1e9)
	jd := jdn + fraction

	return jd, nil
}

// Add method adds the specified duration to the current Jalali
// date and time. It first converts the Jalali time to a Gregorian time, adds
// the duration to the Gregorian time, and then converts the resulting Gregorian
// time back to a Jalali time.
func (j JalaliTime) Add(d time.Duration) JalaliTime {
	// Convert the JalaliTime to a time.Time value
	gTime := j.ToGregorian()

	// Add the duration to the time.Time value
	newGTime := gTime.Add(d)

	// Convert the resulting time.Time value to a JalaliTime value
	newJTime := ToJalali(newGTime)

	return newJTime
}

// Sub returns the duration between the current Jalali date and time and
// the specified Jalali date and time. It first converts both Jalali times to
// equivalent time.Time values, then calculates the duration between them
// using the Sub method of time.Time.
func (j JalaliTime) Sub(u JalaliTime) time.Duration {
	// Convert both JalaliTime values to time.Time values
	t1 := j.ToGregorian()
	t2 := u.ToGregorian()

	// Calculate the duration between the two time values
	duration := t1.Sub(t2)

	return duration
}

// AddYears adds n years to the JalaliTime value j
func (j JalaliTime) AddYears(n int) JalaliTime {
	// Calculate the new year
	newYear := j.year + n

	// If the new year is before year 1, return the zero value
	if newYear < 1 {
		return JalaliTime{}
	}

	// If the new year is a leap year and the current month is Esfand, adjust the day
	if j.IsLeapJalaliYear() && j.month == 12 && j.day == 30 {
		j.day = 29
	}

	// Create the new JalaliTime value
	newTime := JalaliTime{
		year:  newYear,
		month: j.month,
		day:   j.day,
		hour:  j.hour,
		min:   j.min,
		sec:   j.sec,
		nsec:  j.nsec,
		loc:   j.loc,
	}

	return newTime
}

// AddMonths adds n months to the JalaliTime value j
func (j JalaliTime) AddMonths(n int) JalaliTime {
	// Calculate the number of years and months to add
	years := n / 12
	months := n % 12

	// Calculate the new year and month values
	newYear := j.year + years
	newMonth := int(j.month) + months

	// Handle rollover into the next year
	if newMonth > 12 {
		newYear += 1
		newMonth -= 12
	}

	// If the new month has fewer days than the current day, adjust the day
	if j.day > daysInMonth(newYear, Month(newMonth)) {
		j.day = daysInMonth(newYear, Month(newMonth))
	}

	// Return a new JalaliTime value with the updated fields
	return JalaliTime{
		year:  newYear,
		month: Month(newMonth),
		day:   j.day,
		hour:  j.hour,
		min:   j.min,
		sec:   j.sec,
		nsec:  j.nsec,
		loc:   j.loc,
	}
}

// AddDays adds n days to the JalaliTime value j
func (j JalaliTime) AddDays(n int) JalaliTime {
	// Convert the Jalali date to a Gregorian date
	gYear, gMonth, gDay := jalaliToGregorian(j.year, j.month, j.day)

	// Create a time.Time value for the Gregorian date and time
	gDate := time.Date(gYear, time.Month(gMonth), gDay, j.hour, j.min, j.sec, j.nsec, j.loc)

	// Add the number of days to the time.Time value
	newGDate := gDate.AddDate(0, 0, n)

	// Convert the resulting time.Time value to a JalaliTime value
	newJTime := ToJalali(newGDate)

	return newJTime
}

// RecurringEvent represents a struct for a recurring event with a name, start time, end time and frequency.
type RecurringEvent struct {
	StartTime JalaliTime
	EndTime   JalaliTime
	Frequency time.Duration
}

// Occurrences method returns a slice of JalaliTime occurrences between the given start and end dates.
func (e RecurringEvent) Occurrences(startDate JalaliTime, endDate JalaliTime) []JalaliTime {
	var occurrences []JalaliTime
	currentDate := startDate

	// Ensure that the event starts before the given end date
	if e.StartTime.After(endDate) {
		return occurrences
	}

	// Adjust the current date to the first occurrence after or on the start date
	offset := e.StartTime.DaysUntil(currentDate)
	if offset > 0 {
		currentDate = e.StartTime.AddDays(offset / int(e.Frequency) * int(e.Frequency))
	}

	// Loop through each occurrence in the date range
	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		// Check if the current date falls within the event's start and end times
		if (currentDate.After(e.StartTime) || currentDate.Equal(e.StartTime)) &&
			(e.EndTime.IsZero() || currentDate.Before(e.EndTime) || currentDate.Equal(e.EndTime)) {
			occurrences = append(occurrences, currentDate)
		}

		// Increment the current date by the event's frequency
		currentDate = currentDate.Add(e.Frequency)
	}

	return occurrences
}

// DaysUntil calculates the number of days between the current date and the target date.
func (j JalaliTime) DaysUntil(targetDate JalaliTime) int {
	jd1 := j.JulianDate()
	jd2 := targetDate.JulianDate()

	return int(jd2 - jd1)
}

// ParseInLocation function takes a layout and a value as input and returns a JalaliTime and an error.
// The layout parameter defines the format of the value parameter in a similar way to the standard
// library's time.Parse function.
// The function uses regular expressions to match the format defined by the layout and extract the relevant values.
// It then validates the Jalali date and returns a JalaliTime object with the parsed values.
// If the value cannot be parsed or the Jalali date is invalid, an error is returned.
func ParseInLocation(layout, value string, loc *time.Location) (JalaliTime, error) {
	var year, month, day, hour, min, sec int

	re := regexp.MustCompile(`%([YymdHMS])`)
	layout = re.ReplaceAllStringFunc(layout, func(match string) string {
		return `(\d+)`
	})

	re = regexp.MustCompile(layout)
	matches := re.FindStringSubmatch(value)

	if len(matches) < 7 {
		return JalaliTime{}, errors.New("unable to parse value using the provided layout")
	}

	year, _ = strconv.Atoi(matches[1])
	month, _ = strconv.Atoi(matches[2])
	day, _ = strconv.Atoi(matches[3])
	hour, _ = strconv.Atoi(matches[4])
	min, _ = strconv.Atoi(matches[5])
	sec, _ = strconv.Atoi(matches[6])

	if !isValidJalaliDate(year, month, day) {
		return JalaliTime{}, fmt.Errorf("invalid Jalali date: %d/%02d/%02d", year, month, day)
	}

	return JalaliTime{
		year:  year,
		month: Month(month),
		day:   day,
		hour:  hour,
		min:   min,
		sec:   sec,
		nsec:  0,
		loc:   loc,
	}, nil
}

func Parse(layout, value string) (JalaliTime, error) {
	return ParseInLocation(layout, value, time.Local)
}

// AddDate adds the specified years, months, and days to the JalaliTime object and returns the updated JalaliTime.
func (j JalaliTime) AddDate(years int, months int, days int) JalaliTime {
	return j.AddYears(years).AddMonths(months).AddDays(days)
}

// JalaliFromTime takes a time.Time argument and returns a JalaliTime value. The purpose of
// this function is to convert a given Gregorian date and time to the corresponding Jalali date and time.
func JalaliFromTime(t time.Time) JalaliTime {
	jyear, jmonth, jday := gregorianToJalali(t.Year(), t.Month(), t.Day())
	// Create a new JalaliTime value with the calculated date and time components
	return JalaliTime{
		year:  jyear,
		month: jmonth,
		day:   jday,
		hour:  t.Hour(),
		min:   t.Minute(),
		sec:   t.Second(),
		nsec:  t.Nanosecond(),
		loc:   t.Location(),
	}
}

// Tehran returns the *time.Location representing the Iran Standard JalaliTime (IRST),
// which is the time zone used in Tehran, Iran.
func Tehran() *time.Location {
	loc, err := time.LoadLocation("Asia/Tehran")
	if err != nil {
		panic(err)
	}
	return loc
}

// IRST returns the *time.Location representing the Iran Standard JalaliTime (IRST)
func IRST() *time.Location {
	return Tehran()
}

// gregorianToJalali converts a Gregorian date (year, month, and day) to a Jalali date (year, month, and day)
func gregorianToJalali(gYear int, gMonth time.Month, gDay int) (jYear int, jMonth Month, jDay int) {
	// Convert the Gregorian year, month, and day to the day number
	// since 1 January 1 AD (Gregorian) using the formula from the
	// book "Calendrical Calculations" by Dershowitz and Reingold.
	gy := gYear - 1600
	gm := gMonth - 1
	gd := gDay - 1

	gDayNo := 365*gy + (gy+3)/4 - (gy+99)/100 + (gy+399)/400
	for i := 0; i < int(gm); i++ {
		gDayNo += gregorianDaysInMonth[i]
	}
	if gm > 1 && ((gy%4 == 0 && gy%100 != 0) || (gy%400 == 0)) {
		gDayNo++ // leap and after Feb
	}
	gDayNo += gd

	// Convert the day number to the Jalali year, month, and day
	// using the inverse of the formula from the book "Calendrical
	// Calculations" by Dershowitz and Reingold.
	jDayNo := gDayNo - 79
	jNp := jDayNo / 12053
	jDayNo %= 12053

	jYear = 979 + 33*jNp + 4*(jDayNo/1461)
	jDayNo %= 1461

	if jDayNo >= 366 {
		jYear += (jDayNo - 1) / 365
		jDayNo = (jDayNo - 1) % 365
	}

	var i int
	for i = 0; i < 11 && jDayNo >= jalaliDaysInMonth[i]; i++ {
		jDayNo -= jalaliDaysInMonth[i]
	}
	jMonth = Month(i + 1)
	jDay = jDayNo + 1

	return jYear, jMonth, jDay
}

type JalaliDuration struct {
	Years  int
	Months int
	Days   int
}

func (j JalaliTime) AddJalaliDuration(d JalaliDuration) JalaliTime {
	newYear := j.year + d.Years
	newMonth := int(j.month) + d.Months
	newDay := j.day + d.Days

	for newMonth > 12 {
		newYear++
		newMonth -= 12
	}

	for newMonth < 1 {
		newYear--
		newMonth += 12
	}

	maxDays := daysInMonth(newYear, Month(newMonth))
	for newDay > maxDays {
		newDay -= maxDays
		newMonth++

		if newMonth > 12 {
			newYear++
			newMonth = 1
		}

		maxDays = daysInMonth(newYear, Month(newMonth))
	}

	for newDay < 1 {
		newMonth--
		if newMonth < 1 {
			newYear--
			newMonth = 12
		}

		maxDays = daysInMonth(newYear, Month(newMonth))
		newDay += maxDays
	}

	return JalaliTime{
		year:  newYear,
		month: Month(newMonth),
		day:   newDay,
		hour:  j.hour,
		min:   j.min,
		sec:   j.sec,
		nsec:  j.nsec,
		loc:   j.loc,
	}
}

func (j JalaliTime) SubJalaliDuration(d JalaliDuration) JalaliTime {
	negativeDuration := JalaliDuration{
		Years:  -d.Years,
		Months: -d.Months,
		Days:   -d.Days,
	}
	return j.AddJalaliDuration(negativeDuration)
}

// jalaliToGregorian that takes in three parameters: jYear (an integer representing the Jalali year),
// jMonth (a value of type Month representing the Jalali month), and jDay (an integer representing the Jalali day).
// The function returns three values: gYear (an integer representing the Gregorian year),
// gMonth (a value of type time.Month representing the Gregorian month), and gDay (an integer representing the Gregorian day).
func jalaliToGregorian(jYear int, jMonth Month, jDay int) (gYear int, gMonth time.Month, gDay int) {
	jy := jYear - 979
	jm := int(jMonth - 1)
	jd := jDay - 1

	jDayNo := 365*jy + (jy/33)*8 + (jy%33+3)/4
	for i := 0; i < jm; i++ {
		jDayNo += jalaliDaysInMonth[i]
	}
	jDayNo += jd

	// Convert the Jalali day number to the Gregorian year, month,
	// and day using the inverse of the formula from the book
	// "Calendrical Calculations" by Dershowitz and Reingold.
	gDayNo := jDayNo + 79
	gy := 1600 + 400*(gDayNo/146097)
	gDayNo %= 146097

	leap := 1
	if gDayNo >= 36525 {
		gDayNo--
		gy += 100 * (gDayNo / 36524)
		gDayNo = gDayNo % 36524

		if gDayNo >= 365 {
			gDayNo++
		} else {
			leap = 0
		}
	}

	gy += 4 * (gDayNo / 1461)
	gDayNo %= 1461

	if gDayNo >= 366 {
		leap = 0
		gDayNo--
		gy += gDayNo / 365
		gDayNo = gDayNo % 365
	}

	var i int
	for i = 0; gDayNo >= gregorianDaysInMonth[i]+boolToInt(i == 1 && leap != 0); i++ {
		gDayNo -= gregorianDaysInMonth[i] + boolToInt(i == 1 && leap != 0)
	}
	gMonth = time.Month(i + 1)
	gDay = gDayNo + 1

	return gy, gMonth, gDay
}

// daysInMonth returns the number of days in the Jalali month for the given year.
func daysInMonth(year int, month Month) int {
	if month < Farvardin || month > Esfand {
		return 0
	}

	// Calculate the number of days in the month
	if month <= Shahrivar {
		return 31
	} else if month <= Bahman {
		return 30
	} else {
		if isLeapJalaliYear(year) {
			return 30
		} else {
			return 29
		}
	}
}

// boolToInt  takes a boolean input and returns an integer output.
func boolToInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

// isLeapJalaliYear returns true if the given Jalali year is a leap year
func isLeapJalaliYear(year int) bool {
	remainder := year % 33
	leapYearsRemainders := map[int]bool{
		1:  true,
		5:  true,
		9:  true,
		13: true,
		17: true,
		22: true,
		26: true,
		30: true,
	}
	_, isLeap := leapYearsRemainders[remainder]
	return isLeap
}

// isValidJalaliDate checks whether the given year, month, and day constitute a valid Jalali date or not.
func isValidJalaliDate(year, month, day int) bool {
	// Check if the year, month, or day values are out of bounds or invalid.
	if year < 1 || month < 1 || month > 12 || day < 1 {
		// If any of the values are out of bounds or invalid, return false.
		return false
	}

	// Check if the given day is greater than the number of days in the given month.
	if day > daysInMonth(year, Month(month)) {
		// If the given day is greater than the number of days in the given month, return false.
		return false
	}

	// If none of the above conditions were met, return true.
	return true
}
