package main

import (
	"fmt"
)

/*
Bài tập 3
*/
func main() {
	uniqueSlice := removeDuplicateElements([]int{3, 7, 1, 2, 5, 8, 6, 9, 5, 2, 8, 9, 7, 6})
	fmt.Println("Slice without duplicates is: ", uniqueSlice)
}

/*
Remove những phần tử bị trùng nhau và trả về mảng gồm các phần tử duy nhất
*/
func removeDuplicateElements(inputSlice []int) (result []int) {
	keys := map[int]bool{}
	for _, element := range inputSlice {
		// Nếu element không tồn tại thì add vào keys và append vào result
		if _, value := keys[element]; !value {
			keys[element] = true
			result = append(result, element)
		}
	}
	return
}
