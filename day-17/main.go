package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	containers := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		size, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}

		containers = append(containers, size)
	}

	count := 0
	for i := 1; i <= len(containers); i++ {
		for _, combo := range perms(containers, i) {
			sum := 0
			for _, size := range combo {
				sum += size
			}

			if sum == 150 {
				count++
				fmt.Printf("combo = %+v\n", combo)
			}
		}
	}
	fmt.Printf("count = %+v\n", count)
}

func perms(a []int, choose int) [][]int {
	result := [][]int{}

	if choose == 1 {
		for _, e := range a {
			result = append(result, []int{e})
		}
		return result
	}

	for i := 0; i < len(a); i++ {
		subSlice := []int{}
		for ii := i + 1; ii < len(a); ii++ {
			subSlice = append(subSlice, a[ii])
		}

		subs := perms(subSlice, choose-1)
		for _, sub := range subs {
			candidate := []int{a[i]}
			for _, e := range sub {
				candidate = append(candidate, e)
			}

			result = append(result, candidate)
		}
	}

	return result
}
