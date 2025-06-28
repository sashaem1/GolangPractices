package list

import (
	"strconv"
)

type Node struct {
	val  int
	prev *Node
	next *Node
}

type List struct {
	head  *Node
	tail  *Node
	count int
}

func NewList(values ...int) *List {
	list := &List{}
	var prev *Node
	for _, v := range values {
		node := &Node{val: v, prev: prev}
		if prev != nil {
			prev.next = node
		} else {
			list.head = node
		}
		prev = node
		list.count++
	}
	list.tail = prev
	return list
}

func NewListFillFromSlice(slice []int) *List {
	return NewList(slice...)
}

func (l *List) Append(val int) {
	if l.head == nil {
		l.head = &Node{val: val}
		l.tail = l.head
		l.count++
	} else {
		l.tail.next = &Node{val: val, prev: l.tail, next: l.head}
		l.head.prev = l.tail.next
		l.tail = l.tail.next
		l.count++
	}
}

func (l *List) GetCount() int {
	return l.count
}

func (l *List) GetHead() int {
	return l.head.val
}

func (l *List) GetTail() int {
	return l.tail.val
}

func (l *List) Get(number int) int {
	tempNode := l.head
	for i := 0; i < number; i++ {
		tempNode = tempNode.next
	}

	return tempNode.val
}

func (l *List) ListPrintString() string {
	tempNode := l.head
	resultStr := "["
	for i := 0; i < l.count; i++ {
		resultStr += strconv.Itoa(tempNode.val)
		if i != l.count-1 {
			resultStr += ", "
		}
		tempNode = tempNode.next
	}
	resultStr += "]"
	return resultStr
}

func IsListEqual(List1, list2 *List) bool {
	if List1.count != list2.count {
		return false
	}

	tmpList1, tmpList2 := List1.head, list2.head

	for i := 0; i < List1.count; i++ {
		if tmpList1.val != tmpList2.val {
			return false
		}
		tmpList1, tmpList2 = List1.head.next, list2.head.next
	}

	return true
}
