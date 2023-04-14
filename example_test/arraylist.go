package main

import (
	"github.com/olaConge/go-datastructure/list"
	"github.com/olaConge/go-datastructure/util"
)

func main() {
	l := list.NewArrayList[int](util.Equals[int], 1, 2, 3, 4, 5, 6, 7)
	_, _ = l.Get(0)
	l.Add(10, 9, 8)
	l.Sort(func(a, b int) int { return a - b })
	_, _ = l.Set(0, 100)
	_ = l.Contains(100)
	_ = l.Contains(100, 200)
	_ = l.Insert(4, 200, 300, 400)
	_ = l.Swap(4, 5)
	_, _ = l.Remove(4)
	_ = l.Size()
	_ = l.Empty()
	_ = l.String()
	_ = l.Values()
	l.Clear()
	_ = l.Empty()
}
