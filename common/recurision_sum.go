package main

import "fmt"

func sum(arr []int) int {
	var result int
	for _, item := range arr {
		result += item
	}
	return result
}

func main() {
	fmt.Println(sum([]int{1, 2, 3, 4, 5}))
	fmt.Println(recurisionSum([]int{1, 2, 3, 4, 5}))
}

func recurisionSum(arr []int) int {
	var sum int
	if len(arr) == 0 {
		return sum
	}
	//return arr[0] + recurisionSum(arr[:len(arr)])//runtime: goroutine stack exceeds 1000000000-byte limit
	//fatal error: stack overflow
	//return arr[0] + recurisionSum(arr[:len(arr)-1])
	//return arr[0] + recurisionSum(arr[1:len(arr)-1]) 注意arr[1:len(arr)-1]会下届超标，也就是不要减一了
	return arr[0] + recurisionSum(arr[1:len(arr)])
}
