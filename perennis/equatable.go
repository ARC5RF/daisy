package perennis

type Equatable[T any] interface {
	Equals(other T) bool
}
