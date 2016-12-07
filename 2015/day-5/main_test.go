package main

import "testing"

func TestHasTwoDupes(t *testing.T) {
	cases := map[string]bool{
		"xyxy":       true,
		"aabcdefgaa": true,
		"abc":        false,
		"aaa":        false,
		"aaaa":       true,
		"aaaaaa":     true,
	}

	for s, expected := range cases {
		actual := hasTwoDupes(s)
		if actual != expected {
			t.Errorf("hasTwoDupes(%s) == %v, want %v", s, actual, expected)
		}
	}
}
