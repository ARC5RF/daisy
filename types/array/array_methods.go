package array

import (
	"log"
	"runtime/debug"

	"github.com/ARC5RF/daisy/perennis"
)

// type Of[T comparable] interface {}

type Of[T comparable] []T

// returns 0 if arr (a pointer to Of[T], itself an array) is nil
//
// returns len(*arr)
func (arr *Of[T]) Length() int {
	if arr == nil {
		return 0
	}
	return len(*arr)
}

// returns nil if arr (a pointer to Of[T], itself an array) is nil
//
// returns nil if index >= arr.Length() or if arr.Length()-index < 0
//
// returns a reference to the value at index or, if index is less than zero, arr.Length()-index
func (arr *Of[T]) At(index int) *T {
	if arr == nil {
		return nil
	}

	l := arr.Length()
	want := index
	if index < 0 {
		want = l + index
	}

	if want >= l || want < 0 {
		return nil
	}

	return &((*arr)[want])
}

// Finds the first occurrence of the element
//
// evaluates a.Equals(b) if present on T before evaluating a == b
//
// returns array.IS_NIL if arr (a pointer to Of[T], itself an array) is nil
//
// returns array.IS_EMPTY if the underlying array has no values
//
// returns the index of the first result if it is found
//
// returns array.NOT_FOUND if nothing is found
func (arr *Of[T]) FirstIndexByValue(value T) int {
	if arr == nil {
		return IS_NIL
	}

	if len(*arr) == 0 {
		return IS_EMPTY
	}

	if equatable, is_equatable := any(value).(perennis.Equatable[T]); is_equatable {
		return arr.FirstIndexByValueFunc(equatable.Equals)
	}

	for i, element := range *arr {
		if element == value {
			return i
		}
	}

	return NOT_FOUND
}

// Finds the last occurrence of the element
//
// evaluates a.Equals(b) if present on T before evaluating a == b
//
// returns array.IS_NIL if arr (a pointer to Of[T], itself an array) is nil
//
// returns array.IS_EMPTY if the underlying array has no values
//
// returns the index of the first result if it is found
//
// returns array.NOT_FOUND if nothing is found
func (array *Of[T]) LastIndexByValue(value T) int {
	if array == nil {
		return IS_NIL
	}

	if len(*array) == 0 {
		return IS_EMPTY
	}

	if equatable, is_equatable := any(value).(perennis.Equatable[T]); is_equatable {
		return array.LastIndexByValueFunc(equatable.Equals)
	}

	for i := array.Length() - 1; i >= 0; i-- {
		if (*array)[i] == value {
			return i
		}
	}

	return NOT_FOUND
}

// Finds all occurrences of the element
//
// evaluates a.Equals(b) if present on T before evaluating a == b
//
// returns nil if arr (a pointer to Of[T], itself an array) is nil
//
// returns an empty *Of[int] if the underlying array has no values
//
// returns all indexes if any are found
//
// returns an empty *Of[int] if nothing is found
func (array *Of[T]) AllIndexesByValue(value T) *Of[int] {
	if array == nil {
		return nil
	}

	if len(*array) == 0 {
		return Make[int](0)
	}

	if equatable, is_equatable := any(value).(perennis.Equatable[T]); is_equatable {
		return array.AllIndexesByValueFunc(equatable.Equals)
	}

	output := Make[int](0)
	for i, element := range *array {
		if element == value {
			output.Push(i)
		}
	}

	return output
}

// Finds the first occurrence of an element matching the condition
//
// returns array.IS_NIL if arr (a pointer to Of[T], itself an array) is nil
//
// returns array.IS_EMPTY if the underlying array has zero elements
//
// returns the index of the first result if it is found
//
// returns array.NOT_FOUND if nothing is found
func (array *Of[T]) FirstIndexByValueFunc(condition func(other T) bool) int {
	if array == nil {
		return IS_NIL
	}

	if len(*array) == 0 {
		return IS_EMPTY
	}

	for i, element := range *array {
		if condition(element) {
			return i
		}
	}

	return NOT_FOUND
}

// Finds the last occurrence of an element matching the condition
//
// returns array.IS_NIL if arr (a pointer to Of[T], itself an array) is nil
//
// returns array.IS_EMPTY if the underlying array has zero elements
//
// returns the index of the first result if it is found
//
// returns array.NOT_FOUND if nothing is found
func (array *Of[T]) LastIndexByValueFunc(condition func(other T) bool) int {
	if array == nil {
		return IS_NIL
	}

	if len(*array) == 0 {
		return IS_EMPTY
	}

	for i := array.Length() - 1; i >= 0; i-- {
		if condition((*array)[i]) {
			return i
		}
	}

	return NOT_FOUND
}

// Finds the first occurrence of an element matching the condition
//
// returns nil if arr (a pointer to Of[T], itself an array) is nil
//
// returns an empty *Of[int] if the underlying array has zero elements
//
// returns all indexes if any are found
//
// returns an empty *Of[int] if nothing is found
func (array *Of[T]) AllIndexesByValueFunc(condition func(other T) bool) *Of[int] {
	if array == nil {
		return nil
	}

	if len(*array) == 0 {
		return Make[int](0)
	}

	output := Make[int](0)
	for i, element := range *array {
		if condition(element) {
			output.Push(i)
		}
	}

	return output
}

// Finds the first occurrence of an element with the same address as value
//
// if array.WARN_ME_ABOUT has bit array.WARN_ABOUT_BAD_NIL_ARGUMENTS set
// it will log a warning when nil is passed to the function
//
// returns the index of the first result if it is found
//
// returns array.IS_NIL if arr (a pointer to Of[T], itself an array) is nil
//
// returns array.IS_EMPTY if the underlying array has no values
//
// returns array.NOT_FOUND if nothing is found
func (arr *Of[T]) FirstIndexOfReferenceWarn(value *T) int {
	if WARN_ME_ABOUT&WARN_ABOUT_BAD_NIL_ARGUMENTS == WARN_ABOUT_BAD_NIL_ARGUMENTS && value == nil {
		log.Println("WARN: first argument of FirstIndexByReference was nil\n\t\tthis will always return a result that is less than zero")
		// log.Println("")
		if WARN_IS_FATAL {
			panic("")
		}
		if WARN_ME_WITH_STACK {
			log.Panicln(debug.Stack())
		}
	}
	return arr.FirstIndexOfReference(value)
}

// Finds the first occurrence of an element with the same address as value
//
// returns array.IS_NIL if arr (a pointer to Of[T], itself an array) is nil
//
// returns array.IS_EMPTY if the underlying array has no values
//
// returns the index of the first result if it is found
//
// returns array.NOT_FOUND if nothing is found
func (arr *Of[T]) FirstIndexOfReference(value *T) int {
	if arr == nil {
		return IS_NIL
	}

	if len(*arr) == 0 {
		return IS_EMPTY
	}

	for i := 0; i < arr.Length(); i++ {
		if arr.At(i) == value {
			return i
		}
	}

	return NOT_FOUND
}

// Finds the first occurrence that satisfies condition(arr.At(i))
//
// returns array.IS_NIL if arr (a pointer to Of[T], itself an array) is nil
//
// returns array.IS_EMPTY if the underlying array has no values
//
// returns the index of the first result if it is found
//
// returns array.NOT_FOUND if nothing is found
func (arr *Of[T]) FirstIndexOfReferenceFunc(condition func(v *T) bool) int {
	if arr == nil {
		return IS_NIL
	}

	if len(*arr) == 0 {
		return IS_EMPTY
	}

	for i := 0; i < arr.Length(); i++ {
		if condition(arr.At(i)) {
			return i
		}
	}

	return NOT_FOUND
}

// returns nil if arr (a pointer to Of[T], itself an array) is nil
//
// returns nil if from >= arr.Length() or if arr.Length()-from < 0
//
// returns nil if to >= arr.Length() or if arr.Length()-to < 0
//
// returns a pointer to a slice of arr[want_from:want_to] where want_ is arr.Length() - want_ if from or to are negative
func (arr *Of[T]) Slice(from int, to int) *Of[T] {
	if arr == nil {
		return nil
	}

	l := arr.Length()
	wfrom := from
	if from < 0 {
		wfrom = l + from
	}

	if wfrom >= l || wfrom < 0 {
		return nil
	}

	wto := to
	if to < 0 {
		wto = l + to
	}

	if wto >= l || wto < 0 {
		return nil
	}

	if wfrom > wto {
		return nil
	}

	o := ((*arr)[wfrom:wto])
	return &o
}

// Adds the specified elements to the end of
//
// noop if the array if arr (a pointer to Of[T], itself an array) is nil
//
// returns itself (even if itself is nil).
func (arr *Of[T]) Push(values ...T) *Of[T] {
	if arr == nil {
		return arr
	}
	*arr = append(*arr, values...)
	return arr
}

// Adds the specified elements to the end of an array
// only if the individual elements are not already in the array.
//
// an element is evaluated to be in the array after the previous item is added
//
// all elements will prioritize an Equatable evaluation before an == evaluation
//
// noop if the array if arr (a pointer to Of[T], itself an array) is nil
//
// returns itself (even if itself is nil).
func (arr *Of[T]) PushMissing(values ...T) *Of[T] {
	if arr == nil {
		return arr
	}

	for _, value := range values {
		if arr.FirstIndexByValue(value) >= 0 {
			continue
		}
		arr.Push(value)
	}

	return arr
}

// Filters the array against a condition
//
// returns a new empty array if arr (a pointer to Of[T], itself an array) is nil
//
// returns a new array containing all matches
//
// arr.At(n) does not have the same address as output.At(n)
func (arr *Of[T]) Filter(condition func(v T, k int, source *Of[T]) bool) (output *Of[T]) {
	if arr == nil {
		return Make[T](0)
	}

	output = Make[T](0)
	for k, v := range *arr {
		if condition(v, k, arr) {
			output.Push(v)
		}
	}

	return output
}

// Tests whether all elements in the array pass the test implemented by the provided function.
//
// returns false if arr (a pointer to Of[T], itself an array) is nil
//
// returns true if condition returns true for all elements
func (arr *Of[T]) Every(condition func(v T, k int, source *Of[T]) bool) bool {
	if arr == nil {
		return false
	}

	for k, v := range *arr {
		if !condition(v, k, arr) {
			return false
		}
	}

	return true
}
