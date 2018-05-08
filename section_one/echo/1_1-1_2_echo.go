package main

import (
	"fmt"
	"os"
	"strings"
)

// SimpleEcho is the default echo program
func SimpleEcho() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// ArgZeroEcho echoes the arguments with the invoke path
// 1.1 Modify echo to also print os.Args[0]
func ArgZeroEcho() {
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// WithIndex echoes the index and value of each argumet in individual lines
// 1.2 Modify echo to print index and value of each arguments, one per line
func WithIndex() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

// WithJoin is a simple echo function that uses the strings.Join lib
// 1.3 Compare strings.Join with simple Echo in benchmark tests
func WithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	SimpleEcho()
}
