package climbingstairs_test

import (
	"testing"

	climbingstairs "github.com/d-jo/leetcode/climbing_stairs"
)

func TestClimbStairs1(t *testing.T) {
	input := 2
	expected := 2

	actual := climbingstairs.ClimbStairs(input)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}

func TestClimbStairs2(t *testing.T) {
	input := 3
	expected := 3

	actual := climbingstairs.ClimbStairs(input)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}

func TestClimbStairs3(t *testing.T) {
	input := 4
	expected := 5

	actual := climbingstairs.ClimbStairs(input)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}

func TestClimbStairs4(t *testing.T) {
	input := 5
	expected := 8

	actual := climbingstairs.ClimbStairs(input)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}

	input2 := 10
	expected2 := 89

	actual2 := climbingstairs.ClimbStairs(input2)

	if actual2 != expected2 {
		t.Errorf("expected %d gt %d", expected2, actual2)
	}
}
