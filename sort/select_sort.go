package sort

import "fmt"

func FindSmallest(arr []int) int {
	var smallest = arr[0]
	var smallest_idx int
	for i := 0; i < len(arr); i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallest_idx = i
		}
	}
	fmt.Println("最小idx:=>", smallest_idx)
	return smallest_idx
}

func selectionSort(arr []int) []int {
	var new_arr []int
	ori_len := len(arr)

	for i := 0; i < ori_len; i++ {
		smallest_idx := FindSmallest(arr)
		new_arr = append(new_arr, arr[smallest_idx])
		//delete(arr, smallest_idx)
		arr = append(arr[:smallest_idx], arr[smallest_idx+1:]...)
	}

	return new_arr
}
