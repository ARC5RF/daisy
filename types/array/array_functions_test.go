package array_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/ARC5RF/daisy/types/array"
)

// TODO better test for array.Map
func TestMap(t *testing.T) {
	data := array.Make[string](10)
	for i := range *data {
		*data.At(i) = fmt.Sprint(i)
	}

	ints := array.Map(data, func(v *string, k int, input *array.Of[string]) *int {
		o, err := strconv.Atoi(*v)
		if err != nil {
			return nil
		}
		return &o
	})

	second := ints.At(1)
	if second == nil {
		t.Fail()
		return
	}
	if *second == nil {
		t.Fail()
		return
	}
	if **second != 1 {
		t.Fail()
		return
	}

}
