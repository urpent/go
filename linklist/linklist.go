package linklist

import (
	"fmt"
	"strings"
)

type Node[T any] struct {
	Data T
	next *Node[T]
	prev *Node[T]
}

type DoublyLinkedList[T any] struct {
	length int
	head   *Node[T]
	tail   *Node[T]
}

func NewDoublyLinkedList[T any]() DoublyLinkedList[T] {
	return DoublyLinkedList[T]{}
}

func (l *DoublyLinkedList[T]) AddFront(data T) (newNode *Node[T]) {
	defer func() { l.length++ }()

	newNode = &Node[T]{
		Data: data,
		next: l.head,
	}

	if l.head == nil {
		l.tail = newNode
	} else {
		l.head.prev = newNode
	}

	l.head = newNode

	return
}

func (l *DoublyLinkedList[T]) AddEnd(data T) (newNode *Node[T]) {
	defer func() {
		l.length++
	}()

	newNode = &Node[T]{
		Data: data,
	}

	if l.head == nil {
		l.head = newNode
		l.tail = newNode
		return
	}

	l.tail.next = newNode
	newNode.prev = l.tail
	l.tail = newNode

	return
}

func (l *DoublyLinkedList[T]) MoveNodeToFront(node *Node[T]) {
	if l.length == 0 || node == nil || node == l.head {
		return
	}

	l.Remove(node)

	l.length++
	l.head.prev = node
	node.next = l.head
	l.head = node
	return
}

func (l *DoublyLinkedList[T]) Remove(node *Node[T]) {
	if l.length == 0 || node == nil {
		return
	}

	defer func() {
		node.next = nil
		node.prev = nil
		l.length--
	}()

	// If this is the only Node
	if l.length == 1 {
		l.head = nil
		l.tail = nil
		return
	}

	if l.head == node { // If head Node
		// next Node become head
		node.next.prev = nil
		l.head = node.next
		return
	}

	if l.tail == node { // If tail Node
		l.tail = node.prev
		l.tail.next = nil
		return
	}

	// Node in the middle of the list
	node.prev.next = node.next
	node.next.prev = node.prev

	return
}

func (l *DoublyLinkedList[T]) RemoveLast() (removedNode *Node[T]) {
	removedNode = l.tail
	l.Remove(l.tail)

	return
}

func (l *DoublyLinkedList[T]) Len() int {
	return l.length
}

func (l *DoublyLinkedList[T]) string() string {
	var b strings.Builder
	b.Grow(l.length * 3)

	p := l.head
	for p != nil {
		fmt.Fprintf(&b, "%v -> ", p.Data)
		p = p.next
	}

	return strings.TrimSuffix(b.String(), " -> ")
}
