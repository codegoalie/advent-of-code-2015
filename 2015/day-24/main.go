package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type presents []int

func (p presents) Len() int {
	return len(p)
}

func (p presents) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p presents) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	weights := presents{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		weight, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		weights = append(weights, weight)

	}
	fmt.Printf("weights = %+v\n", weights)
	fmt.Printf("len(weights) = %+v\n", len(weights))

	fmt.Printf("sumWeight(weights) = %+v\n", sumWeight(weights))

	candidates := [][]int{}

	start := 3
	end := 14
	thirdSum := 384

	//test
	// start := 2
	// end := 4
	// thirdSum := 20

	for i := start; i < end; i++ {
		fmt.Printf("i = %+v\n", i)
		for _, combo := range perms(weights, i) {
			if sumWeight(combo) == thirdSum {
				candidates = append(candidates, combo)
			}
		}
		if len(candidates) > 0 {
			break
		}
	}

	// mathutil.PermutationFirst(weights)
	// for {
	// 	for i := 0; i < len(weights)-1; i++ {
	// 		for ii := i + 1; ii < len(weights)-1; ii++ {

	// 			one := weights[0:i]
	// 			if len(one) <= len(candidates[0]) {
	// 				two := weights[i:ii]
	// 				three := weights[ii:len(weights)]
	// 				oneSum := sumWeight(one)
	// 				twoSum := sumWeight(two)
	// 				threeSum := sumWeight(three)

	// 				if oneSum == twoSum && twoSum == threeSum {
	// 					if len(candidates[0]) == len(one) {
	// 						candidates = append(candidates, one)
	// 					} else {
	// 						candidates = [][]int{one}
	// 					}
	// 				}
	// 			}
	// 		}
	// 	}
	// 	if !mathutil.PermutationNext(weights) {
	// 		break
	// 	}
	// }

	fmt.Println("Done with perms, finding lowest entanglement")
	fmt.Printf("candidates = %+v\n", candidates)

	lowEnt := math.MaxInt64
	for _, cand := range candidates {
		quan := quantumEntangelment(cand)
		if quan < lowEnt {
			lowEnt = quan
		}
	}

	fmt.Printf("lowEnt = %+v\n", lowEnt)

}

func sumWeight(items []int) int {
	var sum int
	for _, w := range items {
		sum += w
	}

	return sum
}

func quantumEntangelment(items []int) int {
	entanglement := 1
	for _, w := range items {
		entanglement *= w
	}

	return entanglement
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
