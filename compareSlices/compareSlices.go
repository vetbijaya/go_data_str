package main

import "fmt"

func compareSlices(slice1, slice2 []int) []int {
	//Initialize a map to store unique values from slice1
	uniqueValues := make(map[int]bool)

	// Create a slice to store the common elements
	commonElements := []int{}

	// Populate the map with values from slice1
	for _, val := range slice1 {
		uniqueValues[val] = true
	}

	// Iterate through slice2 and check for common elements
	for _, val := range slice2 {
		if uniqueValues[val] {
			commonElements = append(commonElements, val)
		}
	}
	return commonElements
}

func main() {
	dataset1 := []int{1, 2, 3, 4, 5}
	dataset2 := []int{3, 4, 5, 6, 7}

	common := compareSlices(dataset1, dataset2)

	fmt.Println("Common elements:", common) // Output: [3 4 5]
}
