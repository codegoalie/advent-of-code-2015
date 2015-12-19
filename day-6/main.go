// Package main provides Probably a Fire Hazard

// Because your neighbors keep defeating you in the holiday house decorating
// contest year after year, you've decided to deploy one million lights in a
// 1000x1000 grid.

// Furthermore, because you've been especially nice this year, Santa has mailed
// you instructions on how to display the ideal lighting configuration.

// Lights in your grid are numbered from 0 to 999 in each direction; the lights
// at each corner are at 0,0, 0,999, 999,999, and 999,0. The instructions
// include whether to turn on, turn off, or toggle various inclusive ranges
// given as coordinate pairs. Each coordinate pair represents opposite corners
// of a rectangle, inclusive; a coordinate pair like 0,0 through 2,2 therefore
// refers to 9 lights in a 3x3 square. The lights all start turned off.

// To defeat your neighbors this year, all you have to do is set up your lights
// by doing the instructions Santa sent you in order.

// For example:

// - turn on 0,0 through 999,999 would turn on (or leave on) every light.
// - toggle 0,0 through 999,0 would toggle the first line of 1000 lights,
// 	turning off the ones that were on, and turning on the ones that were off.
// - turn off 499,499 through 500,500 would turn off (or leave off) the middle
// 	four lights.

// After following the instructions, how many lights are lit?

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("configuration.txt")
	if err != nil {
		fmt.Println("Error opening configuration.txt", err)
	}

	var grid [1000][1000]bool

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		words := strings.Split(line, " ")
		if words[0] == "turn" {
			words = words[1:]
		}

		start := strings.Split(words[1], ",")
		end := strings.Split(words[3], ",")
		startX, err := strconv.Atoi(start[0])
		startY, err := strconv.Atoi(start[1])
		endX, err := strconv.Atoi(end[0])
		endY, err := strconv.Atoi(end[1])
		if err != nil {
			panic(err)
		}

		fmt.Println(words[0], startX, endX, startY, endY)
		for i := startX; i <= endX; i++ {
			for ii := startY; ii <= endY; ii++ {
				switch words[0] {
				case "on":
					grid[i][ii] = true
				case "off":
					grid[i][ii] = false
				case "toggle":
					grid[i][ii] = !grid[i][ii]
				}
			}
		}

	}

	lightsOn := 0

	for i := 0; i < 1000; i++ {
		for ii := 0; ii < 1000; ii++ {
			if grid[i][ii] {
				lightsOn++
			}
		}
	}

	fmt.Println("The ideal configuration has", lightsOn, "lights on.")
}
