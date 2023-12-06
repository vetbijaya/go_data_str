package main

import (
	"fmt"
	"strings"
)

// Variadic function to join strings and separate them with hyphen
func joinString(element ...string) string {
	return strings.Join(element, "-")
}

func main() {
	// to show zero argument
	fmt.Println(joinString())
	// to show multiple arguments
	fmt.Println(joinString("Golang", "Programmer"))
	fmt.Println(joinString("Go", "Coding", "is", "fun"))
}
