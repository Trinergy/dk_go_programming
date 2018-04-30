package echo

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

// 1.1 Modify echo to also print os.Args[0]
func argZeroEcho() {
	s, sep := "", ""
	for _, arg := range os.Args[:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

// 1.2 Modify echo to print index and value of each arguments, one per line
func echoWithIndex() {
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
	}
}

// 1.3 Compare strings.Join with simple Echo in benchmark tests
func echoWithJoin() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	simpleEcho()
	argZeroEcho()
	echoWithIndex()
	echoWithJoin()
}
