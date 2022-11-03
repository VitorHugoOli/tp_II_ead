package linked

// LinkedList is a linked list withouth a head
type LinkedListFull[T comparable] struct {
	Head  *Node[T]
	First *Node[T]
	Last  *Node[T]

	Len int
}

// NewLinkedListFull
func NewLinkedListFull[T comparable]() *LinkedListFull[T] {
	linked := &LinkedListFull[T]{
		Head:  &Node[T]{Next: nil},
		First: nil,
		Last:  nil,
		Len:   0,
	}

	linked.Last = linked.Head
	linked.First = linked.Head

	return linked
}

func (l *LinkedListFull[T]) Add(element T) {
	l.Last.Next = &Node[T]{Data: element}
	l.Last = l.Last.Next
	l.Len++
}

func (l *LinkedListFull[T]) Remove(element T) {
	l.remove(element, l.First)
}

func (l *LinkedListFull[T]) remove(element T, node *Node[T]) {
	if node == nil {
		panic("Element not found")
	}
	if node.Next == nil {
		l.Last = node
	}
	if node.Next.Data == element {
		node.Next = node.Next.Next
		l.Len--
		return
	}
	l.remove(element, node.Next)
}

func (l *LinkedListFull[T]) Search(element T) bool {
	return l.search(element, l.First)
}

func (l *LinkedListFull[T]) search(element T, node *Node[T]) bool {
	if node == nil {
		return false
	}
	if node.Next.Data == element {
		return true
	}
	return l.search(element, node.Next)
}

// At
func (l *LinkedListFull[T]) At(index int) T {
	return l.at(index, l.First)
}

func (l *LinkedListFull[T]) at(index int, node *Node[T]) T {
	if index == 0 {
		return node.Data
	}
	return l.at(index-1, node.Next)
}

// print valuie -> value
func (l *LinkedListFull[T]) Print() {
	l.print(l.First)
}

func (l *LinkedListFull[T]) print(node *Node[T]) {
	if node == nil {
		println()
	}
	print(node.Data, " -> ")
	l.print(node.Next)
}
