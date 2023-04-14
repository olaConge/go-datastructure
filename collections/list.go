package collections

import (
	"github.com/olaConge/go-datastructure/util"
)

type List[T any] interface {
	Add(values ...T)
	Get(index int) (oldValue T, err error)
	Set(index int, value T) (oldValue T, err error)
	Remove(index int) (oldValue T, err error)
	Contains(values ...T) bool
	Insert(index int, values ...T) error
	Sort(comparator util.Comparator[T])
	Swap(index1, index2 int) error

	Collection[T]
}
