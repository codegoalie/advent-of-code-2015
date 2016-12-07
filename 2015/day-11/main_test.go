package main

import "testing"

func TestHasStraight(t *testing.T) {
	cases := map[string]bool{
		"a":       false,
		"abd":     false,
		"abc":     true,
		"dgxyz":   true,
		"ihrstwa": true,
		"werfklj": false,
	}
	for input, expected := range cases {
		actual := hasStraight(input)
		if actual != expected {
			t.Errorf("hasStraight(%v) = %v want %v", input, actual, expected)
		}
	}
}

func TestHasNoInvalids(t *testing.T) {
	cases := map[string]bool{
		"i":     false,
		"o":     false,
		"l":     false,
		"abb":   true,
		"dfgds": true,
		"ofdd":  false,
		"ffise": false,
		"rfdsl": false,
	}
	for input, expected := range cases {
		actual := hasNoInvalids(input)
		if actual != expected {
			t.Errorf("hasNoInvalids(%v) = %v want %v", input, actual, expected)
		}
	}
}

func TestHasTwoPairs(t *testing.T) {
	cases := map[string]bool{
		"i":       false,
		"abb":     false,
		"efffe":   false,
		"oofdd":   true,
		"ffiseff": true,
		"rffddsl": true,
	}
	for input, expected := range cases {
		actual := hasTwoPairs(input)
		if actual != expected {
			t.Errorf("hasTwoPairs(%v) = %v want %v", input, actual, expected)
		}
	}
}

func TestIncrementPassword(t *testing.T) {
	cases := map[string]string{
		"a":    "b",
		"z":    "aa",
		"efg":  "efh",
		"xz":   "ya",
		"azzz": "baaa",
	}
	for input, expected := range cases {
		actual := incrementPassword(input)
		if actual != expected {
			t.Errorf("incrementPassword(%v) = %v want %v", input, actual, expected)
		}
	}
}
