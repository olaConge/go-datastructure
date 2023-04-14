package util

type Comparator[T any] func(a, b T) int

// EqualsFn is a function that returns whether 'a' is equal to 'b'.
type EqualsFn[T any] func(a, b T) bool

// Equals wraps the '==' operator for comparable types.
func Equals[T comparable](a, b T) bool {
	return a == b
}
