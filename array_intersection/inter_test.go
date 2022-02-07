package inter_test

import (
	"testing"

	inter "github.com/d-jo/leetcode/array_intersection"
)

func TestIntersection1(t *testing.T) {
	inputA := []int{1, 2, 2, 1}
	inputB := []int{2, 2}
	expected := []int{2, 2}

	actual := inter.Intersect(inputA, inputB)

	if len(expected) != len(actual) {
		t.Errorf("expected %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("expected %d, got %d", v, actual[i])
		}
	}

}

func TestIntersection2(t *testing.T) {
	inputA := []int{4, 9, 5}
	inputB := []int{9, 4, 9, 8, 4}
	expected := []int{9, 4}

	actual := inter.Intersect(inputA, inputB)

	if len(expected) != len(actual) {
		t.Errorf("expected %d, got %d", len(expected), len(actual))
	}

	for i, v := range expected {
		if v != actual[i] {
			t.Errorf("expected %d, got %d", v, actual[i])
		}
	}

}
