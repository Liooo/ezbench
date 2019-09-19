package ezbench

import "fmt"

func toOrder(n int) string {
	if n == 1 {
		return "1st"
	} else if n == 2 {
		return "2nd"
	} else if n == 3 {
		return "3rd"
	}
	return fmt.Sprint(n) + "th"
}

func toCount(n int) string {
	if n == 1 {
		return "once"
	} else if n == 2 {
		return "twice"
	}
	return fmt.Sprint(n) + " times"
}
