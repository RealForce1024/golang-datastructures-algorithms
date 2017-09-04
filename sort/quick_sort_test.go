package sort

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/gomaster-me/go_commons"
)

func Test_quick_sort(t *testing.T) {
	s1 := quickSort([]int{3, 2, 5, 1, 8})
	s2 := []int{1, 2, 3, 5, 8}
	//s2 := []int{1, 2, 3, 5, 8, 9}
	//assert.Equal(t, s2, s1, "排序错误")

	result := go_commons.Compare(s1, s2)
	assert.Equal(t, result, true, "排序错误")

}
