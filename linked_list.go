package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head *node
	length int
}

func (l *LinkedList) prepend(n *node) {
	oldHead := l.head
	n.next = oldHead
	l.head = n
	l.length++
}

func (l *LinkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return
	}

	prevToDelete := l.head
	for prevToDelete.next.data != value {
		if prevToDelete.next.next == nil {
			return
		}
		prevToDelete = prevToDelete.next
	}
	prevToDelete.next = prevToDelete.next.next
	l.length--
}

func (l LinkedList) printLinkedList() {
	for l.length != 0 {
		fmt.Printf("%d ", l.head.data)
		l.head = l.head.next
		l.length--
	}

	fmt.Println("")
}

func main() {
	node1 := &node{data: 3}
	node2 := &node{data: 5}
	node3 := &node{data: 7}
	node4 := &node{data: 9}

	myLinkedList := LinkedList{}
	myLinkedList.prepend(node1)
	myLinkedList.prepend(node2)
	myLinkedList.prepend(node3)
	myLinkedList.prepend(node4)

	myLinkedList.printLinkedList()

	myLinkedList.deleteWithValue(5)
	myLinkedList.printLinkedList()
}