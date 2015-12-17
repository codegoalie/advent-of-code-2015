package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
)

var secret = "yzbqklnj"

func main() {
	candidate := 1

	for {
		hash := fmt.Sprintf("%x", md5.Sum([]byte(secret+strconv.Itoa(candidate))))

		if stringStartsWithZeros(6, hash) {
			fmt.Println("We did it!", candidate)
			return
		}

		candidate++
	}
}

func stringStartsWithZeros(leadingZeros int, input string) bool {
	var zero = []byte("0")[0]
	for _, b := range []byte(input)[:leadingZeros] {
		if b != zero {
			return false
		}
	}
	return true
}
