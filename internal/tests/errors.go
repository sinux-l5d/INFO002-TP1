package tests

import "fmt"

type BoundaryError struct {
	Min uint64
	Max uint64
}

func (e BoundaryError) Error() string {
	return fmt.Sprintf("number must be between %d and %d (inclusive)", e.Min, e.Max)
}
