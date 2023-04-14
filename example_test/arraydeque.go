package main

import (
	"github.com/olaConge/go-datastructure/queue"
)

func main() {
	q := queue.NewArrayDeque[int]()
	q.Offer(1)
	q.OfferLast(2)
	q.OfferFirst(3)

	v, _ := q.Peek()
	v, _ = q.PeekLast()
	v, _ = q.PeekFirst()

	v, _ = q.Poll()
	v, _ = q.PollFirst()
	v, _ = q.PollLast()

	q.Push(4)
	q.Push(5)
	v, _ = q.Pop()

	_ = v
}
