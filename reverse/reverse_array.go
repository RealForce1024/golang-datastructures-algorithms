package main

import "fmt"

func main() {
	// reverse1() // 利用对称角标和的数学关系交换
	// reverse2()//利用go语法的便利,多变量声明,直接交换位置
	reverse3() //逼近法 从左到右
}

func reverse1() {
	arr := [5]int{1, 2, 3, 4, 5}
	length := len(arr)
	arr_new := [5]int{} // arr_new := [length]int{} //non-constant array bound length

	for i := 0; i < length; i++ {
		j := length - i - 1
		arr_new[j] = arr[i]
	}

	for i, v := range arr_new {
		fmt.Println(i, v)
	}
}

func reverse2() {
	arr := [5]int{1, 2, 3, 4, 5}
	length := len(arr)

	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	for i, v := range arr {
		fmt.Println(i, v)
	}
}

func reverse3() {
	arr := [5]int{1, 2, 3, 4, 5}
	length := len(arr)
	left := 0
	right := length - 1

	for left < right {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}

	for i, v := range arr {
		fmt.Printf("i=%d,v=%v\n", i, v)
	}
}
