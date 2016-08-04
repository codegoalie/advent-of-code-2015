// Package main provides our traveling present giver, Santa, with the shortest
// route to deliver all his presents.

// Every year, Santa manages to deliver all of his presents in a single night.

// This year, however, he has some new locations to visit; his elves have
// provided him the distances between every pair of locations. He can start and
// end at any two (different) locations he wants, but he must visit each
// location exactly once. What is the shortest distance he can travel to
// achieve this?

// For example, given the following distances:

// London to Dublin = 464
// London to Belfast = 518
// Dublin to Belfast = 141
// The possible routes are therefore:

// Dublin -> London -> Belfast = 982
// London -> Dublin -> Belfast = 605
// London -> Belfast -> Dublin = 659
// Dublin -> Belfast -> London = 659
// Belfast -> Dublin -> London = 605
// Belfast -> London -> Dublin = 982

// The shortest of these is London -> Dublin -> Belfast = 605, and so the
// answer is 605 in this example.

// What is the distance of the shortest route?
package main

import "fmt"

func main() {
	globe := parseMap("routes.txt")

	size := len(globe.destinations)
	longest := 0
	for _, p := range perm(size) {
		trial := 0
		j := 1
		for i := 0; i < len(p)-1; i++ {
			origin := globe.destinations[p[i]]
			destination := globe.destinations[p[j]]
			trial += globe.distances[origin+destination]
			j++
		}
		fmt.Printf("trial = %+v\n", trial)
		fmt.Printf("longest = %+v\n", longest)
		if trial > longest {
			longest = trial
		}
	}
	fmt.Printf("\n\nAnd the winner is...\n")
	fmt.Printf("longest = %+v\n", longest)
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
