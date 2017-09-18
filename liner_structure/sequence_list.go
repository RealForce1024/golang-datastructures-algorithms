package main

import "fmt"

type ElementType int

const MAXSIZE = 3

type SeqList struct {
	data   [MAXSIZE]ElementType // 线性表的容量 数组存储数据元素，最大为maxsize
	length int                  // 线性表的当前长度
}

func main() {
	var sqList = SeqList{length: 10}
	fmt.Println(sqList)
}
