[![Go Reference](https://pkg.go.dev/badge/github.com/mshafiee/jalali.svg)](https://pkg.go.dev/github.com/mshafiee/jalali)
![Coverage](https://img.shields.io/badge/Coverage-77.2%25-yellow)
[![Go Report Card](https://goreportcard.com/badge/github.com/mshafiee/jalali)](https://goreportcard.com/report/github.com/mshafiee/jalali)
# Jalali
Jalali is a Go package designed for working with the Persian calendar, also known as the Jalali calendar. This comprehensive library offers an easy-to-use API, allowing developers to efficiently convert dates between the Gregorian and Jalali calendars, perform date arithmetic, and format Jalali dates. The package is built to be compatible with Go's standard time package, ensuring seamless integration into any project requiring support for the Jalali calendar system. With its powerful set of tools, this package is an essential resource for applications that deal with the Persian (Jalali) calendar.

## Key features include:
- Conversion between Gregorian and Jalali dates
- Date arithmetic and manipulation
- Support for formatting Jalali dates and times
- Integration with Go's standard time package
- Working with recurring events in the Jalali calendar

Whether you're developing an application with Persian language support or simply need to manipulate Jalali dates, this package offers a powerful and efficient solution for your needs.

## Installation
To use Jalali in your Go project, simply run:

```bash
go get github.com/mshafiee/jalali
```
## Usage
Import the package:

```go
import "github.com/mshafiee/jalali"
```
## Creating Jalali Time
Create a new Jalali time by calling one of the following functions:

```go
// Create a Jalali time for a specific date and time
jalaliTime := jalali.Date(year, month, day, hour, minute, second, nanosecond, location)

// Create a Jalali time for the current date and time
jalaliTime := jalali.Now()

// Convert a Gregorian time to a Jalali time
jalaliTime := jalali.ToJalali(gregorianTime)

// Create a Jalali time from a Unix timestamp
jalaliTime := jalali.JalaliFromTime(unixTimestamp)
```
## Getting Jalali Time Components
You can get the individual components of a Jalali time using the following methods:

```go
year := jalaliTime.Year()
month := jalaliTime.Month()
day := jalaliTime.Day()
hour := jalaliTime.Hour()
minute := jalaliTime.Minute()
second := jalaliTime.Second()
weekday := jalaliTime.Weekday()
```
## Converting Jalali Time to Other Formats
You can convert a Jalali time to a Gregorian time using the ToGregorian method:

```go
gregorianTime := jalaliTime.ToGregorian()
```
You can convert a Jalali time to a Unix timestamp using the Unix or UnixNano methods:

```go
unixTimestamp := jalaliTime.Unix()
unixNanoTimestamp := jalaliTime.UnixNano()
```
## Formatting Jalali Time
You can format a Jalali time using the Format method:


```go
formattedTime := jalaliTime.Format("%Y/%m/%d %T %p")
```
Jalali also provides convenience methods for formatting dates in short and long formats:

```go
formattedShort := jalaliTime.FormatShort()
formattedLong := jalaliTime.FormatLong()
```

## Performing Date Arithmetic
You can add or subtract time from a Jalali time using the Add and Sub methods:

```go
newTime := jalaliTime.Add(24 * time.Hour)
duration := jalaliTime.Sub(otherTime)
```
You can add or subtract years, months, or days from a Jalali time using the AddYears, AddMonths, and AddDays methods:

```go
newTime := jalaliTime.AddYears(1)
newTime := jalaliTime.AddMonths(3)
newTime := jalaliTime.AddDays(7)
```
You can also add or subtract a JalaliDuration using the AddJalaliDuration and SubJalaliDuration methods:

```go
duration := jalali.NewJalaliDuration(1, 2, 3)
newTime := jalaliTime.AddJalaliDuration(duration)
newTime := jalaliTime.SubJalaliDuration(duration)
```
## Working with Recurring Events
Jalali provides a RecurringEvent type that represents an event that occurs on a regular schedule. You can use this type to generate a list of occurrences for an event between two dates:

```go
// Create a new recurring event that occurs every two weeks
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
```
## List of functions
```go
func Date(year int, month Month, day, hour, min, sec, nsec int, loc *time.Location) JalaliTime
func JalaliFromTime(t time.Time) JalaliTime
func ToJalali(t time.Time) JalaliTime
func ParseJalali(layout, value string) (JalaliTime, error)
func Now() JalaliTime
func Tehran() *time.Location
func IRST() *time.Location
func (w Weekday) String() string
func (w Weekday) FaString() string
func (m Month) String() string
func (m Month) FaString() string
func (j JalaliTime) Year() int
func (j JalaliTime) Month() Month
func (j JalaliTime) Day() int
func (j JalaliTime) Hour() int
func (j JalaliTime) Minute() int
func (j JalaliTime) Second() int
func (j JalaliTime) YearDay() int
func (j JalaliTime) Weekday() Weekday
func (j JalaliTime) DaysInMonth() int
func (j JalaliTime) UTC() JalaliTime
func (j JalaliTime) ToGregorian() time.Time
func (j JalaliTime) ToTime() time.Time
func (j JalaliTime) Local() JalaliTime
func (j JalaliTime) In(loc *time.Location) JalaliTime
func (j JalaliTime) Location() *time.Location
func (j JalaliTime) Zone() (name string, offset int)
func (j JalaliTime) Unix() int64
func (j JalaliTime) UnixNano() int64
func (j JalaliTime) Format(layout string) string
func (j JalaliTime) FormatShort() string
func (j JalaliTime) FormatLong() string
func (j JalaliTime) String() string
func (j JalaliTime) DaysBetween(u JalaliTime) int
func (j JalaliTime) After(u JalaliTime) bool
func (j JalaliTime) Before(t JalaliTime) bool
func (j JalaliTime) Equal(other JalaliTime) bool
func (j JalaliTime) IsZero() bool
func (j JalaliTime) IsLeapJalaliYear() bool
func (j JalaliTime) JulianDate() float64
func (j JalaliTime) Add(d time.Duration) JalaliTime
func (j JalaliTime) Sub(u JalaliTime) time.Duration
func (j JalaliTime) AddYears(n int) JalaliTime
func (j JalaliTime) AddMonths(n int) JalaliTime
func (j JalaliTime) AddDays(n int) JalaliTime
func (e RecurringEvent) Occurrences(startDate JalaliTime, endDate JalaliTime) []JalaliTime
func (j JalaliTime) DaysUntil(targetDate JalaliTime) int
func (j JalaliTime) AddDate(years int, months int, days int) JalaliTime
func (j JalaliTime) AddJalaliDuration(d JalaliDuration) JalaliTime
func (j JalaliTime) SubJalaliDuration(d JalaliDuration) JalaliTime
```
