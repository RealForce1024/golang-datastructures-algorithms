package main

import "fmt"

func main() {
	fmt.Println(find_max([]int{-1, 1, 2, 3, 4, 5}))
	fmt.Println(find_max([]int{-1, -2, -3}))

	fmt.Println(find_max_recur([]int{1, 2, 3, 4, 5, 10, 20}))
	fmt.Println(find_max_recur([]int{-1, -2, -3}))
	fmt.Println(find_max_recur([]int{3, 2, 1}))
	fmt.Println(find_max_recur([]int{-1, 9, 3, 2, 30, 1, 10}))
}

func find_max_recur(arr []int) int {

	if len(arr) == 0 {
		return 0
	}

	if len(arr) == 1 {
		return arr[0]
	}

	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		}
		return arr[1]
	}

	sub_slice_max := find_max_recur(arr[1:])
	if arr[0] > sub_slice_max {
		return arr[0]
	}
	return sub_slice_max
}

func find_max(arr []int) int {
	var result int
	/*for _, item := range arr {
		if item > result {
			result = item
		}
	}*/

	for i, item := range arr {
		if item > result || i == 0 {
			result = item
		}
	}
	return result
}
