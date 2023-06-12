package main

import "fmt"

func main() {

	for i := 1; i <= 15; i++ {
		fmt.Println(fizzBuzz(i))
	}
}

func fizzBuzz(i int) string {
	if i%3 == 0 && i%5 == 0 {
		return "fizzbuzz"
	} else if i%3 == 0 {
		return "fizz"
	} else if i%5 == 0 {
		return "buzz"
	} else {
		return fmt.Sprintf("%d", i)
	}
}
