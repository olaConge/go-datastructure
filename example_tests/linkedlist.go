package main

import (
	"github.com/olaConge/go-datastructure/lists"
	"github.com/olaConge/go-datastructure/utils"
)

func main() {
	l := lists.NewLinkedList[int](utils.Equals[int], 1, 2, 3, 4, 5, 6, 7)
	g, _ := l.Get(0)
	_ = g

	l.Add(10, 9, 8)
	v := l.Values()
	_ = v

	l.Sort(func(a, b int) int { return a - b })
	v = l.Values()

	_, _ = l.Set(0, 100)
	v = l.Values()

	b := l.Contains(100)
	b = l.Contains(100, 200)
	_ = b

	_ = l.Insert(4, 200, 300, 400)
	v = l.Values()

	_ = l.Swap(4, 5)
	v = l.Values()

	_, _ = l.Remove(4)
	v = l.Values()

	_ = l.Size()
	b = l.Empty()
	s := l.String()
	v = l.Values()
	_ = s
	l.Clear()
	b = l.Empty()
}
