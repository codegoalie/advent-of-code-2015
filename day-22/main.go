package main

import "fmt"

type person struct {
	health int
	damage int
}

type wizard struct {
	health int
	armor  int
	mana   int
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
	boss := person{health: 55, damage: 8}
	me := wizard{health: 50, mana: 500}
	fmt.Printf("boss = %+v\n", boss)
	fmt.Printf("me = %+v\n", me)

	fmt.Printf("spells = %+v\n", spells)
}
