package array

import (
	"testing"
)

func Test_Claim(t *testing.T) {
	actual := [3]int{1, 2}
	//Please fill all the array elements of the actual
	expect := [3]int{1, 2, 0}
	if actual != expect {
		t.Fail()
	}
}
