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

// --- Part Two ---

// Realizing the error of his ways, Santa has switched to a better model of
// determining whether a string is naughty or nice. None of the old rules
// apply, as they are all clearly ridiculous.

// Now, a nice string is one with all of the following properties:

// - It contains a pair of any two letters that appears at least twice in the
//   string without overlapping, like xyxy (xy) or aabcdefgaa (aa), but not like
//   aaa (aa, but it overlaps).
// - It contains at least one letter which repeats with exactly one letter
//   between them, like xyx, abcdefeghi (efe), or even aaa.

// For example:

// - `qjhvhtzxzqqjkmpb` is nice because is has a pair that appears twice (qj)
//   and a letter that repeats with exactly one letter between them (zxz).
// - `xxyxx` is nice because it has a pair that appears twice and a letter that
//   repeats with one between, even though the letters used by each rule
//   overlap.
// - `uurcxstgmygtbstg` is naughty because it has a pair (tg) but no repeat
//   with a single letter between them.
// - `ieodomkazucvgmuy` is naughty because it has a repeating letter with one
//   between (odo), but no pair that appears twice.

// How many strings are nice under these new rules?

package main

import (
	"bufio"
	"fmt"
	"math"
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
	niceToo := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if isNice(line) {
			nice = append(nice, line)
		} else {
			naughty = append(naughty, line)
		}

		if isAlsoNice(line) {
			niceToo++
		}
	}

	fmt.Println("Santa, there are", len(nice), "nice strings this year.")
	fmt.Println("Also, there are", niceToo, "nice (v2) strings this year.")
}

func isNice(s string) bool {
	return hasThreeVowels(s) && hasADup(s) && !hasBadSub(s)
}

func isAlsoNice(s string) bool {
	return hasAbA(s) && hasTwoDupes(s)
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

func hasTwoDupes(s string) bool {
	pairs := make(map[string][]int)
	for i := 1; i < len(s); i++ {
		index := s[i-1 : i+1]
		pairs[index] = append(pairs[index], i-1)
	}

	for _, indexes := range pairs {
		if len(indexes) < 2 {
			continue
		}

		if len(indexes) > 2 {
			return true
		}

		diff := math.Abs(float64(indexes[0] - indexes[1]))
		if diff > 1 {
			return true
		}
	}
	return false
}

func hasAbA(s string) bool {
	for i := 2; i < len(s); i++ {
		if s[i-2] == s[i] {
			return true
		}
	}
	return false
}
