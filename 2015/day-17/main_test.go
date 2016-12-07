package main

import "testing"

func TestPerm(t *testing.T) {
	set := []int{1, 2, 3, 4}
	cases := map[int][][]int{
		0: {},
		1: {{1}, {2}, {3}, {4}},
		2: {
			{1, 2}, {1, 3}, {1, 4},
			{2, 3}, {2, 4},
			{3, 4},
		},
		3: {
			{1, 2, 3}, {1, 2, 4}, {1, 3, 4},
			{2, 3, 4},
		},
		4: {{1, 2, 3, 4}},
	}

	for n, expected := range cases {
		actual := perms(set, n)
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
