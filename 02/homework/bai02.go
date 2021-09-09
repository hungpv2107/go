package main

import (
	"fmt"
)

/*
Bài tập 2
*/
func main() {
	longestElement := findMaxLengthElement([]string{"aba", "aa", "ad", "c", "vcd"})
	fmt.Println("Slice with logest elements: ", longestElement)
}

/*
Tìm các phần tử có độ dài lớn nhất trong mảng string truyền vào
*/
func findMaxLengthElement(inputSlice []string) (result []string) {
	max := -1 // -1 để đảm bảo bé hơn length của string
	for _, element := range inputSlice {
		if len(element) < max {
			// Skip element có độ dài ngắn hơn
			continue
		} else if len(element) > max {
			// Found longer string, reset result
			result = result[:0]
		}
		// Update max và add to result
		max = len(element)
		result = append(result, element)
	}
	return
}
