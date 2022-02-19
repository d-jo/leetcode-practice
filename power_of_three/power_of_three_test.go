package power_of_three_test

import (
	"math"
	"testing"

	"github.com/d-jo/leetcode/power_of_three"
)

func TestIsInt(t *testing.T) {
	input := 33.0

	if input != math.Trunc(input) {
		t.Fail()
	}
}

func TestIsInt1(t *testing.T) {
	input := 33.3333

	if input == math.Trunc(input) {
		t.Fail()
	}
}

func TestAffirm1(t *testing.T) {
	input := 3
	expected := true

	actual := power_of_three.IsPowerOfThree(input)

	if actual != expected {
		t.Fail()
	}
}

func TestAffirm2(t *testing.T) {
	input := 9
	expected := true

	actual := power_of_three.IsPowerOfThree(input)

	if actual != expected {
		t.Fail()
	}
}
