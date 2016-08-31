package main

import "testing"

func TestPresentsAthouse(t *testing.T) {
	presents := presentsAtHouse(4)
	if presents != 77 {
		t.Errorf("wanted 77 presents but got %d", presents)
	}
}
