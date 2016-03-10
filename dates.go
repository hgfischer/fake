package fake

import (
	"math/rand"
	"time"
)

// Day generates day of the month
func Day() int {
	return rand.Intn(31) + 1
}

// WeekDay generates name ot the week day
func WeekDay() string {
	return lookup(lang, "weekdays", true)
}

// WeekDayShort generates abbreviated name of the week day
func WeekDayShort() string {
	return lookup(lang, "weekdays_short", true)
}

// WeekdayNum generates number of the day of the week
func WeekdayNum() int {
	return rand.Intn(7) + 1
}

// Month generates month name
func Month() string {
	return lookup(lang, "months", true)
}

// MonthShort generates abbreviated month name
func MonthShort() string {
	return lookup(lang, "months_short", true)
}

// MonthNum generates month number (from 1 to 12)
func MonthNum() int {
	return rand.Intn(12) + 1
}

// Year generates year using the given boundaries
func Year(from, to int) int {
	n := rand.Intn(to-from) + 1
	return from + n
}

// Birthday
func Birthdate(age int) time.Time {
	now := time.Now()
	startWindow := now.AddDate(-1*(age+1), 0, 0)
	endWindow := startWindow.AddDate(1, 0, 0)
	randomUnix := startWindow.Unix() + rand.Int63n(endWindow.Unix()-startWindow.Unix())
	return time.Unix(randomUnix, 0)
}
