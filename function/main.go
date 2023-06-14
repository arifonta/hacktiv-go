package main

import (
	"fmt"
)

func main() {
	fib(5)
	fib(-1)
}

func fib(num int) {
	a := 0
	b := 1
	c := b
	fmt.Printf("Print: %d %d", a, b)
	for true {
		c = b
		b = a + b
		if b >= num {
			fmt.Println()
			break
		}
		a = c
		fmt.Printf(" %d", b)
	}
}
