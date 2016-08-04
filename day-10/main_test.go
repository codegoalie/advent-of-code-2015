package main

import "testing"

func TestLookAndSee(t *testing.T) {
	cases := map[string]string{
		"1":      "11",
		"11":     "21",
		"21":     "1211",
		"1211":   "111221",
		"111221": "312211",
	}
	for input, expected := range cases {
		actual := lookAndSee(input)
		if actual != expected {
			t.Errorf("lookAndSee(%v) = %v want %v", input, actual, expected)
		}
	}

}
