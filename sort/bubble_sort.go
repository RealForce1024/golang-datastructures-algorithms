package sort

func bubble_sort(arr []int) []int {
	length := len(arr)
	for i := 0; i < length; i++ {
		for j := i + 1; j < length - 1; j++ {
			if arr[j] < arr[i] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}
