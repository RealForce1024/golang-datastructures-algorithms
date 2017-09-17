package main

import "fmt"

func main() {
	// s := []int{1, 2, 3, 4, 5}
	// printArr(reverse1(s)) // 利用对称角标和的数学关系交换

	// s := []int{1, 2, 3, 4, 5}
	// printArr(reverse2(s)) //利用go语法的便利,多变量声明,直接交换位置

	s := []int{1, 2, 3, 4, 5}
	printArr(reverse3(s)) //逼近法 从左到右

}

func reverse1(s []int) []int {
	length := len(s)

	for i := 0; i < length; i++ {
		j := length - i - 1
		s[j] = s[i]
	}

	return s
}

func printArr(s []int) {
	for i, v := range s {
		fmt.Println(i, v)
	}
}

func reverse2(s []int) []int {
	length := len(s)

	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}

func reverse3(s []int) []int {
	length := len(s)
	left := 0
	right := length - 1

	for left < right {
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

	return s
}
