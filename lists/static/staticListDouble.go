package static

import "fmt"

type StaticListDouble[T ListTypes] struct {
	Data []T
	Next []int
	Back []int
	Len  int

	First int
	Last  int
	Free  int
}

// StaticListDouble Contiguous dynamic list with a pointer to the next element
func NewStaticListDouble[T ListTypes](size int) *StaticListDouble[T] {
	return &StaticListDouble[T]{
		Data:  make([]T, size),
		Next:  make([]int, size),
		Back:  make([]int, size),
		Len:   0,
		First: -1,
		Last:  -1,
		Free:  0,
	}
}

// Add, Remove, Search, At, Print
func (l *StaticListDouble[T]) Add(element T) {
	if l.Len == len(l.Data) {
		panic("List is full")
	}
	if l.First == -1 {
		l.First = l.Free
		l.Last = l.Free
	} else {
		l.Next[l.Last] = l.Free
		l.Back[l.Free] = l.Last
		l.Last = l.Free
	}
	l.Data[l.Free] = element
	l.Free = l.Next[l.Free]
	l.Len++
}

func (l *StaticListDouble[T]) Remove(element T) {
	if l.First == -1 {
		panic("List is empty")
	}
	if l.Data[l.First] == element {
		l.First = l.Next[l.First]
		l.Len--
		return
	}
	l.remove(element, l.First)
}

func (l *StaticListDouble[T]) remove(element T, node int) {
	if l.Next[node] == -1 {
		panic("Element not found")
	}
	if l.Data[l.Next[node]] == element {
		l.Next[node] = l.Next[l.Next[node]]
		l.Len--
		return
	}
	l.remove(element, l.Next[node])
}

func (l *StaticListDouble[T]) Search(element T) bool {
	return l.search(element, l.First)
}

func (l *StaticListDouble[T]) search(element T, node int) bool {
	if node == -1 {
		return false
	}
	if l.Data[node] == element {
		return true
	}
	return l.search(element, l.Next[node])
}

func (l *StaticListDouble[T]) At(index int) T {
	return l.at(index, l.First)
}

func (l *StaticListDouble[T]) at(index int, node int) T {
	if index == 0 {
		return l.Data[node]
	}
	return l.at(index-1, l.Next[node])
}

func (l *StaticListDouble[T]) Print() {
	l.print(l.First)
}

func (l *StaticListDouble[T]) print(node int) {
	if node == -1 {
		return
	}
	fmt.Println(l.Data[node])
	l.print(l.Next[node])
}
