package main

import "fmt"

func main() {
	fmt.Println(sum_elem([]int{1, 2, 3, 4, 5,8}))
	fmt.Println(recurisive_count([]int{1, 2, 3, 4, 5,8}))
}

func recurisive_count(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	return len(arr[:1]) + recurisive_count(arr[1:len(arr)])
}

func sum_elem(arr []int) int {
	var result int
	for i := 0; i < len(arr); i++ {
		result ++
	}
	return result
}
