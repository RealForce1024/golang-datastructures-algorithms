package main

import "fmt"

type ElementType int // 元素类型，假设为int

const MAXSIZE = 10 // 存储空间初始分配量

type SeqList struct {
	// 数组存储数据元素，最大为maxsize。  线性表的最大容量maxsize。存储空间的起始位置就是数组的存储位置。
	// 线性表的长度<=数组的长度
	data   [MAXSIZE]ElementType
	length int // 线性表的当前长度
}

func (list SeqList) getElem(i int) (ElementType, error) {
	if len(list.data) == 0 {
		panic("顺序表数组为空,请添加元素后再进行获取指定位置的元素")
	}

	if i < 1 || i > MAXSIZE {
		panic("参数有误，请传递正确范围1-MaxSize")
	}

	return list.data[i-1], nil
}

func main() {
	// var arr = [MAXSIZE]ElementType{1, 2, 3}
	var arr = [MAXSIZE]ElementType{}
	var sqList = SeqList{data: arr, length: 10}
	fmt.Println(sqList)

	fmt.Println(sqList.getElem(11))
}
