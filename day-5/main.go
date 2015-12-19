// Package main provides differntiation between naughty and nice strings as
// follows:
// --- Day 5: Doesn't He Have Intern-Elves For This? ---
//
// Santa needs help figuring out which strings in his text file are naughty or
// nice.

// A nice string is one with all of the following properties:

// - It contains at least three vowels (aeiou only), like aei, xazegov, or
//   aeiouaeiouaeiou.
// - It contains at least one letter that appears twice in a row, like xx,
//   abcdde (dd), or aabbccdd (aa, bb, cc, or dd).
// - It does not contain the strings ab, cd, pq, or xy, even if they are part
//   of one of the other requirements.
// For example:

// - ugknbfddgicrmopn is nice because it has at least three vowels
//   (u...i...o...), a double letter (...dd...), and none of the disallowed
//   substrings.
// - aaa is nice because it has at least three vowels and a double letter, even
//   though the letters used by different rules overlap.
// - jchzalrnumimnmhp is naughty because it has no double letter.
// - haegwjzuvuyypxyu is naughty because it contains the string xy.
// - dvszwmarrgswjxmb is naughty because it contains only one vowel.

// How many strings are nice?
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("strings.txt")
	if err != nil {
		fmt.Println("Error opening strings.txt", err)
	}

	var nice []string
	var naughty []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if isNice(line) {
			nice = append(nice, line)
		} else {
			naughty = append(naughty, line)
		}
	}

	fmt.Println("Santa, there are", len(nice), "strings this year.")
}

func isNice(s string) bool {
	return hasThreeVowels(s) && hasADup(s) && !hasBadSub(s)
}

func hasThreeVowels(s string) bool {
	vowelCount := 0
	for _, char := range s {
		switch char {
		case 'a', 'e', 'i', 'o', 'u':
			vowelCount++
		}
		if vowelCount > 2 {
			return true
		}
	}
	return false
}

func hasADup(s string) bool {
	prevChar := '_'
	for _, char := range s {
		if char == prevChar {
			return true
		}
		prevChar = char
	}
	return false
}

func hasBadSub(s string) bool {
	for _, sub := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(s, sub) {
			return true
		}
	}
	return false
}
