package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sue struct {
	number int
	attrs  map[string]int
}

func newSue() sue {
	return sue{
		attrs: map[string]int{},
	}
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	sues := []sue{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		splits := strings.SplitN(line, ": ", 2)

		sue := newSue()
		identifiers := strings.Split(splits[0], " ")

		number, err := strconv.Atoi(string(identifiers[1]))
		if err != nil {
			panic(err)
		}
		sue.number = number
		for _, attrString := range strings.Split(splits[1], ", ") {
			attrs := strings.Split(attrString, ": ")
			count, err := strconv.Atoi(attrs[1])
			if err != nil {
				panic(err)
			}
			sue.attrs[attrs[0]] = count
		}

		sues = append(sues, sue)
	}

	knowns := map[string]int{
		"children":    3,
		"cats":        7,
		"samoyeds":    2,
		"pomeranians": 3,
		"akitas":      0,
		"vizslas":     0,
		"goldfish":    5,
		"trees":       3,
		"cars":        2,
		"perfumes":    1,
	}

	for _, sue := range sues {
		good := true
		for attr, count := range sue.attrs {
			switch attr {
			case "cats", "trees":
				good = knowns[attr] < count
			case "pomeranians", "goldfish":
				good = knowns[attr] > count
			default:
				good = knowns[attr] == count
			}
			if !good {
				break
			}
		}
		if good {
			fmt.Printf("sue.number = %+v\n", sue.number)
		}
	}
}
