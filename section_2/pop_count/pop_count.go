package main

import (
	"fmt"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// PopCountWithLoop returns the population count (number of set bits) of x.
func PopCountWithLoop(x uint64) int {
	var popcount int
	for i := 0; i <= 8; i++ {
		popcount += int(pc[byte(x>>(uint64(i)*8))])
	}
	return popcount
}

func main() {
	pop := PopCount(2038998)
	fmt.Printf("POP COUNT: %v", pop)
}
