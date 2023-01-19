package collections

type Queue[T any] interface {
	Offer(value T)
	Poll() (value T, ok bool)
	Peek() (value T, ok bool)
	Collection[T]
}

type Deque[T any] interface {
	OfferFirst(value T)
	OfferLast(value T)
	PollFirst() (value T, ok bool)
	PollLast() (value T, ok bool)
	PeekFirst() (value T, ok bool)
	PeekLast() (value T, ok bool)
	Queue[T]
	Stack[T]
}
