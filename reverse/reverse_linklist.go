package main

type List struct {
	value int
	next  *List
}

func printList(l *List) {
	curr := l

	for curr != nil {
		print(curr.value)
		if curr.next != nil {
			print("->")
		}
		curr = curr.next
	}
}

func reverse(l *List) *List {
	curr := l
	var prev *List

	for curr != nil {
		//记录当前下一节点
		next := curr.next
		curr.next = prev
		if next == nil {
			break
		}
		prev = curr
		curr = next
	}
	return curr
}

func main() {
	list := &List{
		value: 1,
		next: &List{
			value: 2,
			next: &List{
				value: 3,
				next: &List{
					value: 4,
				},
			},
		},
	}

	printList(list)
	println()

	printList(reverse(list))

}
