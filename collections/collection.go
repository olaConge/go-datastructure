package collections

type Collection[T any] interface {
	Empty() bool
	Size() int
	Clear()
	Values() []T
	String() string
}
