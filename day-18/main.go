package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	grid := [100][100]bool{}
	row := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		chars := strings.Split(line, "")

		for col, char := range chars {
			grid[row][col] = char == "#"
		}
		row++
	}

	for z := 0; z < 100; z++ {
		newGrid := [100][100]bool{}

		for r, row := range grid {
			for c, lightOn := range row {
				neighborsOn := 0
				if c > 0 {
					if r > 0 && grid[r-1][c-1] {
						neighborsOn++
					}
					if grid[r][c-1] {
						neighborsOn++
					}
					if r < len(row)-1 && grid[r+1][c-1] {
						neighborsOn++
					}
				}

				if r > 0 && grid[r-1][c] {
					neighborsOn++
				}

				if c < len(row)-1 {
					if r > 0 && grid[r-1][c+1] {
						neighborsOn++
					}
					if grid[r][c+1] {
						neighborsOn++
					}
					if r < len(row)-1 && grid[r+1][c+1] {
						neighborsOn++
					}
				}

				if r < len(row)-1 && grid[r+1][c] {
					neighborsOn++
				}

				if lightOn && (neighborsOn == 2 || neighborsOn == 3) {
					newGrid[r][c] = true
				}
				if !lightOn && neighborsOn == 3 {
					newGrid[r][c] = true
				}
			}
		}
		grid = newGrid
	}

	totalOn := 0
	for _, row := range grid {
		for _, lightOn := range row {
			if lightOn {
				totalOn++
			}
		}
	}
	fmt.Printf("totalOn = %+v\n", totalOn)
}
