package main

import (
	"fmt"
	"unicode/utf8"
)

func deformation(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	return true

}

func main() {
	str1 := "hello world"
	fmt.Println(str1)
	fmt.Printf("%s,%d", str1, len(str1))
	str2 := "abcd 中国"
	fmt.Printf("%s,%d\n", str2, len(str2))
	len_str2 := utf8.RuneCountInString(str2)
	fmt.Println(len_str2)
	arr2 := []rune(str2)

	for i := 0; i < len_str2; i++ {
		fmt.Println(arr2[i])
	}
}
