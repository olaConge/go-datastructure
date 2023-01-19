package lists

import (
	"fmt"
	"github.com/olaConge/go-datastructure/errors"
	"github.com/olaConge/go-datastructure/utils"
	"strings"
)

// Assert List implementation
var _ List[any] = new(LinkedList[any])

type LinkedList[T any] struct {
	first  *node[T]
	last   *node[T]
	size   int
	equals func(a, b T) bool
}

type node[T any] struct {
	value T
	prev  *node[T]
	next  *node[T]
}

func NewLinkedList[T any](equals utils.EqualsFn[T], values ...T) *LinkedList[T] {
	list := &LinkedList[T]{
		equals: equals,
	}
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

// Add appends the specified value to the end of this list.
func (list *LinkedList[T]) Add(values ...T) {
	for _, value := range values {
		newNode := &node[T]{
			value: value,
			prev:  list.last,
		}
		if list.size == 0 {
			list.first = newNode
			list.last = newNode
		} else {
			list.last.next = newNode
			list.last = newNode
		}
		list.size++
	}
}

// AddLast appends a value (one or more) at the end of the list (same as Add())
func (list *LinkedList[T]) AddLast(values ...T) {
	list.Add(values...)
}

// AddFirst Inserts the values at the beginning of this list
func (list *LinkedList[T]) AddFirst(values ...T) {
	// in reverse to keep passed order i.e. ["c","d"] -> Prepend(["a","b"]) -> ["a","b","c",d"]
	for v := len(values) - 1; v >= 0; v-- {
		newNode := &node[T]{value: values[v], next: list.first}
		if list.size == 0 {
			list.first = newNode
			list.last = newNode
		} else {
			list.first.prev = newNode
			list.first = newNode
		}
		list.size++
	}
}

// Get return the value at the specified position and possibly OutOfRangeError
func (list *LinkedList[T]) Get(index int) (oldValue T, err error) {
	if !list.checkIndex(index) {
		return oldValue, new(errors.OutOfRangeError)
	}

	// determine traversal direction, last to first or first to last
	if index > (list.size >> 1) {
		curNode := list.last
		for i := list.size - 1; i != index; i, curNode = i-1, curNode.prev {
		}
		return curNode.value, nil
	}
	curNode := list.first
	for i := 0; i != index; i, curNode = i+1, curNode.next {
	}
	return curNode.value, nil
}

// Set replaces the value at the specified position in this list with the specified value.
// return the value previously at the specified position and possibly OutOfRangeError
func (list *LinkedList[T]) Set(index int, value T) (oldValue T, err error) {
	if !list.checkIndex(index) {
		return oldValue, new(errors.OutOfRangeError)
	}

	foundNode := list.foundNodeByIndex(index)
	oldValue = foundNode.value
	foundNode.value = value
	return oldValue, nil
}

// Remove removes the element at the specified position in this list
func (list *LinkedList[T]) Remove(index int) (oldValue T, err error) {
	if !list.checkIndex(index) {
		return oldValue, new(errors.OutOfRangeError)
	}
	if list.size == 1 {
		oldValue = list.first.value
		list.Clear()
		return oldValue, nil
	}

	element := list.foundNodeByIndex(index)
	if element == list.first {
		list.first = element.next
	}
	if element == list.last {
		list.last = element.prev
	}
	if element.prev != nil {
		element.prev.next = element.next
	}
	if element.next != nil {
		element.next.prev = element.prev
	}

	oldValue = element.value
	element = nil
	list.size--
	return oldValue, nil
}

// Contains return true if this list contains all the specified elements.
func (list *LinkedList[T]) Contains(values ...T) bool {
	if len(values) == 0 {
		return true
	}
	if list.size == 0 {
		return false
	}
	for _, value := range values {
		found := false
		for element := list.first; element != nil; element = element.next {
			if list.equals(element.value, value) {
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
func (list *LinkedList[T]) Insert(index int, values ...T) error {
	if !list.checkIndex(index) {
		return new(errors.OutOfRangeError)
	}

	list.size += len(values)

	tempHead := (*node[T])(nil)
	tempTail := (*node[T])(nil)
	for i, value := range values {
		newNode := &node[T]{
			value: value,
			prev:  tempTail,
		}
		if i == 0 {
			tempHead = newNode
			tempTail = newNode
		} else {
			tempTail.next = newNode
			tempTail = newNode
		}
	}

	foundNode := list.foundNodeByIndex(index)
	if foundNode == list.first {
		tempTail.next = list.first
		list.first.prev = tempTail
		list.first = tempHead
	} else {
		tempHead.prev = foundNode.prev
		foundNode.prev.next = tempHead
		foundNode.prev = tempTail
		tempTail.next = foundNode
	}
	return nil
}

// Sort sorts values (in-place) using.
func (list *LinkedList[T]) Sort(comparator utils.Comparator[T]) {
	if list.size < 2 {
		return
	}

	values := list.Values()
	utils.Sort(values, comparator)

	list.Clear()

	list.Add(values...)
}

// Swap swaps the two elements' value in index1 and index2, don't change the address of the node.
// if any of the two index is negative or bigger than list's size will return OutOfRangeError
func (list *LinkedList[T]) Swap(index1, index2 int) error {
	if !list.checkIndex(index1) || !list.checkIndex(index2) {
		return new(errors.OutOfRangeError)
	}
	if index1 == index2 {
		return nil
	}
	var node1, node2 *node[T]
	for i, curNode := 0, list.first; node1 == nil || node2 == nil; i, curNode = i+1, curNode.next {
		if i == index1 {
			node1 = curNode
		}
		if i == index2 {
			node2 = curNode
		}
	}
	node1.value, node2.value = node2.value, node1.value
	return nil
}

// Empty return true if this list contains no elements
func (list *LinkedList[T]) Empty() bool {
	return list.size == 0
}

// Size return the number of elements in this list
func (list *LinkedList[T]) Size() int {
	return list.size
}

// Clear removes all elements from the list.
func (list *LinkedList[T]) Clear() {
	list.size = 0
	list.first = nil
	list.last = nil
}

// Values returns all elements in the list.
func (list *LinkedList[T]) Values() []T {
	values := make([]T, list.size, list.size)
	for i, node := 0, list.first; node != nil; i, node = i+1, node.next {
		values[i] = node.value
	}
	return values
}

func (list *LinkedList[T]) String() string {
	str := "LinkedList\n"
	values := []string{}
	for node := list.first; node != nil; node = node.next {
		values = append(values, fmt.Sprintf("%v", node.value))
	}
	str += strings.Join(values, ", ")
	return str
}

func (list *LinkedList[T]) checkIndex(index int) bool {
	return index >= 0 && index < list.size
}

func (list *LinkedList[T]) foundNodeByIndex(index int) *node[T] {
	var foundNode *node[T]
	// determine traversal direction, last to first or first to last
	if index > (list.size >> 1) {
		foundNode = list.last
		for i := list.size - 1; i != index; {
			i, foundNode = i-1, foundNode.prev
		}
	} else {
		foundNode = list.first
		for i := 0; i != index; {
			i, foundNode = i+1, foundNode.next
		}
	}
	return foundNode
}
