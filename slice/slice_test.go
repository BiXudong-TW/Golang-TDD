package slice

import (
	"testing"
)

func Test_should_get_correct_array_slice1(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	actual := array[1:3]
	//please fill the element is the expect array
	expect := []int{}

	if !Equal(expect, actual) {
		t.Fail()
	}
}

func Test_should_get_correct_array_slice2(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	actual := array[1:]
	//please fill the element is the expect array
	expect := []int{}

	if !Equal(expect, actual) {
		t.Fail()
	}
}

func Test_should_get_correct_array_slice3(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	actual := array[:4]
	//please fill the element is the expect array
	expect := []int{}

	if !Equal(expect, actual) {
		t.Fail()
	}
}

func Test_should_get_correct_array_slice4(t *testing.T) {
	var actual []int

	//Please update the following operator to pass the test
	if actual == nil {
		t.Fail()
	}
}

func Test_should_get_correct_array_slice5(t *testing.T) {
	actual := make([]int, 5)
	//please fill the element is the expect array
	expect := []int{}

	if !Equal(expect, actual) {
		t.Fail()
	}
}

func Test_should_get_full_array_slice(t *testing.T) {
	array := [5]int{1, 2, 3, 4, 5}
	actual := array[:]
	//please fill the element is the expect array
	expect := []int{}

	if !Equal(expect, actual) {
		t.Fail()
	}
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
