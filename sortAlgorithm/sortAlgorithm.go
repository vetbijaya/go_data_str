package main

import (
	"fmt"
	"sort"
)

func main() {
	// sort a slice of integers in ascending order
	numbers := []int{5, 9, 2, 4, 1, 8, 3, 6}
	sort.Ints(numbers)
	fmt.Println("sorted slice of integers", numbers)

	// sort a slice of integers in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
	fmt.Println("sorted slice of integers in descending order", numbers)

	// search for an index of an element in sorted slice of integers
	x := []int{1, 2, 3, 4}
	index := sort.SearchInts(x, 3)
	fmt.Println("index of given number is", index)

	// sort a slice of strings
	colors := []string{"Red", "Blue", "Pink", "Yellow"}
	sort.Strings(colors)
	fmt.Println("sorted slice of strings", colors)

	// sort a slice of stings in descending order
	sort.Sort(sort.Reverse(sort.StringSlice(colors)))
	fmt.Println("sorted slice of strings", colors)
}
