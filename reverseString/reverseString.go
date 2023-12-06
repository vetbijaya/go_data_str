package main

import "fmt"

// The reverseString function takes a string as input, converts it to a rune slice, and
// then reverses the order of the elements in the slice by swapping elements from the beginning to the end.
// Finally, it converts the rune slice back to a string and returns the reversed string.

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	original := "Hello, Golang!"
	reversed := reverseString(original)
	fmt.Printf("Original: %s\nReversed: %s\n", original, reversed)
}
