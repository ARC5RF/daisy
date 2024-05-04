// TODO doc comment for whole array package
package array

import (
	"fmt"
	"log"
	"reflect"
	"runtime/debug"
	"strings"
)

const (
	WARN_ABOUT_NOTHING = 1 << iota
	WARN_ABOUT_BAD_NIL_ARGUMENTS
	WARN_ABOUT_BAD_EQUATABLE
)

var WARN_ME_ABOUT = WARN_ABOUT_NOTHING
var WARN_ME_WITH_STACK = false
var WARN_IS_FATAL = false

const (
	IS_NIL    = -1
	IS_EMPTY  = -2
	NOT_FOUND = -3
	NIL_REF   = -4
)

const EV_FOOTPRINT = "Equals has the wrong footprint\n\texpected func (%s) Equals(%s)\n\tgot func (%s) Equals(%s)\n"

type footprint struct {
	name string
	recv string
	args []string
}

func footprints(rt reflect.Type) []footprint {
	output := []footprint{}
	for i := 0; i < rt.NumMethod(); i++ {
		m := rt.Method(i)
		if m.Name == "Equals" {
			o := footprint{}
			o.name = m.Name
			raw := m.Type.String()
			end := strings.Index(raw, ")")
			lr := strings.Split(raw[5:end], ", ")
			o.recv = lr[0]
			if len(lr) > 1 {
				o.args = lr[1:]
			}
			output = append(output, o)
		}
	}
	return output
}

func warn_for_bad_equatable_v(rto reflect.Type) (bool, string) {
	fmt.Println("non pointer type passed to check")
	// fmt.Println(rto.String())

	had := map[string]uint8{}
	for _, got := range footprints(rto) {
		if _, has := had[got.name]; !has {
			had[got.name] = 0
		}
		if len(got.args) != 1 || got.args[0] != got.recv {
			return true, fmt.Sprintf(EV_FOOTPRINT, got.recv, got.recv, got.recv, strings.Join(got.args, ", "))
		}
	}

	rto_p := reflect.PointerTo(rto)
	for _, got := range footprints(rto_p) {
		if _, has := had[got.name]; has {
			continue
		}
		return true, fmt.Sprintf(EV_FOOTPRINT, got.recv[1:], got.recv[1:], got.recv, strings.Join(got.args, ", "))
	}

	return false, ""
}

func warn_for_bad_equatable_p(rto reflect.Type) (bool, string) {
	// fmt.Println("pointer type passed to check")

	// var v *reflect.Method

	rto_e := rto.Elem()

	for _, got := range footprints(rto_e) {
		// if len(got.args) != 1 || got.args[0] == got.recv {
		return true, fmt.Sprintf(EV_FOOTPRINT, "*"+got.recv, "*"+got.recv, got.recv, strings.Join(got.args, ", "))
		// }
	}

	// rto_p := reflect.PointerTo(rto)
	for _, got := range footprints(rto) {
		if len(got.args) != 1 || got.args[0] != got.recv {
			return true, fmt.Sprintf(EV_FOOTPRINT, got.recv, got.recv, got.recv, strings.Join(got.args, ", "))
		}
	}

	return false, ""
}

// todo check footprint for valid interface
func warn_for_bad_equatable[T comparable]() (bool, string) {
	var temp T

	rto := reflect.TypeOf(temp)

	if rto.Kind() == reflect.Pointer {
		return warn_for_bad_equatable_p(rto)
	}

	return warn_for_bad_equatable_v(rto)
}

func warn[T comparable]() {
	if (WARN_ME_ABOUT & WARN_ABOUT_BAD_EQUATABLE) == WARN_ABOUT_BAD_EQUATABLE {
		if bad, why := warn_for_bad_equatable[T](); bad {
			if WARN_IS_FATAL {
				panic(why)
			}
			log.Println(why)
			if WARN_ME_WITH_STACK {
				log.Println(string(debug.Stack()))
			}
		}
	}
}
