package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Message is the heding text and attachments for the Slack message
type Message struct {
	Text string `json:"text"`
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	transforms := map[string][]string{}
	mode := "transforms"
	molecule := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			mode = "molecule"
			continue
		}

		if mode == "transforms" {
			splits := strings.Split(line, " => ")
			transforms[splits[0]] = append(transforms[splits[0]], splits[1])
		} else {
			molecule = line
		}
	}

	// molecule = molecule[:40]

	fmt.Printf("len(transforms) = %+v\n", len(transforms))
	fmt.Printf("molecule = %+v\n", molecule)

	distincts := map[string]bool{"e": true}
	newDistincts := map[string]bool{}
	round := 1

	for {
		longest := 0
		for mol := range distincts {
			for match, replacements := range transforms {
				// fmt.Sprintf("match = %+v\n", match)
				start := 0
				for i := 0; i < strings.Count(mol, match); i++ {
					// fmt.Sprintf("  start = %+v\n", start)
					for _, replacement := range replacements {
						// fmt.Sprintf("  replacement = %+v\n", replacement)
						new := strings.Replace(mol[start:], match, replacement, 1)
						new = mol[:start] + new
						// fmt.Sprintf("  new = %+v\n", new)
						newDistincts[new] = true
						if len(new) > longest {
							longest = len(new)
						}
						if new == molecule {
							fmt.Printf("found it in round = %+v\n", round)
							return
						}
					}
					start = strings.Index(mol[start:], match) + start + len(match)
				}
			}
		}
		distincts = newDistincts
		newDistincts = map[string]bool{}
		fmt.Printf("  len(distincts) = %+v\n", len(distincts))
		fmt.Printf("  longest = %+v\n", longest)
		round++
		fmt.Printf("round = %+v\n", round)
	}
}
