package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type ingredient struct {
	capacity, durability, flavor, texture, calories int
}

func main() {
	filename := "input.txt"
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error opening file:", filename, err)
	}

	ingredients := map[string]ingredient{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		splits := strings.Split(line, " ")

		name := splits[0]
		name = name[:len(name)-1]
		capacity, err := strconv.Atoi(splits[2][:len(splits[2])-1])
		if err != nil {
			log.Fatal(err)
		}
		durability, err := strconv.Atoi(splits[4][:len(splits[4])-1])
		if err != nil {
			log.Fatal(err)
		}
		flavor, err := strconv.Atoi(splits[6][:len(splits[6])-1])
		if err != nil {
			log.Fatal(err)
		}
		texture, err := strconv.Atoi(splits[8][:len(splits[8])-1])
		if err != nil {
			log.Fatal(err)
		}
		calories, err := strconv.Atoi(splits[10])
		if err != nil {
			log.Fatal(err)
		}

		ingredients[name] = ingredient{
			capacity:   capacity,
			durability: durability,
			flavor:     flavor,
			texture:    texture,
			calories:   calories,
		}
	}

	fmt.Printf("ingredients = %+v\n", ingredients)

	highScore := 0

	maxIngredients := 100
	for i := 0; i <= maxIngredients; i++ {
		for j := 0; j <= maxIngredients; j++ {
			for k := 0; k <= maxIngredients; k++ {
				for l := 0; l <= maxIngredients; l++ {
					if i+j+k+l != maxIngredients {
						continue
					}

					capacity := ingredients["Sprinkles"].capacity*i +
						ingredients["Butterscotch"].capacity*j +
						ingredients["Chocolate"].capacity*k +
						ingredients["Candy"].capacity*l

					if capacity < 0 {
						capacity = 0
					}

					durability := ingredients["Sprinkles"].durability*i +
						ingredients["Butterscotch"].durability*j +
						ingredients["Chocolate"].durability*k +
						ingredients["Candy"].durability*l

					if durability < 0 {
						durability = 0
					}

					flavor := ingredients["Sprinkles"].flavor*i +
						ingredients["Butterscotch"].flavor*j +
						ingredients["Chocolate"].flavor*k +
						ingredients["Candy"].flavor*l

					if flavor < 0 {
						flavor = 0
					}

					texture := ingredients["Sprinkles"].texture*i +
						ingredients["Butterscotch"].texture*j +
						ingredients["Chocolate"].texture*k +
						ingredients["Candy"].texture*l

					if texture < 0 {
						texture = 0
					}

					calories := ingredients["Sprinkles"].calories*i +
						ingredients["Butterscotch"].calories*j +
						ingredients["Chocolate"].calories*k +
						ingredients["Candy"].calories*l

					if calories < 0 {
						calories = 0
					}

					if calories != 500 {
						continue
					}

					score := texture * flavor * durability * capacity

					if score > highScore {
						highScore = score
					}
				}
			}
		}
	}

	fmt.Printf("highScore = %+v\n", highScore)
}
