package utils

import "strconv"

func LastDigit(str string) int {
	var last int
	for i := len(str) - 1; i > 0; i-- {
		ch := str[i]
		if ch <= 57 && ch >= 48 {
			last = i
			break
		}
	}
	number, _ := strconv.Atoi(str[last : last+1])
	return number
}
