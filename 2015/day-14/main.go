package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type reindeer struct {
	name                    string
	speed, goTime, restTime int
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	deer := []reindeer{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		splits := strings.Split(line, " ")

		name := splits[0]
		speed, err := strconv.Atoi(splits[3])
		if err != nil {
			log.Fatal(err)
		}
		goTime, err := strconv.Atoi(splits[6])
		if err != nil {
			log.Fatal(err)
		}
		restTime, err := strconv.Atoi(splits[13])
		if err != nil {
			log.Fatal(err)
		}

		deer = append(deer, reindeer{
			name:     name,
			speed:    speed,
			goTime:   goTime,
			restTime: restTime,
		})
	}

	totalTime := 2503

	points := map[string]int{}

	for i := 1; i <= totalTime; i++ {
		best := 0
		round := map[string]int{}
		for _, d := range deer {
			fullCircuits := i / (d.goTime + d.restTime)
			distance := fullCircuits * d.speed * d.goTime
			partialTime := i % (d.goTime + d.restTime)
			if partialTime > d.goTime {
				distance += (d.speed * d.goTime)
			} else {
				distance += (d.speed * partialTime)
			}

			if distance > best {
				best = distance
			}
			round[d.name] = distance
			fmt.Printf("%+v = %+v\n", d.name, distance)
		}
		// Get winner and add point
		for name, distance := range round {
			if distance == best {
				points[name]++
			}
		}
		fmt.Printf("points = %+v\n", points)
	}

	best := 0
	for _, points := range points {
		if points > best {
			best = points
		}
	}
	fmt.Printf("best = %+v\n", best)
}
