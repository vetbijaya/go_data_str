/*
Write a program that prints the numbers from 1 to 100.
But for multiples of 3, print "Fizz" instead of the number, and for the multiples of 5, print "Buzz."
For numbers that are multiples of both three and five, print "FizzBuzz."
*/

package main

import "fmt"

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
		} else if i%3 == 0 {
			fmt.Println("Fizz")
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func main() {
	fizzBuzz()
}
