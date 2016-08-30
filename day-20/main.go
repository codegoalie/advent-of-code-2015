package main

import (
	"fmt"
	"math"
)

func main() {
	var i, house, last, target, presents, maxPresents int64
	last = 1
	house = 1
	target = 29000000
	for {
		presents = presentsAtHouse(house)
		if presents > maxPresents {
			maxPresents = presents
		}
		if presents >= target {
			fmt.Printf("ALMOST THERE House %d got %d presents.\n", house, presents)
			for i = last; i <= house; i++ {
				presents = presentsAtHouse(i)
				fmt.Printf("House %d got %d presents.\n", i, presents)
				if presents >= target {
					return
				}
			}
		}
		last = house

		if math.Mod(float64(house), 10000) == 0 {
			fmt.Printf("house = %+v\n", house)
			fmt.Printf("maxPresents = %+v\n", maxPresents)
		}
		house++
	}
}

func presentsAtHouse(house int64) int64 {
	var i, presents int64
	for i = 1; i <= house; i++ {
		if math.Mod(float64(house), float64(i)) == 0 {
			presents += i * 10
		}
	}

	return presents
}
