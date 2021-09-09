package main

import (
	"fmt"
	"sort"
)

/*
Bài tập 1
*/
func main() {
	inputSlice := []int{2, 1, 3, 4, 9, 8}
	fmt.Println("Origin list: ", inputSlice)

	// Cach 1
	fmt.Println("\n--- Cach 1: Keep Order ---")
	secondLargestElement01 := findSecondLargestElementKeepOrder(inputSlice)
	fmt.Println("Second largest element is: ", secondLargestElement01)
	fmt.Println("Origin list: ", inputSlice)

	// Cach 2
	fmt.Println("\n--- Cach 2: Not Keep Order ---")
	secondLargestElement02 := findSecondLargestElementNotKeepOrder(inputSlice)
	fmt.Println("Second largest element is: ", secondLargestElement02)
	fmt.Println("Origin list: ", inputSlice)
}

/*
Tìm số lớn thứ 2 trong slice truyền vào, bảo toàn thứ tự phần tử trong slice
*/
func findSecondLargestElementKeepOrder(inputSlice []int) int {
	largestElement := 0
	secondLargestElement := 0
	for _, element := range inputSlice {
		if largestElement < element {
			secondLargestElement = largestElement
			largestElement = element
		} else if secondLargestElement < element {
			secondLargestElement = element
		}
	}
	return secondLargestElement
}

/*
Tìm số lớn thứ 2 trong slice truyền vào, không bảo toàn thứ tự phần tử trong slice
*/
func findSecondLargestElementNotKeepOrder(inputSlice []int) int {
	sort.Ints(inputSlice)
	return inputSlice[len(inputSlice)-2]
}
