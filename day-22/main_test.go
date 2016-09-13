package main

import (
	"strconv"
	"testing"
)

func TestMakeStrategy(t *testing.T) {
	cases := map[int][9]int{
		0:        [9]int{-1, -1, -1, -1, -1, -1, -1, -1, -1},
		1:        [9]int{0, -1, -1, -1, -1, -1, -1, -1, -1},
		2:        [9]int{1, -1, -1, -1, -1, -1, -1, -1, -1},
		3:        [9]int{2, -1, -1, -1, -1, -1, -1, -1, -1},
		4:        [9]int{3, -1, -1, -1, -1, -1, -1, -1, -1},
		5:        [9]int{4, -1, -1, -1, -1, -1, -1, -1, -1},
		6:        [9]int{-1, 0, -1, -1, -1, -1, -1, -1, -1},
		7:        [9]int{0, 0, -1, -1, -1, -1, -1, -1, -1},
		12:       [9]int{-1, 1, -1, -1, -1, -1, -1, -1, -1},
		10077695: [9]int{4, 4, 4, 4, 4, 4, 4, 4, 4},
	}

	for mask, expected := range cases {
		t.Run(strconv.Itoa(mask), func(t *testing.T) {
			actual := makeStrategy(mask)
			if actual != expected {
				t.Errorf("makeStrategy(%d) = %+v; want %+v", mask, actual, expected)
			}
		})
	}
}
