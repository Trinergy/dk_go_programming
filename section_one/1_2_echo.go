package main

import (
	"fmt"
	"os"
	"strings"
)

// default
func simpleEcho() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// 1.1
func argZeroEcho() {
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// 1.2
func echoWithIndex() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

// 1.3
func echoWithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	simpleEcho()
	argZeroEcho()
	echoWithIndex()
	echoWithJoin()
}
