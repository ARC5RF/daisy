package array_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ARC5RF/daisy/types/array"
)

type bad_equatable_a struct{}

func (e bad_equatable_a) Equals(other bad_equatable_a) {}

type bad_equatable_b struct{}

func (e bad_equatable_b) Equals(other *bad_equatable_b) {}

type bad_equatable_c struct{}

func (e *bad_equatable_c) Equals(other bad_equatable_c) {}

type bad_equatable_d struct{}

func (e *bad_equatable_d) Equals(other *bad_equatable_d) {}

type bad_equatable_e struct{}

func (e bad_equatable_e) Equals(other bad_equatable_a) {}

type bad_equatable_f struct{}

func (e *bad_equatable_f) Equals(other *bad_equatable_b) {}

func TestMakeWarn(t *testing.T) {
	old := array.WARN_ME_ABOUT
	array.WARN_ME_ABOUT = array.WARN_ABOUT_BAD_EQUATABLE

	a_p := "*array_test.bad_equatable_a"
	a_v := "array_test.bad_equatable_a"
	a := fmt.Sprintf(array.EV_FOOTPRINT+"\n", a_p, a_p, a_v, a_v)
	if out := cap(func() { array.MakeWarn[*bad_equatable_a](0) }); !strings.HasSuffix(out, a) {
		t.Log(out)
		t.Fail()
		return
	}

	b_p := "*array_test.bad_equatable_b"
	b_v := "array_test.bad_equatable_b"
	b := fmt.Sprintf(array.EV_FOOTPRINT+"\n", b_v, b_v, b_v, b_p)
	if out := cap(func() { array.MakeWarn[bad_equatable_b](0) }); !strings.HasSuffix(out, b) {
		t.Log(out)
		t.Fail()
		return
	}

	c_p := "*array_test.bad_equatable_c"
	c_v := "array_test.bad_equatable_c"
	c := fmt.Sprintf(array.EV_FOOTPRINT+"\n", c_v, c_v, c_p, c_v)
	if out := cap(func() { array.MakeWarn[bad_equatable_c](0) }); !strings.HasSuffix(out, c) {
		t.Log(out)
		t.Fail()
		return
	}

	d_p := "*array_test.bad_equatable_d"
	d_v := "array_test.bad_equatable_d"
	d := fmt.Sprintf(array.EV_FOOTPRINT+"\n", d_v, d_v, d_p, d_p)
	if out := cap(func() { array.MakeWarn[bad_equatable_d](0) }); !strings.HasSuffix(out, d) {
		t.Log(out)
		t.Fail()
		return
	}

	e_p := "*array_test.bad_equatable_e"
	e_v := "array_test.bad_equatable_e"
	e := fmt.Sprintf(array.EV_FOOTPRINT+"\n", e_p, e_p, e_v, a_v)
	if out := cap(func() { array.MakeWarn[*bad_equatable_e](0) }); !strings.HasSuffix(out, e) {
		t.Log(out)
		t.Fail()
		return
	}

	f_p := "*array_test.bad_equatable_f"
	f_v := "array_test.bad_equatable_f"
	f := fmt.Sprintf(array.EV_FOOTPRINT+"\n", f_v, f_v, f_p, b_p)
	if out := cap(func() { array.MakeWarn[bad_equatable_f](0) }); !strings.HasSuffix(out, f) {
		t.Log(out)
		t.Fail()
		return
	}

	array.WARN_ME_ABOUT = old
}

func TestFromWarn(t *testing.T) {
	old := array.WARN_ME_ABOUT
	array.WARN_ME_ABOUT = array.WARN_ABOUT_BAD_EQUATABLE

	a_p := "*array_test.bad_equatable_a"
	a_v := "array_test.bad_equatable_a"
	a := fmt.Sprintf(array.EV_FOOTPRINT+"\n", a_p, a_p, a_v, a_v)
	if out := cap(func() { array.FromWarn([]*bad_equatable_a{}) }); !strings.HasSuffix(out, a) {
		t.Log(out)
		t.Fail()
		return
	}

	// f_p := "array_test.bad_equatable_f"
	// f_v := "array_test.bad_equatable_f"
	b_p := "*array_test.bad_equatable_b"
	b_v := "array_test.bad_equatable_b"
	b := fmt.Sprintf(array.EV_FOOTPRINT+"\n", b_v, b_v, b_v, b_p)
	if out := cap(func() { array.FromWarn([]bad_equatable_b{}) }); !strings.HasSuffix(out, b) {
		t.Log(out)
		t.Fail()
		return
	}

	c_p := "*array_test.bad_equatable_c"
	c_v := "array_test.bad_equatable_c"
	c := fmt.Sprintf(array.EV_FOOTPRINT+"\n", c_v, c_v, c_p, c_v)
	if out := cap(func() { array.FromWarn([]bad_equatable_c{}) }); !strings.HasSuffix(out, c) {
		t.Log(out)
		t.Fail()
		return
	}

	d_p := "*array_test.bad_equatable_d"
	d_v := "array_test.bad_equatable_d"
	d := fmt.Sprintf(array.EV_FOOTPRINT+"\n", d_v, d_v, d_p, d_p)
	if out := cap(func() { array.FromWarn([]bad_equatable_d{}) }); !strings.HasSuffix(out, d) {
		t.Log(out)
		t.Fail()
		return
	}

	e_p := "*array_test.bad_equatable_e"
	e_v := "array_test.bad_equatable_e"
	e := fmt.Sprintf(array.EV_FOOTPRINT+"\n", e_p, e_p, e_v, a_v)
	if out := cap(func() { array.FromWarn([]*bad_equatable_e{}) }); !strings.HasSuffix(out, e) {
		t.Log(out)
		t.Fail()
		return
	}

	f_p := "*array_test.bad_equatable_f"
	f_v := "array_test.bad_equatable_f"
	f := fmt.Sprintf(array.EV_FOOTPRINT+"\n", f_v, f_v, f_p, b_p)
	if out := cap(func() { array.FromWarn([]bad_equatable_f{}) }); !strings.HasSuffix(out, f) {
		t.Log(out)
		t.Fail()
		return
	}

	array.WARN_ME_ABOUT = old
}

type from_foo struct{}

func TestFrom(t *testing.T) {
	basic := []string{"a"}
	advanced := array.From(basic)

	if fmt.Sprintf("%p", &basic) == fmt.Sprintf("%p", advanced) {
		t.Log("both have the same address")
		t.Fail()
		return
	}

	if len(*advanced) != len(basic) {
		t.Log("both have different lengths")
		t.Fail()
		return
	}

	first := advanced.At(0)
	if &basic[0] == first {
		t.Log("both have the same address for their 0th element")
		t.Fail()
		return
	}

	basic_b := []*from_foo{{}}
	advanced_b := array.From(basic_b)

	if fmt.Sprintf("%p", &basic_b) == fmt.Sprintf("%p", advanced_b) {
		t.Log("both have the same address for their 0th element")
		t.Fail()
		return
	}

	first_b := advanced_b.At(0)
	if first_b == nil {
		t.Log("the created array does not have a 0th element")
		t.Fail()
		return
	}

	if basic_b[0] != *first_b {
		t.Log("both have different values for their 0th element")
		t.Fail()
		return
	}
}

func TestWrapWarn(t *testing.T) {
	old := array.WARN_ME_ABOUT
	array.WARN_ME_ABOUT = array.WARN_ABOUT_BAD_EQUATABLE

	a_p := "*array_test.bad_equatable_a"
	a_v := "array_test.bad_equatable_a"
	a := fmt.Sprintf(array.EV_FOOTPRINT+"\n", a_p, a_p, a_v, a_v)
	if out := cap(func() { array.WrapWarn(&[]*bad_equatable_a{}) }); !strings.HasSuffix(out, a) {
		t.Log(out)
		t.Fail()
		return
	}

	b_p := "*array_test.bad_equatable_b"
	b_v := "array_test.bad_equatable_b"
	b := fmt.Sprintf(array.EV_FOOTPRINT+"\n", b_v, b_v, b_v, b_p)
	if out := cap(func() { array.WrapWarn(&[]bad_equatable_b{}) }); !strings.HasSuffix(out, b) {
		t.Log(out)
		t.Fail()
		return
	}

	c_p := "*array_test.bad_equatable_c"
	c_v := "array_test.bad_equatable_c"
	c := fmt.Sprintf(array.EV_FOOTPRINT+"\n", c_v, c_v, c_p, c_v)
	if out := cap(func() { array.WrapWarn(&[]bad_equatable_c{}) }); !strings.HasSuffix(out, c) {
		t.Log(out)
		t.Fail()
		return
	}

	d_p := "*array_test.bad_equatable_d"
	d_v := "array_test.bad_equatable_d"
	d := fmt.Sprintf(array.EV_FOOTPRINT+"\n", d_v, d_v, d_p, d_p)
	if out := cap(func() { array.WrapWarn(&[]bad_equatable_d{}) }); !strings.HasSuffix(out, d) {
		t.Log(out)
		t.Fail()
		return
	}

	e_p := "*array_test.bad_equatable_e"
	e_v := "array_test.bad_equatable_e"
	e := fmt.Sprintf(array.EV_FOOTPRINT+"\n", e_p, e_p, e_v, a_v)
	if out := cap(func() { array.WrapWarn(&[]*bad_equatable_e{}) }); !strings.HasSuffix(out, e) {
		t.Log(out)
		t.Fail()
		return
	}

	f_p := "*array_test.bad_equatable_f"
	f_v := "array_test.bad_equatable_f"
	f := fmt.Sprintf(array.EV_FOOTPRINT+"\n", f_v, f_v, f_p, b_p)
	if out := cap(func() { array.WrapWarn(&[]bad_equatable_f{}) }); !strings.HasSuffix(out, f) {
		t.Log(out)
		t.Fail()
		return
	}

	array.WARN_ME_ABOUT = old
}

func TestWrap(t *testing.T) {
	basic := []string{"a"}
	advanced := array.Wrap(&basic)

	if fmt.Sprintf("%p", &basic) != fmt.Sprintf("%p", advanced) {
		t.Log("both have different addresses")
		t.Fail()
		return
	}

	if len(*advanced) != len(basic) {
		t.Log("both have different lengths")
		t.Fail()
		return
	}

	first := advanced.At(0)
	if &basic[0] != first {
		t.Log("both have different addresses for their 0th element")
		t.Fail()
		return
	}

	basic_b := []*from_foo{{}}
	advanced_b := array.Wrap(&basic_b)

	if fmt.Sprintf("%p", &basic_b) != fmt.Sprintf("%p", advanced_b) {
		t.Log("both have different addresses for their 0th element")
		t.Fail()
		return
	}

	first_b := advanced_b.At(0)
	if first_b == nil {
		t.Log("the created array does not have a 0th element")
		t.Fail()
		return
	}

	if basic_b[0] != *first_b {
		t.Log("both have different values for their 0th element")
		t.Fail()
		return
	}
}
