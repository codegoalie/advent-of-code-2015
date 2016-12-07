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

	santa := Point{x: 0, y: 0}
	roboSanta := Point{x: 0, y: 0}
	visits := make(map[Point]int)
	visits[santa] += 1
	visits[roboSanta] += 1
	santasTurn := true

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanRunes)
	for scanner.Scan() {
		if santasTurn {
			santa.move(scanner.Text())
			visits[santa] += 1
			santasTurn = false
		} else {
			roboSanta.move(scanner.Text())
			visits[roboSanta] += 1
			santasTurn = true
		}
	}

	fmt.Println("Together, you've visited", len(visits), "houses")
}

func (point *Point) move(direction string) {
	switch direction {
	case "^":
		point.y += 1
	case ">":
		point.x += 1
	case "v":
		point.y -= 1
	case "<":
		point.x -= 1
	}
}
