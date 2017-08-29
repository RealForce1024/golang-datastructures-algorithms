package sort

import "testing"

//func binary_search_test(t *testing.T) {
func Test_binary_search(t *testing.T) {
	result := binary_search([]int{1, 3, 5, 7}, 3)
	if result != true {
		t.Errorf("算法出错，应该找到")
	}
}

func Benchmark_binary_search(b *testing.B) {
	arr := []int{1, 3, 5, 7, 9, 11}
	for i := 0; i < b.N; i++ {
		result := binary_search(arr, 9)
		if !result {
			b.Errorf("benchmark 错误")
		}
	}

}
