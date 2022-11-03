package static

import "II/analysis"

type StaticList[T ListTypes] struct {
	Data []T
	Len  int
}

// NewStaticList creates a new list of fixed size
func NewStaticList[T ListTypes](size int) *StaticList[T] {
	return &StaticList[T]{
		Data: make([]T, size),
		Len:  0,
	}
}

// Add adds an element to the list
func (l *StaticList[T]) Add(element T) {
	l.Data[l.Len] = element
	l.Len++
}

// Remove removes an element from the list
func (l *StaticList[T]) Remove(element T) {
	for i := 0; i < l.Len; i++ {
		analysis.Analysis.CurrentMeasure.Iterations++
		if l.Data[i] == element {
			l.Data[i] = l.Data[l.Len-1]
			l.Len--
			break
		}
	}
}

// Search searches for an element in the list
func (l *StaticList[T]) Search(element T) bool {
	for i := 0; i < l.Len; i++ {
		analysis.Analysis.CurrentMeasure.Iterations++
		if l.Data[i] == element {
			return true
		}
	}
	return false
}

// At returns the element at the given index
func (l *StaticList[T]) At(index int) T {
	if index < 0 || index >= l.Len {
		analysis.Analysis.CurrentMeasure.Iterations++
		panic("Index out of bounds")
	}
	return l.Data[index]
}

// Print prints the list
func (l *StaticList[T]) Print() {
	for i := 0; i < l.Len; i++ {
		print(l.Data[i], " ")
	}
	println()
}
