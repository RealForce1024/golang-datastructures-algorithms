package main

import "fmt"

type SinglyLinkListNode struct {
	data int
	next *SinglyLinkListNode
}

func (l *SinglyLinkListNode) show() {
	curr := l
	for curr != nil {
		fmt.Print(curr.data)
		if curr.next != nil {
			fmt.Print("->")
		}
		curr = curr.next
	}
}

func (head *SinglyLinkListNode) reverse_recurise() *SinglyLinkListNode {
	if head == nil || head.next == nil {
		return head
	}
	// 思路: 只关心第一个和第二个(递归调用产生的节点)的关系，递归后重建second节点与旧的head的节点关系
	second := head.next
	new_head := second.reverse_recurise()

	second.next = head
	head.next = nil
	return new_head
}

func (l *SinglyLinkListNode) reverse() *SinglyLinkListNode {
	curr := l
	var prew *SinglyLinkListNode
	for curr != nil {
		pnext := curr.next
		curr.next = prew
		if pnext == nil {
			break
		}
		prew = curr
		curr = pnext
	}
	return curr
}

func main() {
	list := &SinglyLinkListNode{data: 1, next: &SinglyLinkListNode{data: 2, next: &SinglyLinkListNode{data: 3}}}
	list.show()
	fmt.Println()
	// list.reverse().show()
	list.reverse_recurise().show()

}
