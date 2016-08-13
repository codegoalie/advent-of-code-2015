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

	replacements := map[string][]string{}
	mode := "replacements"
	molecule := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			mode = "molecule"
			continue
		}

		if mode == "replacements" {
			splits := strings.Split(line, " => ")
			replacements[splits[0]] = append(replacements[splits[0]], splits[1])
		} else {
			molecule = line
		}
	}

	// molecule = molecule[:40]

	fmt.Printf("len(replacements) = %+v\n", len(replacements))
	fmt.Printf("molecule = %+v\n", molecule)

	distincts := map[string]bool{}

	for match, replacements := range replacements {
		// fmt.Printf("match = %+v\n", match)
		start := 0
		for i := 0; i < strings.Count(molecule, match); i++ {
			// fmt.Printf("  start = %+v\n", start)
			for _, replacement := range replacements {
				// fmt.Printf("  replacement = %+v\n", replacement)
				new := strings.Replace(molecule[start:], match, replacement, 1)
				new = molecule[:start] + new
				// fmt.Printf("  new = %+v\n", new)
				distincts[new] = true
			}
			start = strings.Index(molecule[start:], match) + start + len(match)
		}
	}

	// fmt.Printf("distincts = %+v\n", distincts)
	fmt.Printf("len(distincts) = %+v\n", len(distincts))
}
