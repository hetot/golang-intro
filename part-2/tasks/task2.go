package main

import (
	"fmt"
	"time"
)

func main() {
	dobStr := "05.06.1999" // Replace this date with your birthday
	givenDate, err := time.Parse("02.01.2006", dobStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s is %s\n", givenDate.Format("02-01-2006"), FindWeekday(givenDate))
}

func FindWeekday(date time.Time) (weekday string) {
	switch day := date.Weekday(); day {
	case time.Monday:
		weekday = "Dushanba"
	case time.Tuesday:
		weekday = "Seshanba"
	case time.Wednesday:
		weekday = "Chorshanba"
	case time.Thursday:
		weekday = "Payshanba"
	case time.Friday:
		weekday = "Juma"
	case time.Saturday:
		weekday = "Shanba"
	case time.Sunday:
		weekday = "Yakshanba"
	}
	return
}
