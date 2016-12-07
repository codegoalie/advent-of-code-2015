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
	destinations []string
}

func newCartogram() cartogram {
	return cartogram{
		distances:    make(map[string]int),
		destinations: []string{},
	}
}

func parseMap(filename string) cartogram {
	var destinations = make(map[string]bool)
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
		globe.distances[parts[2]+parts[0]] = distance
		destinations[parts[0]] = true
		destinations[parts[2]] = true
	}

	for dest := range destinations {
		globe.destinations = append(globe.destinations, dest)
	}

	return globe
}
