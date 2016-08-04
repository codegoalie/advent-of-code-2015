package main

import "testing"

func TestPerm(t *testing.T) {
	cases := map[int][][]int{
		1: {{0}},
		2: {{0, 1}, {1, 0}},
		3: {
			{0, 1, 2}, {1, 0, 2},
			{2, 0, 1}, {0, 2, 1},
			{1, 2, 0}, {2, 1, 0},
		},
		4: {
			{0, 1, 2, 3}, {1, 0, 2, 3}, {2, 0, 1, 3}, {0, 2, 1, 3}, {1, 2, 0, 3}, {2, 1, 0, 3},
			{3, 1, 2, 0}, {1, 3, 2, 0}, {2, 3, 1, 0}, {3, 2, 1, 0}, {1, 2, 3, 0}, {2, 1, 3, 0},
			{3, 0, 2, 1}, {0, 3, 2, 1}, {2, 3, 0, 1}, {3, 2, 0, 1}, {0, 2, 3, 1}, {2, 0, 3, 1},
			{3, 0, 1, 2}, {0, 3, 1, 2}, {1, 3, 0, 2}, {3, 1, 0, 2}, {0, 1, 3, 2}, {1, 0, 3, 2},
		},
	}

	for n, expected := range cases {
		actual := perm(n)
		if !deepCompareSlice(actual, expected) {
			t.Errorf("perm(%d) = %+v, want %+v", n, actual, expected)
		}
	}
}

func deepCompareSlice(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		for ii := range a[i] {
			if a[i][ii] != b[i][ii] {
				return false
			}
		}
	}

	return true
}
