package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkList Node

func main() {
	var linkList LinkList
	fmt.Println(linkList)
}
