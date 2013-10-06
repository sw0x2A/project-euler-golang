// https://projecteuler.net/problem=19

package main

import "time"

func p19() int {
	month, year, days, sun := 1, 1900, 1, 0
	for {
		switch month {
		case 4, 6, 9, 11:
			days += 30
		case 1, 3, 5, 7, 8, 10, 12:
			days += 31
		case 2:
			if year%4 == 0 && year%100 != 0 {
				days += 29
			} else {
				days += 28
			}
		}

		if month == 12 {
			month = 1
			year++
		} else {
			month++
		}

		if year >= 1901 && year <= 2000 {
			if days%7 == 0 {
				sun++
			}
		}

		if year >= 2001 {
			break
		}
	}
	return sun
}

func p19weekday() int {
	sun := 0
	for year := 1901; year <= 2000; year++ {
		for month := 1; month <= 12; month++ {
			t := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
			if t.Weekday() == 0 {
				sun++
			}
		}
	}
	return sun
}
