package main

import "fmt"

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

func sequenceNumber(row, col int) int {
	top := row + col - 1
	offset := top - col
	return ((top * (top + 1)) / 2) - offset
}

func main() {
	row := 2981
	col := 3075
	seqNum := sequenceNumber(row, col)
	fmt.Printf("sequenceNumber(%d, %d) = %+v\n", row, col, seqNum)

	code := 20151125
	for i := 1; i < seqNum; i++ {
		code = (code * 252533) % 33554393
	}

	fmt.Printf("code = %+v\n", code)
}
