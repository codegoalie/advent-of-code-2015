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

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cartogram struct {
	distances    map[string]int
	destinations map[string]bool
}

func main() {
	globe := parseMap("routes.txt")

	size := len(globe.destinations)
	perm(size, func(c []int) {
		fmt.Println(c)
	})

	// for destination := range destinations {
	// 	fmt.Printf("destination = %+v\n", destination)
	// }

	// fmt.Println(places)
	// for route := range GenerateCombinations(places) {
	// 	fmt.Println("Testing", route)
	// 	distance := 0
	// 	for i := 0; i < len(route)-1; i++ {
	// 		d := distances[route[i]+route[i+1]]
	// 		fmt.Println(route[i]+route[i+1], d)
	// 		distance += d
	// 	}
	// 	if distance < shortestDistance {
	// 		shortestDistance = distance
	// 	}
	// }

	// fmt.Println("Santa, the shortest distance you can fly is", shortestDistance)
}

// func GenerateCombinations(destinations []string) <-chan []string {
// 	c := make(chan []string)
// 	length := len(destinations)

// 	go func(c chan []string) {
// 		defer close(c)

// 		AddDestination(c, []string{}, destinations, length)
// 	}(c)

// 	return c
// }

// func AddDestination(c chan []string, combo, destinations []string, length int) {
// 	if length <= 0 {
// 		return
// 	}

// 	var newCombo []string
// 	for _, destination := range destinations {
// 		fmt.Println("Adding", destination)
// 		newCombo = append(combo, destination)
// 		c <- newCombo
// 		AddDestination(c, newCombo, destinations, length-1)
// 	}
// }

func perm(n int, emit func([]int)) {
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i] = i
	}

	emit(s)

	return
}

// func comb(n, m int, emit func([]int)) {
// 	s := make([]int, m)
// 	last := m - 1
// 	var rc func(int, int)
// 	rc = func(i, next int) {
// 		for j := next; j < n; j++ {
// 			s[i] = j
// 			if i == last {
// 				emit(s)
// 			} else {
// 				rc(i+1, j+1)
// 			}
// 		}
// 		return
// 	}
// 	rc(0, 0)
// }

func newCartogram() cartogram {
	return cartogram{
		distances:    make(map[string]int),
		destinations: make(map[string]bool),
	}
}

func parseMap(filename string) cartogram {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	globe := newCartogram()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		parts := strings.Split(line, " ")
		distance, err := strconv.Atoi(parts[4])
		if err != nil {
			panic(err)
		}
		globe.distances[parts[0]+parts[2]] = distance
		globe.destinations[parts[0]] = true
		globe.destinations[parts[2]] = true
	}

	return globe
}
