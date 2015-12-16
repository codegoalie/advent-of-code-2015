package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("gifts.txt")
	if err != nil {
		fmt.Println("Error opening gifts.txt", err)
	}

	totalArea := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		dimensions := strings.Split(line, "x")

		var areas []int
		for first := 0; first < len(dimensions); first++ {
			for second := first + 1; second < len(dimensions); second++ {
				sideA, err := strconv.Atoi(dimensions[first])
				if err != nil {
					panic(err)
				}
				sideB, err := strconv.Atoi(dimensions[second])
				if err != nil {
					panic(err)
				}
				areas = append(areas, sideA*sideB)
			}
		}

		packageArea := 0
		minArea := math.MaxInt64
		for _, area := range areas {
			if area < minArea {
				minArea = area
			}
			packageArea += 2 * area
		}
		packageArea += minArea
		totalArea += packageArea
	}
	fmt.Println("Elves, get", totalArea, "square feet of wrapping paper")
}
