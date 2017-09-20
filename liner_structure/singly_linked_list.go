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

	second := head.next
	new_head := second.reverse_recurise()

	second.next = head
	head.next = nil
	return new_head
}

func (l *SinglyLinkListNode) reverse() *SinglyLinkListNode {
	curr := l
	// 用于记录当前结点的前驱结点
	// 前驱结点开始为nil，因为是反转后的最后一个结点的下一个结点，即nil
	var prew *SinglyLinkListNode
	for curr != nil {
		// 记录当前节点的下一节点
		next := curr.next
		// 将当前节点的下一节点指向前驱节点。也就是将当前节点插入到了反转链表的头部
		curr.next = prew
		if next == nil {
			break
		}
		// 记录当前节点为前驱节点
		prew = curr
		// 当前节点移动到下一个节点
		curr = next
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
