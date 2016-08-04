package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	seen := "1113122113"
	for i := 0; i < 40; i++ {
		seen = lookAndSee(seen)
		fmt.Printf("i = %+v\n", i)
	}
	fmt.Printf("len(seen) = %+v\n", len(seen))
}

func lookAndSee(in string) string {
	chars := strings.Split(in, "")
	current := chars[0]
	count := 0
	saw := ""

	for _, char := range chars {
		if current != char {
			saw += strconv.Itoa(count) + current
			current = char
			count = 1
		} else {
			count++
		}
	}
	saw += strconv.Itoa(count) + current
	return saw
}
