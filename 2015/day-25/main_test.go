package main

import (
	"fmt"
	"testing"
)

//    | 1   2   3   4   5   6   7   8   9   10
// ---+---+---+---+---+---+---+---+---+---+---+
//  1 |  1   3   6  10  15  21  28  36  45  55
//  2 |  2   5   9  14  20  27  35  44  54
//  3 |  4   8  13  19  26  34  42  53
//  4 |  7  12  18  25  33  42  52
//  5 | 11  17  24  32  41  51
//  6 | 16  23  31  40  50
//  7 | 22  30  39  49  60
//  8 | 29  38  48  59
//  9 | 37  47  58
// 10 | 46  57
// 11 | 56

func TestSequenceNumber(t *testing.T) {
	testCases := []struct {
		row, col, seqNum int
	}{
		{row: 1, col: 1, seqNum: 1},
		{row: 3, col: 1, seqNum: 4},
		{row: 1, col: 6, seqNum: 21},
		{row: 7, col: 5, seqNum: 60},
		{row: 10, col: 2, seqNum: 57},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d:%d", tc.row, tc.col), func(t *testing.T) {
			if got := sequenceNumber(tc.row, tc.col); got != tc.seqNum {
				t.Errorf("sequenceNumber(%d, %d) = %d, want %d", tc.row, tc.col, got, tc.seqNum)
			}
		})
	}
}
