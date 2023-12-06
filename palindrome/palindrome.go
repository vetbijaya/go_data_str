package main

import "fmt"

//The isPalindrome function checks if a given string is a palindrome. It uses two pointers (i and j) starting from the beginning and end of the string,
//comparing characters until they meet in the middle. If at any point the characters are not equal, the function returns false; otherwise, it returns true.

func isPalindrome(s string) bool {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		if runes[i] != runes[j] {
			return false
		}
	}
	return true
}

func isPalindromeNumber(num int) bool {
	if num < 0 {
		return false // Negative numbers are not palindromes
	}
	originalNum := num
	reversedNum := 0

	for num > 0 {
		digit := num % 10
		reversedNum = reversedNum*10 + digit
		num /= 10
	}
	return originalNum == reversedNum
}

func main() {
	result := isPalindrome("level")
	fmt.Println(result)
	number := 121
	if isPalindromeNumber(number) {
		fmt.Println(number, "is a palindrome number")
	} else {
		fmt.Println(number, "is not a palindrome number")
	}
}
