package lists

import "II/lists/static"

func Test() {
	list := static.NewStaticList[*static.StaticList[int]](11)
	for i := 0; i < 10; i++ {
		item := static.NewStaticList[int](10)
		list.Add(item)
	}
	list.At(0).Add(1)
	list.At(0).Add(2)
	list.At(0).Print()
}
