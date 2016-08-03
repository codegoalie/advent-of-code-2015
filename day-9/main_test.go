package main

import "testing"

func TestPerm(t *testing.T) {
	cases := map[int][][]int{
		1: [][]int{[]int{0}},
		2: [][]int{[]int{0, 1}, []int{1, 0}},
	}

	for n, expected := range cases {
		actual := [][]int{}
		perm(n, func(s []int) {
			actual = append(actual, s)
		})
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
