package main

import (
	"fmt"
	"math"
)

type player struct {
	health int
	damage int
	armor  int
}

type item struct {
	name   string
	cost   int
	damage int
	armor  int
}

// Weapons:    Cost  Damage  Armor
// Dagger        8     4       0
// Shortsword   10     5       0
// Warhammer    25     6       0
// Longsword    40     7       0
// Greataxe     74     8       0

// Armor:      Cost  Damage  Armor
// Leather      13     0       1
// Chainmail    31     0       2
// Splintmail   53     0       3
// Bandedmail   75     0       4
// Platemail   102     0       5

// Rings:      Cost  Damage  Armor
// Damage +1    25     1       0
// Damage +2    50     2       0
// Damage +3   100     3       0
// Defense +1   20     0       1
// Defense +2   40     0       2
// Defense +3   80     0       3

func main() {
	boss := player{health: 103, damage: 9, armor: 2}

	weapons := []item{
		{name: "Dagger", cost: 8, damage: 4},
		{name: "Shortsword", cost: 10, damage: 5},
		{name: "Warhammer", cost: 25, damage: 6},
		{name: "Longsword", cost: 40, damage: 7},
		{name: "Greataxe", cost: 74, damage: 8},
	}

	armors := []item{
		{name: "Leather", cost: 13, armor: 1},
		{name: "Chainmail", cost: 31, armor: 2},
		{name: "Splintmail", cost: 53, armor: 3},
		{name: "Bandedmail", cost: 75, armor: 4},
		{name: "Platemail", cost: 102, armor: 5},
		{name: "No armor"},
	}

	rings := []item{
		{name: "Damage +1", cost: 25, damage: 1},
		{name: "Damage +2", cost: 50, damage: 2},
		{name: "Damage +3", cost: 100, damage: 2},
		{name: "Defense +1", cost: 20, armor: 1},
		{name: "Defense +2", cost: 40, armor: 2},
		{name: "Defense +3", cost: 80, armor: 3},
		{name: "No ring right"},
		{name: "No ring left"},
	}

	leastGold := math.MaxInt32

	for _, weapon := range weapons {

		for _, armor := range armors {

			for _, ring1 := range rings {

				for _, ring2 := range rings {
					if ring1.name == ring2.name {
						continue
					}
					inventory := []item{weapon, armor, ring1, ring2}

					me := player{
						health: 100,
						damage: totalDamage(inventory),
						armor:  totalArmor(inventory),
					}

					cost := totalCost(inventory)

					if leastGold > cost && playerWins(me, boss) {
						fmt.Printf("inventory = %+v\n", inventory)
						fmt.Printf("me = %+v\n", me)
						fmt.Printf("cost = %+v\n", cost)
						fmt.Printf("playerWins(me, boss) = %+v\n", true)
						leastGold = cost
					}
				}
			}
		}
	}
}

func totalCost(inv []item) int {
	total := 0
	for _, i := range inv {
		total += i.cost
	}
	return total
}

func totalDamage(inv []item) int {
	total := 0
	for _, i := range inv {
		total += i.damage
	}
	return total
}

func totalArmor(inv []item) int {
	total := 0
	for _, i := range inv {
		total += i.armor
	}
	return total
}

func playerWins(me, boss player) bool {
	myDamage := me.damage - boss.armor
	if myDamage < 1 {
		myDamage = 1
	}
	bossDamage := boss.damage - me.armor
	if bossDamage < 1 {
		bossDamage = 1
	}

	bossTurnsToDie := math.Ceil(float64(boss.health) / float64(myDamage))
	meTurnsToDie := math.Ceil(float64(me.health) / float64(bossDamage))
	return bossTurnsToDie <= meTurnsToDie
}
