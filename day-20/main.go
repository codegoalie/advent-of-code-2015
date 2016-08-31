package main

import (
	"fmt"
	"math"
)

func main() {
	var house, target, presents, maxPresents int64
	house = 1
	target = 29000000
	for {
		presents = presentsAtHouse(house)
		if presents > maxPresents {
			maxPresents = presents
		}
		if presents >= target {
			fmt.Printf("House %d got %d presents.\n", house, presents)
			return
		}

		if math.Mod(float64(house), 10000) == 0 {
			fmt.Printf("house = %+v\n", house)
			fmt.Printf("maxPresents = %+v\n", maxPresents)
		}
		house++
	}
}

func presentsAtHouse(house int64) int64 {
	var i, presents int64
	for i = 1; i <= (house/2)+1; i++ {
		if i*50 >= house && math.Mod(float64(house), float64(i)) == 0 {
			presents += i * 11
		}
	}
	presents += house * 11

	return presents
}
