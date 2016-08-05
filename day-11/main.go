package main

import (
	"fmt"
	"regexp"
)

var invalidChars = regexp.MustCompile(`[iol]`)

func main() {
	password := "cqjxjnds"
	password = incrementPassword(password)
	for !validPassword(password) {
		password = incrementPassword(password)
	}
	fmt.Printf("password = %+v\n", password)
	password = incrementPassword(password)
	for !validPassword(password) {
		password = incrementPassword(password)
	}
	fmt.Printf("password = %+v\n", password)
}

func incrementPassword(p string) string {
	var chars []int
	for _, r := range p {
		chars = append(chars, int(r))
	}

	chars[len(chars)-1]++

	for i := len(chars) - 1; i >= 0; i-- {
		if chars[i] > 122 {
			chars[i] = 97
			if i == 0 {
				chars = append([]int{97}, chars...)
			} else {
				chars[i-1]++
			}
		} else {
			break
		}
	}

	str := ""
	for _, point := range chars {
		str += string(rune(point))
	}

	return str
}

func validPassword(p string) bool {
	return hasStraight(p) &&
		hasNoInvalids(p) &&
		hasTwoPairs(p)
}

func hasStraight(p string) bool {
	if len(p) < 3 {
		return false
	}

	count := 0
	var last rune
	for _, code := range p {
		if last+1 == code {
			count++
		} else {
			count = 1
		}
		last = code

		if count > 2 {
			return true
		}
	}

	return false
}

func hasNoInvalids(p string) bool {
	return !invalidChars.MatchString(p)
}

func hasTwoPairs(p string) bool {
	if len(p) < 4 {
		return false
	}

	count := 0
	var last rune
	for _, code := range p {
		if last == code {
			count++
			last = 0
		} else {
			last = code
		}

		if count > 1 {
			return true
		}
	}

	return false
}
