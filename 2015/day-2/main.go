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
	totalRibbon := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		dimensions := strings.Split(line, "x")
		maxDimension := 0
		perimeter := 0
		volume := 1
		var areas []int
		for first := 0; first < len(dimensions); first++ {
			sideA, err := strconv.Atoi(dimensions[first])
			if err != nil {
				panic(err)
			}

			if maxDimension < sideA {
				maxDimension = sideA
			}
			perimeter += 2 * sideA
			volume *= sideA

			for second := first + 1; second < len(dimensions); second++ {
				sideB, err := strconv.Atoi(dimensions[second])
				if err != nil {
					panic(err)
				}
				areas = append(areas, sideA*sideB)
			}
		}

		perimeter -= 2 * maxDimension
		totalRibbon += perimeter + volume

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
	fmt.Println("Elves, also get", totalRibbon, "feet of ribbon")
}
