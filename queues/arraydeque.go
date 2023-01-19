package queues

import (
	"fmt"
	"github.com/olaConge/go-datastructure/collections"
	"github.com/olaConge/go-datastructure/lists"
	"strings"
)

type ArrayDeque[T any] struct {
	list *lists.ArrayList[T]
}

func NewArrayDeque[T any]() *ArrayDeque[T] {
	return &ArrayDeque[T]{list: lists.NewArrayList[T](nil)}
}

var _ collections.Deque[any] = new(ArrayDeque[any])
var _ collections.Queue[any] = new(ArrayDeque[any])
var _ collections.Stack[any] = new(ArrayDeque[any])

func (queue *ArrayDeque[T]) OfferFirst(value T) {
	if queue.list.Size() == 0 {
		queue.list.Add(value)
	} else {
		_ = queue.list.Insert(0, value)
	}
}

func (queue *ArrayDeque[T]) OfferLast(value T) {
	queue.Offer(value)
}

func (queue *ArrayDeque[T]) PollFirst() (value T, ok bool) {
	value, err := queue.list.Remove(0)
	if err != nil {
		ok = false
	}
	return
}

func (queue *ArrayDeque[T]) PollLast() (value T, ok bool) {
	value, err := queue.list.Remove(queue.list.Size() - 1)
	if err != nil {
		ok = false
	}
	return
}

func (queue *ArrayDeque[T]) PeekFirst() (value T, ok bool) {
	value, err := queue.list.Get(0)
	if err != nil {
		ok = false
	}
	return
}

func (queue *ArrayDeque[T]) PeekLast() (value T, ok bool) {
	value, err := queue.list.Get(queue.list.Size() - 1)
	if err != nil {
		ok = false
	}
	return
}

func (queue *ArrayDeque[T]) Offer(value T) {
	queue.list.Add(value)
}

func (queue *ArrayDeque[T]) Poll() (value T, ok bool) {
	return queue.PollFirst()
}

func (queue *ArrayDeque[T]) Peek() (value T, ok bool) {
	return queue.PeekFirst()
}

func (queue *ArrayDeque[T]) Push(value T) {
	queue.OfferFirst(value)
}

func (queue *ArrayDeque[T]) Pop() (value T, ok bool) {
	return queue.PollFirst()
}

func (queue *ArrayDeque[T]) Empty() bool {
	return queue.list.Empty()
}

func (queue *ArrayDeque[T]) Size() int {
	return queue.list.Size()
}

func (queue *ArrayDeque[T]) Clear() {
	queue.list.Clear()
}

func (queue *ArrayDeque[T]) Values() []T {
	return queue.list.Values()
}

func (queue *ArrayDeque[T]) String() string {
	str := "ArrayDeque\n"
	values := []string{}
	for _, value := range queue.list.Values() {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
