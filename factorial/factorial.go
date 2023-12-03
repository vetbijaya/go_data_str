package main

import "fmt"

// Factorial of a number is the product of multiplication of a number n with every preceding number till it reaches 1. Factorial of 0 is 1.
// factorial functions get the factorial of given number
func factorial(n int) int {
	if n == 0 { // another way: if n <=1
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	fmt.Println("The factorial of 5 is:", factorial(5)) // Output: The factorial of 5 is: 120
}
