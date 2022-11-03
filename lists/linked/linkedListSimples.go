package linked

// LinkedList is a linked list withouth a head
type LinkedListSimple[T comparable] struct {
	First *Node[T]
	Last  *Node[T]

	Len int
}

// NewLinkedListSimple
func NewLinkedListSimple[T comparable]() *LinkedListSimple[T] {
	return &LinkedListSimple[T]{
		First: nil,
		Last:  nil,
		Len:   0,
	}
}

func (l *LinkedListSimple[T]) Add(element T) {
	if l.First == nil {
		l.First = &Node[T]{Data: element}
		l.Last = l.First
	} else {
		l.Last.Next = &Node[T]{Data: element}
		l.Last = l.Last.Next
	}
	l.Len++
}

func (l *LinkedListSimple[T]) Remove(element T) {
	l.remove(element, l.First)
}

func (l *LinkedListSimple[T]) remove(element T, node *Node[T]) {
	if node == nil {
		panic("Element not found")
	}
	if node.Next == nil {
		l.Last = node
	}
	if node.Data == element {
		node.Data = node.Next.Data
		node.Next = node.Next.Next
		l.Len--
		return
	}
	l.remove(element, node.Next)
}

func (l *LinkedListSimple[T]) Search(element T) bool {
	return l.search(element, l.First)
}

func (l *LinkedListSimple[T]) search(element T, node *Node[T]) bool {
	if node == nil {
		return false
	}
	if node.Data == element {
		return true
	}
	return l.search(element, node.Next)
}
