package collection

import "fmt"

type node struct {
	data   interface{}
	next   *node
	before *node
}

type linkedList struct {
	head   *node
	tail   *node
	length int
}

func (l *linkedList) keepPrepend(n *node) {
	if l.length == 0 {
		l.head = n
		l.head.before = nil
		l.head.next = nil
		l.tail = n
	} else {
		nextnode := l.head
		l.head = n
		l.head.before = nil
		l.head.next = nextnode
		nextnode.before = l.head
		l.tail = nextnode
	}
	l.length++
}

func (l *linkedList) keepAppend(n *node) {
	if l.length == 0 {
		l.head = n
		l.tail = n
	} else {
		nextnode := l.head
		for nextnode.next != nil {
			nextnode = nextnode.next
		}
		nextnode.next = n
		n.before = nextnode
		l.tail = n
	}
	l.length++
}

func (l *linkedList) printLinkedList() {
	nextnode := l.head
	for nextnode.next != nil {
		fmt.Println("Data: ", nextnode.data)
		nextnode = nextnode.next
		// print the last node value
		if nextnode.next == nil {
			fmt.Println("Data: ", nextnode.data)
		}
	}
}

func (l *linkedList) reversePrintLinkList() {
	beforelastnode := l.tail
	for beforelastnode.before != nil {
		fmt.Println("Data: ", beforelastnode.data)
		beforelastnode = beforelastnode.before
		// print the first node value
		if beforelastnode.before == nil {
			fmt.Println("Data: ", beforelastnode.data)
		}
	}
}

func (l *linkedList) popTail() *node {
	nextnode := l.head
	for nextnode.next != nil {
		nextnode = nextnode.next
		if nextnode.next == nil {
			newLastNode := nextnode.before
			newLastNode.next = nil
			l.length--
		}
	}
	lastNode := nextnode
	return lastNode
}

func (l *linkedList) popHead() *node {
	popnode := l.head
	nextnode := l.head.next
	nextnode.before = nil
	l.head = nextnode
	l.length--
	return popnode
}

func main() {
	node0 := &node{data: 1}
	node1 := &node{data: 2}
	node2 := &node{data: 3}
	popNode := &node{data: 0}

	fmt.Println("Append linked list")
	appendLinkList := linkedList{}
	appendLinkList.keepAppend(node0)
	appendLinkList.keepAppend(node1)
	appendLinkList.keepAppend(node2)
	appendLinkList.printLinkedList()
	fmt.Println("Append reverse linkedList")
	appendLinkList.reversePrintLinkList()
	popNode = appendLinkList.popHead()
	fmt.Println("Pop Node Data: ", popNode.data)
	fmt.Println("After popping node...")
	appendLinkList.printLinkedList()

	fmt.Println("Prepend linked list")
	prependLinkList := linkedList{}
	prependLinkList.keepPrepend(node0)
	prependLinkList.keepPrepend(node1)
	prependLinkList.keepPrepend(node2)
	prependLinkList.printLinkedList()
	fmt.Println("Prepend reverse linkedList")
	prependLinkList.reversePrintLinkList()
	popNode = prependLinkList.popHead()
	fmt.Println("Pop Node Data: ", popNode.data)
	fmt.Println("After popping node...")
	prependLinkList.printLinkedList()
}
