package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x int
	y int
}

func main() {
	file, err := os.Open("directions.txt")
	if err != nil {
		fmt.Println("Error opening directions.txt", err)
	}

	current := Point{x: 0, y: 0}
	visits := make(map[Point]int)
	visits[current] += 1

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		switch scanner.Text() {
		case "^":
			current.y += 1
		case ">":
			current.x += 1
		case "v":
			current.y -= 1
		case "<":
			current.x -= 1
		}

		visits[current] += 1
	}

	fmt.Println("Santa, you've given", len(visits), "different houses presents.")
}
