package main

import "fmt"

type ElementType int // 元素类型，假设为int

const MAXSIZE = 10 // 存储空间初始分配量

type SeqList struct {
	// 数组存储数据元素，最大为maxsize。  线性表的最大容量maxsize。存储空间的起始位置就是数组的存储位置。
	// 线性表的长度小于等于数组的长度
	data   [MAXSIZE]ElementType
	length int // 线性表的当前长度
}

func main() {
	var arr = [MAXSIZE]ElementType{1, 2, 3}
	var sqList = SeqList{data: arr, length: 10}
	fmt.Println(sqList)
}
