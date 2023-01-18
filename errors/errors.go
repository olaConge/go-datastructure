package errors

type OutOfRangeError struct {
	name string
}

func (e *OutOfRangeError) Error() string {
	return "index out of range"
}
