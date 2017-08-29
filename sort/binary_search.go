package sort

func binary_search(arr []int, item int) bool {
	low := 0
	high := len(arr) - 1
	for low <= high {
		mid := (low + high) / 2
		guess := arr[mid]
		switch {
		case guess == item:
			return true
		case guess > item:
			high = mid - 1
		case guess < item:
			low = mid + 1
		}
	}
	return false
}
