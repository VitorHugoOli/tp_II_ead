package linked

import "fmt"

// Node
type Node[T comparable] struct {
	Data T
	Next *Node[T]
}

// LinkedList is a linked list withouth a head
type LinkedList[T comparable] struct {
	First *Node[T]
	Len   int
}

// NewLinkedList
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		First: nil,
		Len:   0,
	}
}

func (l *LinkedList[T]) Add(element T) {
	if l.First == nil {
		l.First = &Node[T]{Data: element}
	} else {
		l.First = &Node[T]{Data: element, Next: l.First}
	}
	l.Len++
}

func (l *LinkedList[T]) Remove(element T) {
	l.remove(element, l.First)
}

func (l *LinkedList[T]) remove(element T, node *Node[T]) {
	if node == nil {
		panic("Element not found")
	}
	if node.Next == nil {
		panic("Element not found")
	}
	if node.Next.Data == element {
		node.Next = node.Next.Next
		l.Len--
		return
	}
	l.remove(element, node.Next)
}

func (l *LinkedList[T]) Search(element T) bool {
	return l.search(element, l.First)
}

func (l *LinkedList[T]) search(element T, node *Node[T]) bool {
	if node == nil {
		return false
	}
	if node.Data == element {
		return true
	}
	return l.search(element, node.Next)
}

func (l *LinkedList[T]) At(index int) T {
	return l.at(index, l.First)
}

func (l *LinkedList[T]) at(index int, node *Node[T]) T {
	if index == 0 {
		return node.Data
	}
	return l.at(index-1, node.Next)
}

func (l *LinkedList[T]) Print() {
	l.print(l.First)
}

func (l *LinkedList[T]) print(node *Node[T]) {
	if node == nil {
		return
	}
	fmt.Println(node.Data)
	l.print(node.Next)
}
