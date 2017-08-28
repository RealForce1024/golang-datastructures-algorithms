package main

import "fmt"

func main() {
	arr := []int{1, 3, 5, 7}
	binary_search(arr, 3)
}

func binary_search(arr []int, item int) (result int) {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		fmt.Println(guess)
		switch {
		case guess == item:
			result = mid
			fmt.Println("find it")
			return
		case guess > item:
			high = mid - 1
			fmt.Println("大了")
		case guess < item:
			low = mid + 1
			fmt.Println("小了")
		}
	}
	return
}
