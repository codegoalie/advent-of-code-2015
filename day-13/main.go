package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type seat struct {
	s, n string
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	seats := make(map[seat]int)
	subjects := map[string]bool{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		splits := strings.Split(line, " ")

		modification, _ := strconv.Atoi(splits[3])
		if splits[2] == "lose" {
			modification *= -1
		}

		subjects[splits[0]] = true
		seats[seat{s: splits[0], n: splits[10][:len(splits[10])-1]}] = modification
	}

	people := []string{}
	for k := range subjects {
		people = append(people, k)
	}

	var best int
	perms := perm(len(people))
	for _, table := range perms {
		last := len(table) - 1
		var candidate int
		for i := 0; i < len(table); i++ {
			var left int
			var right int

			if i == 0 {
				left = table[last]
			} else {
				left = table[i-1]
			}
			if i == last {
				right = table[0]
			} else {
				right = table[i+1]
			}

			candidate += seats[seat{s: people[table[i]], n: people[left]}]
			candidate += seats[seat{s: people[table[i]], n: people[right]}]
		}

		if candidate > best {
			best = candidate
		}
	}

	fmt.Printf("best = %+v\n", best)
}

func perm(n int) [][]int {
	s := make([]int, n)
	accrue := [][]int{}

	for i := 0; i < n; i++ {
		s[i] = i
	}

	return subPerm(n, s, accrue)

}

func subPerm(n int, s []int, accrue [][]int) [][]int {
	if n == 1 {
		accrue = append(accrue, append([]int(nil), s...))
	} else {
		for i := 0; i < n; i++ {
			accrue = subPerm(n-1, s, accrue)
			if n%2 == 0 {
				s[i], s[n-1] = s[n-1], s[i]
			} else {
				s[0], s[n-1] = s[n-1], s[0]
			}
		}
	}
	return accrue
}
