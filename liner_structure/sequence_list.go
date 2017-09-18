package main

import "fmt"

type ElementType int

const MAXSIZE = 3

type SeqList struct {
	data   [MAXSIZE]ElementType // 线性表的容量 数组存储数据元素，最大为maxsize
	length int                  // 线性表的当前长度
}

func main() {
	var arr = [MAXSIZE]ElementType{1, 2, 3}
	var sqList = SeqList{data: arr, length: 10}
	fmt.Println(sqList)
}
