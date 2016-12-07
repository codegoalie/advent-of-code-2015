package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

const input = `R3, R1, R4, L4, R3, R1, R1, L3, L5, L5, L3, R1, R4, L2, L1, R3, L3, R2, R1, R1, L5, L2, L1, R2, L4, R1, L2, L4, R2, R2, L2, L4, L3, R1, R4, R3, L1, R1, L5, R4, L2, R185, L2, R4, R49, L3, L4, R5, R1, R1, L1, L1, R2, L1, L4, R4, R5, R4, L3, L5, R1, R71, L1, R1, R186, L5, L2, R5, R4, R1, L5, L2, R3, R2, R5, R5, R4, R1, R4, R2, L1, R4, L1, L4, L5, L4, R4, R5, R1, L2, L4, L1, L5, L3, L5, R2, L5, R4, L4, R3, R3, R1, R4, L1, L2, R2, L1, R4, R2, R2, R5, R2, R5, L1, R1, L4, R5, R4, R2, R4, L5, R3, R2, R5, R3, L3, L5, L4, L3, L2, L2, R3, R2, L1, L1, L5, R1, L3, R3, R4, R5, L3, L5, R1, L3, L5, L5, L2, R1, L3, L1, L3, R4, L1, R3, L2, L2, R3, R3, R4, R4, R1, L4, R1, L5`

// const input = `R3, R1, R4, L4`

func main() {
	facing, x, y := "N", 0, 0
	visits := make(map[string]bool)
	splits := strings.Split(input, ", ")
	for _, instruction := range splits {
		fmt.Printf("%s - ", instruction)
		direction := instruction[:1]
		blocks, err := strconv.Atoi(instruction[1:len(instruction)])
		if err != nil {
			log.Fatal(err)
		}
		facing = newFacing(facing, direction)
		switch facing {
		case "N":
			y += blocks
		case "E":
			x += blocks
		case "S":
			y -= blocks
		case "W":
			x -= blocks
		default:
			log.Fatalf("Tried to walk %s", facing)
		}
		fmt.Printf("You are at %d, %d facing %s\n", x, y, facing)
		key := fmt.Sprintf("%d,%d", x, y)
		if visits[key] {
			fmt.Printf("visits = %+v\n", visits)
			fmt.Printf("You are at %d, %d again facing %s\n", x, y, facing)
			fmt.Printf("It'll take %.0f blocks to get there", math.Abs(float64(x))+math.Abs(float64(y)))
			return
		}
		visits[key] = true
	}
	fmt.Printf("You are at %d, %d facing %s\n", x, y, facing)
	fmt.Printf("It'll take %.0f blocks to get there", math.Abs(float64(x))+math.Abs(float64(y)))
}

func newFacing(facing, turn string) string {
	switch facing {
	case "N":
		if turn == "R" {
			return "E"
		}
		return "W"
	case "E":
		if turn == "R" {
			return "S"
		}
		return "N"
	case "S":
		if turn == "R" {
			return "W"
		}
		return "E"
	case "W":
		if turn == "R" {
			return "N"
		}
		return "S"
	default:
		log.Fatalf("Tried to turn %s from %s", turn, facing)
		return ""
	}
}
