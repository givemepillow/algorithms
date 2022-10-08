package main

import "fmt"

func FindMax(array []int, left int, right int) int {
	if left == right {
		return array[left]
	}
	mid := (left + right) / 2
	if array[left] > array[mid+1] {
		return FindMax(array, left, mid)
	} else {
		return FindMax(array, mid+1, right)
	}
}

func main() {
	var array = [10]int{55, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(FindMax(array[:], 0, 9))
}
