package array_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/ARC5RF/daisy/types/array"
)

func TestLength(t *testing.T) {
	a := array.Make[string](10)

	if a.Length() != 10 {
		t.Fail()
		return
	}

	b := array.Wrap(&[]string{"a", "b"})
	if b.Length() != 2 {
		t.Fail()
		return
	}
}

func TestAt(t *testing.T) {
	data := array.Make[string](10)
	for i := range *data {
		*data.At(i) = string(rune('a' + i))
	}
	a := data.At(0)
	if a == nil {
		t.Fail()
		return
	}
	if *a != "a" {
		t.Log(a)
		t.Fail()
		return
	}

	j := data.At(-1)
	if j == nil {
		t.Fail()
		return
	}
	if *j != "j" {
		t.Log(a)
		t.Fail()
		return
	}

	if data.At(20) != nil {
		t.Fail()
		return
	}

	if data.At(-20) != nil {
		t.Fail()
		return
	}
}

type index_by_value_b struct{ q string }

type index_by_value_c struct{ q string }

func (i *index_by_value_c) Equals(other index_by_value_c) bool {
	return i.q == other.q
}

func TestFirstIndexByValue(t *testing.T) {
	a := array.Wrap(&[]string{"a", "b", "a"})

	a_a_idx := a.FirstIndexByValue("a")
	if a_a_idx != 0 {
		t.Log(a_a_idx)
		t.Fail()
		return
	}

	a_b_idx := a.FirstIndexByValue("b")
	if a_b_idx != 1 {
		t.Log(a_b_idx)
		t.Fail()
		return
	}

	b := array.Wrap(&[]index_by_value_b{{"a"}, {"b"}, {"a"}})

	b_a_idx := b.FirstIndexByValue(index_by_value_b{"a"})
	if b_a_idx != 0 {
		t.Log(b_a_idx)
		t.Fail()
		return
	}

	b_b_idx := b.FirstIndexByValue(index_by_value_b{"b"})
	if b_b_idx != 1 {
		t.Log(b_b_idx)
		t.Fail()
		return
	}

	c := array.Wrap(&[]*index_by_value_c{{"a"}, {"b"}, {"a"}})

	c_a_idx := c.FirstIndexByValue(&index_by_value_c{"a"})
	if c_a_idx >= 0 {
		t.Log(c_a_idx)
		t.Fail()
		return
	}

	c_b_idx := c.FirstIndexByValue(*c.At(1))
	if c_b_idx != 1 {
		t.Log(c_b_idx)
		t.Fail()
		return
	}

}

func TestLastIndexByValue(t *testing.T) {
	a := array.Wrap(&[]string{"a", "b", "a", "b"})

	a_a_idx := a.LastIndexByValue("a")
	if a_a_idx != 2 {
		t.Log(a_a_idx)
		t.Fail()
		return
	}

	a_b_idx := a.LastIndexByValue("b")
	if a_b_idx != 3 {
		t.Log(a_b_idx)
		t.Fail()
		return
	}

	b := array.Wrap(&[]index_by_value_b{{"a"}, {"b"}, {"a"}, {"b"}})

	b_a_idx := b.LastIndexByValue(index_by_value_b{"a"})
	if b_a_idx != 2 {
		t.Log(b_a_idx)
		t.Fail()
		return
	}

	b_b_idx := b.LastIndexByValue(index_by_value_b{"b"})
	if b_b_idx != 3 {
		t.Log(b_b_idx)
		t.Fail()
		return
	}

	c := array.Wrap(&[]*index_by_value_c{{"a"}, {"b"}, {"a"}, {"b"}})

	c_a_idx := c.LastIndexByValue(&index_by_value_c{"a"})
	if c_a_idx >= 0 {
		t.Log(c_a_idx)
		t.Fail()
		return
	}

	c_b_idx := c.LastIndexByValue(*c.At(3))
	if c_b_idx != 3 {
		t.Log(c_b_idx)
		t.Fail()
		return
	}

}

func TestFirstIndexByValueFunc(t *testing.T) {
	a := array.Wrap(&[]string{"a", "b", "a", "b"})

	a_a_idx := a.FirstIndexByValueFunc(func(v string) bool { return v == "a" })
	if a_a_idx != 0 {
		t.Log(a_a_idx)
		t.Fail()
		return
	}

	a_b_idx := a.FirstIndexByValueFunc(func(v string) bool { return v == "b" })
	if a_b_idx != 1 {
		t.Log(a_b_idx)
		t.Fail()
		return
	}

	b := array.Wrap(&[]index_by_value_b{{"a"}, {"b"}, {"a"}, {"b"}})

	b_a_idx := b.FirstIndexByValueFunc(func(v index_by_value_b) bool { return v.q == "a" })
	if b_a_idx != 0 {
		t.Log(b_a_idx)
		t.Fail()
		return
	}

	b_b_idx := b.FirstIndexByValueFunc(func(v index_by_value_b) bool { return v.q == "b" })
	if b_b_idx != 1 {
		t.Log(b_b_idx)
		t.Fail()
		return
	}

	c := array.Wrap(&[]*index_by_value_c{{"a"}, {"b"}, {"a"}})

	c_a_idx := c.FirstIndexByValueFunc(func(v *index_by_value_c) bool { return v.q == "a" })
	if c_a_idx != 0 {
		t.Log(c_a_idx)
		t.Fail()
		return
	}

	c_b_idx := c.FirstIndexByValueFunc(func(v *index_by_value_c) bool { return v.q == "b" })
	if c_b_idx != 1 {
		t.Log(c_b_idx)
		t.Fail()
		return
	}

}

func TestLastIndexByValueFunc(t *testing.T) {
	a := array.Wrap(&[]string{"a", "b", "a", "b"})

	a_a_idx := a.LastIndexByValueFunc(func(v string) bool { return v == "a" })
	if a_a_idx != 2 {
		t.Log(a_a_idx)
		t.Fail()
		return
	}

	a_b_idx := a.LastIndexByValueFunc(func(v string) bool { return v == "b" })
	if a_b_idx != 3 {
		t.Log(a_b_idx)
		t.Fail()
		return
	}

	b := array.Wrap(&[]index_by_value_b{{"a"}, {"b"}, {"a"}, {"b"}})

	b_a_idx := b.LastIndexByValueFunc(func(v index_by_value_b) bool { return v.q == "a" })
	if b_a_idx != 2 {
		t.Log(b_a_idx)
		t.Fail()
		return
	}

	b_b_idx := b.LastIndexByValue(index_by_value_b{"b"})
	if b_b_idx != 3 {
		t.Log(b_b_idx)
		t.Fail()
		return
	}

	c := array.Wrap(&[]*index_by_value_c{{"a"}, {"b"}, {"a"}, {"b"}})
	c.LastIndexByValueFunc(func(v *index_by_value_c) bool { return v.q == "a" })

	c_a_idx := c.LastIndexByValueFunc(func(v *index_by_value_c) bool { return v.q == "a" })
	if c_a_idx != 2 {
		t.Log(c_a_idx)
		t.Fail()
		return
	}

	c_b_idx := c.LastIndexByValueFunc(func(v *index_by_value_c) bool { return v.q == "b" })
	if c_b_idx != 3 {
		t.Log(c_b_idx)
		t.Fail()
		return
	}
}

type index_of_reference_a struct{ q string }

func TestFirstIndexOfReferenceWarn(t *testing.T) {
	old := array.WARN_ME_ABOUT
	array.WARN_ME_ABOUT = array.WARN_ABOUT_BAD_EQUATABLE | array.WARN_ABOUT_BAD_NIL_ARGUMENTS

	var a *array.Of[index_of_reference_a]

	if out := cap(func() { a = array.WrapWarn(&[]index_of_reference_a{{"a"}, {"b"}, {"a"}}) }); out != "" {
		t.Log(out)
		t.Fail()
		return
	}

	nil_out := "WARN: first argument of FirstIndexByReference was nil\n\t\tthis will always return a result that is less than zero\n"
	if out := cap(func() { a.FirstIndexOfReferenceWarn(nil) }); !strings.HasSuffix(out, nil_out) {
		t.Log(out)
		t.Fail()
		return
	}

	p := &((*a)[0])

	var p_idx int
	if out := cap(func() { p_idx = a.FirstIndexOfReferenceWarn(p) }); out != "" {
		t.Log(out)
		t.Fail()
		return
	}
	if p_idx != 0 {
		t.Log(p_idx)
		t.Fail()
		return
	}

	array.WARN_ME_ABOUT = old
}

func TestSlice(t *testing.T) {
	var a *array.Of[string]
	b := a.Slice(0, 0)
	if b != nil {
		t.Log(b)
		t.Fail()
		return
	}
	c := array.Wrap(&[]string{"a"})
	d := c.Slice(10, 100)
	if d != nil {
		t.Log(d)
		t.Fail()
		return
	}

	e := array.Wrap(&[]string{"a", "b"})
	f := e.Slice(1, 100)
	if f != nil {
		t.Log(f)
		t.Fail()
		return
	}

	g := []string{"a", "b", "c", "d", "e"}
	h := array.Wrap(&g)

	i := h.Slice(0, 2)
	if i == nil {
		t.Log(i)
		t.Fail()
		return
	}

	h_p := fmt.Sprintf("%p", h.At(0))
	i_p := fmt.Sprintf("%p", i.At(0))
	if h_p != i_p {
		t.Log(h_p, i_p)
		t.Fail()
		return
	}

	j := h.Slice(-2, -1)

	j_p := fmt.Sprintf("%p", j.At(0))
	i_p_ := fmt.Sprintf("%p", h.At(3))
	if j_p != i_p_ {
		t.Log(j, h)
		t.Log(j_p, i_p_)
		t.Fail()
		return
	}
	fmt.Println(j)

	k := h.Slice(0, -1)

	k_p := fmt.Sprintf("%p", k.At(0))
	// i_p_ := fmt.Sprintf("%p", h.At(3))
	if k_p != i_p {
		t.Log(k, h)
		t.Log(k_p, i_p_)
		t.Fail()
		return
	}
	fmt.Println(j)
}

func TestPush(t *testing.T) {
	var a *array.Of[string] = nil
	a.Push("a")
	if a != nil {
		t.Log(a)
		t.Fail()
		return
	}
	b := array.Wrap(&[]string{})
	b.Push("a")
	if b.Length() != 1 {
		t.Log(b)
		t.Fail()
		return
	}
}

type missing_a struct{ q string }

func (me missing_a) Equals(other missing_a) bool {
	return strings.EqualFold(me.q, other.q)
}

type missing_b struct{ q string }

func (me *missing_b) Equals(other *missing_b) bool {
	return strings.EqualFold(me.q, other.q)
}

func TestPushMissing(t *testing.T) {
	old := array.WARN_ME_ABOUT

	//this is still here to catch bad warnings
	array.WARN_ME_ABOUT = array.WARN_ABOUT_BAD_EQUATABLE
	array.WARN_IS_FATAL = true

	var a *array.Of[string] = nil
	a.PushMissing("a")
	if a != nil {
		t.Log(a)
		t.Fail()
		return
	}

	b := array.Wrap(&[]string{})
	b.PushMissing("a")
	if b.Length() != 1 {
		t.Log(b)
		t.Fail()
		return
	}

	b.PushMissing("a")
	if b.Length() != 1 {
		t.Log(b)
		t.Fail()
		return
	}

	b.PushMissing("b")
	if b.Length() != 2 {
		t.Log(b)
		t.Fail()
		return
	}

	c := array.WrapWarn(&[]missing_a{{"a"}})
	c.PushMissing(missing_a{"A"})
	if c.Length() != 1 {
		t.Log(c)
		t.Fail()
		return
	}

	d := array.WrapWarn(&[]*missing_b{{"a"}})
	d.PushMissing(&missing_b{"A"})
	if d.Length() != 1 {
		t.Log(d)
		t.Fail()
		return
	}

	e := array.WrapWarn(&[]*missing_b{{"a"}})
	e.PushMissing(&missing_b{"B"})
	if e.Length() == 1 {
		t.Log(e)
		t.Fail()
		return
	}

	array.WARN_ME_ABOUT = old
}

func TestFilter(t *testing.T) {
	a := array.Wrap(&[]string{"a", "b", "c", "d"})

	b := a.Filter(func(v string, k int, source *array.Of[string]) bool {
		return k < 2
	})

	a_p_a := fmt.Sprintf("%p", a.At(0))
	b_p_a := fmt.Sprintf("%p", b.At(0))
	fmt.Println(b, b.Length(), a_p_a, b_p_a)
}

func TestEvery(t *testing.T) {
	var a *array.Of[string] = nil

	if a.Every(func(v string, k int, source *array.Of[string]) bool {
		return true
	}) {
		t.Log(a)
		t.Fail()
		return
	}

	b := array.Wrap(&[]string{"a", "b", "c", "d"})

	if !b.Every(func(v string, k int, source *array.Of[string]) bool {
		return len(v) == 1
	}) {
		t.Log(a)
		t.Fail()
		return
	}
}
