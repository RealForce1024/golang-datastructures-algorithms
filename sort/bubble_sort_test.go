package sort

import (
	"testing"
	"fmt"
)

func Test_bubble_sort(t *testing.T) {
	arr := []int{3,2,1,4,0,8}
	sorted_arr := bubble_sort(arr)
	fmt.Println(sorted_arr)
	if sorted_arr[0]!=0{
		t.Error("bubble sort error")
	}
}
	