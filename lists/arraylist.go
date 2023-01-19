package lists

import (
	"fmt"
	"github.com/olaConge/go-datastructure/collections"
	"github.com/olaConge/go-datastructure/errors"
	"github.com/olaConge/go-datastructure/utils"
	"strings"
)

var _ collections.List[any] = new(ArrayList[any])

// ArrayList contains the following members:
// elements: a slice store elements of Type T
// size: how many elements there are
// equals: a function used to justify whether elements are equal
type ArrayList[T any] struct {
	elements []T
	size     int
	equals   func(a, b T) bool
}

const (
	expandFactor = float32(2.0)
	shrinkFactor = float32(0.25)
)

func NewArrayList[T any](equals utils.EqualsFn[T], values ...T) *ArrayList[T] {
	list := &ArrayList[T]{
		equals: equals,
	}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends the specified value to the end of this list.
func (list *ArrayList[T]) Add(values ...T) {
	list.expand(len(values))
	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}
}

// Get return the value at the specified position and possibly OutOfRangeError
func (list *ArrayList[T]) Get(index int) (oldValue T, err error) {
	if !list.checkIndex(index) {
		return oldValue, new(errors.OutOfRangeError)
	}
	return list.elements[index], nil
}

// Set replaces the value at the specified position in this list with the specified value.
// return the value previously at the specified position and possibly OutOfRangeError
func (list *ArrayList[T]) Set(index int, value T) (oldValue T, err error) {
	if !list.checkIndex(index) {
		return oldValue, new(errors.OutOfRangeError)
	}
	oldValue = list.elements[index]
	list.elements[index] = value
	return oldValue, nil
}

// Remove removes the element at the specified position in this list,
// and shifts any subsequent elements to the left (subtracts one from their indices).
// return the value that was removed from the list and possibly OutOfRangeError
func (list *ArrayList[T]) Remove(index int) (oldValue T, err error) {
	if !list.checkIndex(index) {
		return oldValue, new(errors.OutOfRangeError)
	}
	ret := list.elements[index]
	list.elements[index] = oldValue
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	list.shrink()
	return ret, nil
}

// Contains return true if this list contains all the specified elements.
// More formally, returns true if and only if every element e in param values such that list.equals(o, e)
func (list *ArrayList[T]) Contains(values ...T) bool {
	for _, searchValue := range values {
		found := false
		for index := 0; index < list.size; index++ {
			if list.equals(list.elements[index], searchValue) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// Insert inserts values at specified index position shifting the value at that position (if any) and any subsequent elements to the right.
// return OutOfRangeError if position is negative or bigger than list's size
func (list *ArrayList[T]) Insert(index int, values ...T) error {
	if !list.checkIndex(index) {
		return new(errors.OutOfRangeError)
	}

	l := len(values)
	list.expand(l)
	list.size += l
	copy(list.elements[index+l:], list.elements[index:list.size-l])
	copy(list.elements[index:], values)
	return nil
}

// Sort sorts values (in-place) using.
func (list *ArrayList[T]) Sort(comparator utils.Comparator[T]) {
	if len(list.elements) < 2 {
		return
	}
	utils.Sort(list.elements[:list.size], comparator)
}

// Swap swaps the two elements in index1 and index2
// if any of the two index is negative or bigger than list's size will return OutOfRangeError
func (list *ArrayList[T]) Swap(index1, index2 int) error {
	if !list.checkIndex(index1) || !list.checkIndex(index2) {
		return new(errors.OutOfRangeError)
	}
	list.elements[index1], list.elements[index2] = list.elements[index2], list.elements[index1]
	return nil
}

// Empty return true if this list contains no elements
func (list *ArrayList[T]) Empty() bool {
	return list.size == 0
}

// Size return the number of elements in this list
func (list *ArrayList[T]) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *ArrayList[T]) Clear() {
	list.size = 0
	list.elements = []T{}
}

// Values returns all elements in the list.
func (list *ArrayList[T]) Values() []T {
	newElements := make([]T, list.size, list.size)
	copy(newElements, list.elements[:list.size])
	return newElements
}

func (list *ArrayList[T]) String() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range list.elements[:list.size] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// Expand the array if necessary, i.e. capacity will be reached if we add n elements
func (list *ArrayList[T]) expand(n int) {
	currentCapacity := cap(list.elements)
	if list.size+n >= currentCapacity {
		newCapacity := int(expandFactor * float32(currentCapacity+n))
		list.resize(newCapacity)
	}
}

// Shrink the array if necessary, i.e. when size is shrinkFactor percent of current capacity
func (list *ArrayList[T]) shrink() {
	if shrinkFactor == 0.0 {
		return
	}
	// Shrink when size is at shrinkFactor * capacity
	currentCapacity := cap(list.elements)
	if list.size <= int(float32(currentCapacity)*shrinkFactor) {
		list.resize(list.size)
	}
}

func (list *ArrayList[T]) resize(cap int) {
	newElements := make([]T, cap, cap)
	copy(newElements, list.elements)
	list.elements = newElements
}

func (list *ArrayList[T]) checkIndex(index int) bool {
	return index >= 0 && index < list.size
}
