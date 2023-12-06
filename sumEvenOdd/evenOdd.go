// Problem: Sum of even and odd numbers

package main

import "fmt"

// sumEven functions add the even numbers
func sumEven(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num%2 == 0 {
			sum += num
		}
	}
	return sum
}

// sumOdd functions add the odd numbers
func sumOdd(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		if num%2 != 0 {
			sum += num
		}
	}
	return sum
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	result := sumEven(numbers)
	res := sumOdd(numbers)
	fmt.Printf("The sum of even numbers: %d\n", result)
	fmt.Printf("The sum of odd numbers: %d\n", res)
}
