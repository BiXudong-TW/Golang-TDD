package pointer

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_should_get_type_of_pointer(t *testing.T) {
	a := 42
	b := &a
	actual := fmt.Sprint(reflect.TypeOf(b))
	//Please update the expect value to pass the test
	expect := ""
	if expect != actual {
		t.Fail()
	}
}

func Test_should_get_correct_value_with_pointer(t *testing.T) {
	a := 42
	b := &a
	c := *b
	//Please update the expect value to pass the test
	expect := 0
	if expect != c {
		t.Fail()
	}
}

func Test_should_get_nil_when_only_declare_pointer(t *testing.T) {
	var pointer *string
	actual := fmt.Sprint(pointer)
	//Please update the expect value to pass the test
	expect := ""
	if expect != actual {
		t.Fail()
	}
}

func Test_should_get_type_of_new_method(t *testing.T) {
	a := new(int)

	actual := fmt.Sprint(reflect.TypeOf(a))
	//Please update the expect value to pass the test
	expect := ""

	if expect != actual {
		t.Fail()
	}
}

func Test_should_get_type_of_make_method(t *testing.T) {
	a := make(map[int]bool)

	actual := fmt.Sprint(reflect.TypeOf(a))
	//Please update the expect value to pass the test
	expect := ""

	if expect != actual {
		t.Fail()
	}

}
