package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Trinergy/dk_go_programming/section_2/length_conv"
)

// 2-2 is an example of package imports
func main() {
	for _, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "length conv: %v\n", err)
			os.Exit(1)
		}
		f := lengthconv.Feet(t)
		m := lengthconv.Meters(t)
		fmt.Printf("%s = %s, %s = %s\n",
			f, lengthconv.FToM(f), m, lengthconv.MToF(m))
	}
}
