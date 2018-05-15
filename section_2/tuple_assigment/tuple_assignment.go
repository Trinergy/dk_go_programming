package main

import "fmt"

func main() {
	gcd := gcd(28, 7)
	fmt.Printf("RESULT: %v\n", gcd)

	fib := fib(6)
	fmt.Printf("FIB RESULT:%d", fib)

}

func gcd(x int, y int) int {
	for y != 0 {
		x, y = y, x%y
		fmt.Printf("x: %d, y: %d\n", x, y)
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}
