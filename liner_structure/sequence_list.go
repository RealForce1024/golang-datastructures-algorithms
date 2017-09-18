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
	if list.length == 0 {
		panic("顺序表数组为空,请添加元素后再进行获取指定位置的元素")
	}

	if i < 1 || i > list.length {
		panic("参数有误，请传递正确范围1~length")
	}

	return list.data[i-1], nil
}

func (list *SeqList) addElem(i int, e ElementType) { //注意传递指针
	if list.length == MAXSIZE {
		panic("顺序表已满")
	}

	// if i < 1 || i > MAXSIZE { //注意:实际的长度与容量，添加元素看实际的长度
	if i < 1 || i > list.length+1 { //
		panic("要添加的元素不在已有位置的位置内")
	}

	for k := list.length - 1; k >= i-1; k-- {
		list.data[k+1] = list.data[k]
	}

	list.data[i-1] = e
	list.length++
}
func main() {
	var arr = [MAXSIZE]ElementType{1, 2, 3}
	// var arr = [MAXSIZE]ElementType{}
	// var arr [MAXSIZE]ElementType
	var sqList = SeqList{data: arr}

	// sqList.addElem(2, 33)
	sqList.addElem(1, 11)
	sqList.addElem(2, 22)
	fmt.Println(sqList.length)
	fmt.Println(sqList)
	fmt.Println(sqList.getElem(2))

}
