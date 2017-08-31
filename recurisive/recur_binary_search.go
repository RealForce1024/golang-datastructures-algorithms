package main

import "fmt"

func main() {
	fmt.Println(recur_binary_search([]int{1, 2, 3, 4, 5}, 29))
	fmt.Println(recur_binary_search([]int{1, 2, 3, 4, 5}, 2))
	fmt.Println(recur_binary_search([]int{1, 2, 3, 4, 5}, 3))
	fmt.Println(recur_binary_search([]int{1, 2, 3, 4, 5, -1}, -1))
}

func recur_binary_search(arr []int, item int) bool {
	if len(arr) == 0 {
		return false
	}
	//mid := (len(arr) - 1) //2  left+right 左右步步紧逼
	mid := len(arr) - 1
	switch {
	case arr[mid] == item:
		return true
	case item < arr[mid]:
		return recur_binary_search(arr[:mid], item) //binary
	case item > arr[mid]:
		return recur_binary_search(arr[mid+1:], item)
	}
	return false
}
