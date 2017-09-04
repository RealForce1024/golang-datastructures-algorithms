package sort

func quickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	less := []int{}
	grater := []int{}
	for _, v := range arr[1:] {
		switch {
		case v < pivot:
			less = append(less, v)
		case v > pivot:
			grater = append(grater, v)
		}
	}

	//return quickSort(less) + []int{pivot} + quickSort(grater)
	//return append(quickSort(less),pivot)+quickSort(grater)
	return append(append(quickSort(less), pivot), quickSort(grater)...)
}
