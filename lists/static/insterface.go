package static

type ListTypes interface {
	comparable
}

type List interface {
	New() List
	Add(value interface{})
	Remove(value interface{})
	Search(value interface{}) bool
	At()
	Print()
}
