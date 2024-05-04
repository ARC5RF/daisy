package array

// Runs converter over each index of input and adds the output to the same index of its return value
//
// if array.WARN_ME_ABOUT has bit array.WARN_ABOUT_BAD_EQUATABLE set
// it will log any mismatch between T or O and EquatableValue
func MapWarn[T comparable, O comparable](input *Of[T], converter func(v T, k int, input *Of[T]) O) *Of[O] {
	warn[T]()
	warn[O]()

	return Map(input, converter)
}

// Runs converter over each index of input and adds the output to the same index of its return value
func Map[T comparable, O comparable](input *Of[T], converter func(v T, k int, input *Of[T]) O) *Of[O] {
	if input == nil {
		return nil
	}
	output := Make[O](len(*input))
	for i, v := range *input {
		(*output)[i] = converter(v, i, input)
	}
	return output
}
