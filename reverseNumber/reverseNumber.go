package main

import "fmt"

func reverseNumber(num []int) {
	for i, j := 0, len(num)-1; i < j; i, j = i+1, j-1 {
		num[i], num[j] = num[j], num[i]
	}
}

func main() {
	num := []int{1, 2, 3}
	reverseNumber(num)
	fmt.Println(num)
}
