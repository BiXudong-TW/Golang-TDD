package array

import (
	"testing"
)

func Test_should_init_array_correctly1(t *testing.T) {
	actual := [3]int{1, 2}
	//Please fill all the elements of the expect array
	expect := [3]int{}

	if expect != actual {
		t.Fail()
	}
}

func Test_should_init_array_correctly2(t *testing.T) {
	actual := [5]int{2: 4, 4: 2}
	//Please fill all the elements of the expect array
	expect := [5]int{}

	if expect != actual {
		t.Fail()
	}
}

func Test_should_get_correct_array_length(t *testing.T) {
	actual := [...]int{1, 2, 3}
	//Please update the expect value correctly
	expect := 0

	if expect != len(actual) {
		t.Fail()
	}
}

func Test_should_get_length_and_capcity_of_array1(t *testing.T) {
	array := [8]int{2, 3, 4, 5, 6, 7}
	expectLen := 0
	expectCapacity := 0
	if expectLen != len(array) {
		t.Fail()
	}
	if expectCapacity != cap(array) {
		t.Fail()
	}
}

func Test_should_get_length_and_capcity_of_array2(t *testing.T) {
	array := [...]int{2, 3, 4, 5, 6, 7}
	expectLen := 0
	expectCapacity := 0
	if expectLen != len(array) {
		t.Fail()
	}
	if expectCapacity != cap(array) {
		t.Fail()
	}
}
