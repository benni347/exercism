package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout = "1/2/2006 15:04:05"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}
	return parsedDate
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const layout = "January 2, 2006 15:04:05"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}

	currentTime := time.Now()
	return parsedDate.Before(currentTime)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout = "Monday, January 2, 2006 15:04:05"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return false
	}
	hour := parsedDate.Hour()
	if hour >= 12 && hour < 18 {
		return true
	}

	return false
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	const layout = "1/2/2006 15:04:05"

	parsedDate, err := time.Parse(layout, date)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("You have an appointment on %v, %v %v, %v, at %v:%v.", parsedDate.Weekday(), parsedDate.Month(), parsedDate.Day(), parsedDate.Year(), parsedDate.Hour(), parsedDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	currentYear := time.Now().UTC().Year()

	anniversary := time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
	return anniversary
}
