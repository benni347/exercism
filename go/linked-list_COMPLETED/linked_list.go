package linkedlist

import "errors"

// Node represents a node in the doubly linked list.
type Node struct {
	Value interface{}
	next  *Node
	prev  *Node
}

// List represents the doubly linked list.
type List struct {
	head *Node
	tail *Node
}

// NewList creates a new List with the given elements.
func NewList(args ...interface{}) *List {
	list := &List{}
	for _, arg := range args {
		list.Push(arg)
	}
	return list
}

// Next returns the next node.
func (n *Node) Next() *Node {
	return n.next
}

// Prev returns the previous node.
func (n *Node) Prev() *Node {
	return n.prev
}

// insertNodeAfter inserts a new node after a given node.
func (l *List) insertNodeAfter(node *Node, newNode *Node) {
	newNode.prev = node
	newNode.next = node.next
	if node.next != nil {
		node.next.prev = newNode
	}
	node.next = newNode
	if l.tail == node {
		l.tail = newNode
	}
}

// insertNodeBefore inserts a new node before a given node.
func (l *List) insertNodeBefore(node *Node, newNode *Node) {
	newNode.next = node
	newNode.prev = node.prev
	if node.prev != nil {
		node.prev.next = newNode
	}
	node.prev = newNode
	if l.head == node {
		l.head = newNode
	}
}

// removeNode removes a node from the list.
func (l *List) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	} else {
		l.head = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	} else {
		l.tail = node.prev
	}
}

// Unshift inserts a new value at the front of the list.
func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}
	if l.head == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.insertNodeBefore(l.head, newNode)
	}
}

// Push inserts a new value at the back of the list.
func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}
	if l.tail == nil {
		l.head = newNode
		l.tail = newNode
	} else {
		l.insertNodeAfter(l.tail, newNode)
	}
}

// Shift removes and returns the first value of the list.
func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("list is empty")
	}
	value := l.head.Value
	l.removeNode(l.head)
	return value, nil
}

// Pop removes and returns the last value of the list.
func (l *List) Pop() (interface{}, error) {
	if l.tail == nil {
		return nil, errors.New("list is empty")
	}
	value := l.tail.Value
	l.removeNode(l.tail)
	return value, nil
}

// First returns the first node of the list.
func (l *List) First() *Node {
	return l.head
}

// Last returns the last node of the list.
func (l *List) Last() *Node {
	return l.tail
}

// Reverse reverses the entire list.
func (l *List) Reverse() {
	current := l.head
	var temp *Node
	l.tail = current

	for current != nil {
		temp = current.prev
		current.prev = current.next
		current.next = temp
		current = current.prev
	}

	if temp != nil {
		l.head = temp.prev
	}
}
