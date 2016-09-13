package main

import (
	"fmt"
	"math"
)

type person struct {
	health int
	damage int
}

type wizard struct {
	health int
	armor  int
	mana   int
}

type effect interface {
	String() string
}

type spell struct {
	name     string
	cost     int
	modifier string
	effect   int
	length   int
}

var spells = [5]spell{
	{
		name:     "Magic Missle",
		cost:     53,
		modifier: "damage",
		effect:   4,
		length:   0,
	},
	{
		name:     "Drain",
		cost:     73,
		modifier: "drain",
		effect:   2,
		length:   0,
	},
	{
		name:     "Shield",
		cost:     113,
		modifier: "armor",
		effect:   7,
		length:   6,
	},
	{
		name:     "Poison",
		cost:     173,
		modifier: "damage",
		effect:   3,
		length:   6,
	},
	{
		name:     "Recharge",
		cost:     229,
		modifier: "mana",
		effect:   101,
		length:   5,
	},
}

func main() {
	lowestCost := math.MaxInt32

	for strategyMask := 0; strategyMask < 10077696; strategyMask++ {
		boss := person{health: 55, damage: 8}
		me := wizard{health: 50, mana: 500}
		strategy := makeStrategy(strategyMask)
		currentEffects := []spell{}
		var round, totalCost int

		for boss.health > 0 && me.health > 0 {
			// fmt.Printf("\n-- Player turn (round %d)--\n", round)
			// fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", me.health, me.armor, me.mana)
			// fmt.Printf("- Boss has %d hit points\n", boss.health)
			woreOff := []int{}
			for i, effect := range currentEffects {
				// fmt.Printf("%s's timer is %d\n", effect.name, effect.length)
				applyEffect(effect, &boss, &me)
				effect.length--
				if effect.length == 0 {
					woreOff = append(woreOff, i)
				}
			}
			for offset, i := range woreOff {
				index := i - offset
				wearEffectOff(currentEffects[index], &boss, &me)
				currentEffects[index] = currentEffects[len(currentEffects)-1]
				currentEffects = currentEffects[:len(currentEffects)-1]
			}

			indexToCast := strategy[round%len(strategy)]
			if indexToCast != -1 {
				spellToCast := spells[indexToCast]
				// fmt.Printf("Player casts %s\n", spellToCast.name)

				for _, effect := range currentEffects {
					if effect.name == spellToCast.name {
						break
					}
				}

				if spellToCast.length == 0 {
					applyEffect(spellToCast, &boss, &me)
				} else {
					currentEffects = append(currentEffects, spellToCast)
				}
				me.mana -= spellToCast.cost
				if me.mana < 0 {
					break
				}
				totalCost += spellToCast.cost
			} else {
				// fmt.Println("Player doesn't cast anything")
			}

			if boss.health <= 0 {
				break
			}

			// fmt.Printf("-- Boss turn (round %d)--\n", round)
			// fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", me.health, me.armor, me.mana)
			// fmt.Printf("- Boss has %d hit points\n", boss.health)

			woreOff = []int{}
			for i, effect := range currentEffects {
				// fmt.Printf("%s's timer is %d\n", effect.name, effect.length)
				applyEffect(effect, &boss, &me)
				currentEffects[i].length--
				if currentEffects[i].length == 0 {
					woreOff = append(woreOff, i)
				}
				if boss.health <= 0 {
					break
				}
			}
			for _, i := range woreOff {
				wearEffectOff(currentEffects[i], &boss, &me)
				currentEffects[i] = currentEffects[len(currentEffects)-1]
				currentEffects = currentEffects[:len(currentEffects)-1]
			}

			// fmt.Printf("Boss attacks for %d - %d = %d damage!\n", boss.damage, me.armor, boss.damage-me.armor)
			me.health -= boss.damage - me.armor

			round++
		}

		if boss.health < 0 && me.health > 0 && totalCost < lowestCost {
			lowestCost = totalCost
			fmt.Printf("new lowestCost = %+v\n", lowestCost)
		}
		// fmt.Println("\n\n GAME OVER")
		// fmt.Printf("- Player has %d hit points, %d armor, %d mana\n", me.health, me.armor, me.mana)
		// fmt.Printf("- Boss has %d hit points\n", boss.health)
		// fmt.Printf("totalCost = %+v\n", totalCost)
	}

	fmt.Printf("lowestCost = %+v\n", lowestCost)

}

func applyEffect(effect spell, boss *person, me *wizard) {
	// fmt.Printf("%s applies ", effect.name)
	switch effect.modifier {
	case "damage":
		boss.health -= effect.effect
		// fmt.Printf("and does %d damage!\n", effect.effect)
	case "drain":
		me.health += effect.effect
		boss.health -= effect.effect
		// fmt.Printf("and takes %d from Boss to Player\n", effect.effect)
	case "mana":
		me.mana += effect.effect
		// fmt.Printf("and gives %d mana to Player\n", effect.effect)
	case "armor":
		me.armor = effect.effect
		// fmt.Printf("and gives Player %d armor\n", effect.effect)
	}
}

func wearEffectOff(effect spell, boss *person, me *wizard) {
	// fmt.Printf("%s wore off.\n", effect.name)
	switch effect.modifier {
	case "armor":
		me.armor -= effect.effect
		// fmt.Printf("Player armor to %d\n", me.armor)
	}

}

func makeStrategy(mask int) [9]int {
	strategy := [9]int{}
	for i := 0; i < 9; i++ {
		strategyforTurn := (mask % 6) - 1
		mask /= 6
		strategy[i] = strategyforTurn
	}

	return strategy
}
