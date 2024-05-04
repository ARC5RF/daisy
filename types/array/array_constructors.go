package array

// Creates a new array.Of[T] with len of initial_capacity
//
// if array.WARN_ME_ABOUT has bit array.WARN_ABOUT_BAD_EQUATABLE set
// it will log any mismatch between T and EquatableValue
func MakeWarn[T comparable](initial_capacity int) *Of[T] {
	warn[T]()

	return Make[T](initial_capacity)
}

// Creates a new array.Of[T] with len of initial_capacity
func Make[T comparable](initial_capacity int) *Of[T] {
	v := make(Of[T], initial_capacity)

	return &v
}

// Creates a new array.Of[T] with a different pointer than source
//
// if array.WARN_ME_ABOUT has bit array.WARN_ABOUT_BAD_EQUATABLE set
// it will log any mismatch between T and EquatableValue
//
// all values should be understood to be handled the same as from := append([]T{}, source...)
func FromWarn[T comparable](source []T) *Of[T] {
	warn[T]()
	return From(source)
}

// Creates a new array.Of[T] with a different pointer than from
//
// all values should be understood to be handled the same as from := append([]T{}, source...)
func From[T comparable](source []T) *Of[T] {
	if source == nil {
		return Make[T](0)
	}
	output := Make[T](len(source))
	for i, v := range source {
		*output.At(i) = v
	}
	return output
}

// Creates a new array.Of[T] with the same pointer as source
//
// if array.WARN_ME_ABOUT has bit array.WARN_ABOUT_BAD_EQUATABLE set
// it will log any mismatch between T and EquatableValue
func WrapWarn[T comparable](original *[]T) *Of[T] {
	warn[T]()

	return Wrap(original)
}

// Wraps original as *array.Of[T] with the same pointer as original
func Wrap[T comparable](original *[]T) *Of[T] {
	return (*Of[T])(original)
}

// Unwraps *array.Of[T] to *[]T with the same pointer as of
func Unwrap[T comparable](of *Of[T]) *[]T {
	return (*[]T)(of)
}
