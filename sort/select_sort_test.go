package sort

import (
	"fmt"
	"testing"
)

func Test_FindSmallest(t *testing.T) {
	var arr = []int{3, 2, 12, 7, 4}
	smallest_idx := FindSmallest(arr)
	fmt.Println("smallest=>", smallest_idx)
	if arr[smallest_idx] != 2 {
		t.Error("sort 有误!")
	}
}

func TestSelectionSort(t *testing.T) {
	arr := []int{3, 2, 12, 7, 4}
	new_arr := []int{2, 3, 4, 7, 12}
	result_arr := selectionSort(arr)
	if len(new_arr) != len(result_arr) {
		t.Error("ok")
	}

	fmt.Println("result_arr",result_arr)
}
