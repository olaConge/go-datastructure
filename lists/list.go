package lists

import (
	"github.com/olaConge/go-datastructure/collections"
	"github.com/olaConge/go-datastructure/utils"
)

type List[T any] interface {
	Add(values ...T)
	Get(index int) (oldValue T, err error)
	Set(index int, value T) (oldValue T, err error)
	Remove(index int) (oldValue T, err error)
	Contains(values ...T) bool
	Insert(index int, values ...T) error
	Sort(comparator utils.Comparator[T])
	Swap(index1, index2 int) error

	collections.Collection[T]
}
