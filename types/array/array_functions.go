package array

// Runs converter over each index of input and adds the output to the same index of its return value
//
// if array.WARN_ME_ABOUT has bit array.WARN_ABOUT_BAD_EQUATABLE set
// it will log any mismatch between T or O and EquatableValue
func MapWarn[T comparable, O comparable](input *Of[T], converter func(v *T, k int, input *Of[T]) O) *Of[O] {
	warn[T]()
	warn[O]()

	return Map(input, converter)
}

// Runs converter over each index of input and adds the output to the same index of its return value
func Map[T comparable, O comparable](input *Of[T], converter func(v *T, k int, input *Of[T]) O) *Of[O] {
	if input == nil {
		return nil
	}
	output := Make[O](len(*input))
	for i := 0; i < input.Length(); i++ {
		(*output)[i] = converter(input.At(i), i, input)
	}
	return output
}

// Runs visitor over each index of input
//
// breaks out of the loop if visitor returns true
//
// returns input
func Each[T comparable](input *Of[T], visitor func(v *T, k int, input *Of[T]) (exit bool)) *Of[T] {
	for i := 0; i < input.Length(); i++ {
		if visitor(input.At(0), i, input) {
			break
		}
	}
	return input
}
