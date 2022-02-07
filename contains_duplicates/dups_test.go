package dups_test

import (
	"testing"

	dups "github.com/d-jo/leetcode/contains_duplicates"
)

func TestContainsDuplicate1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := false

	actual := dups.ContainsDuplicate(input)

	if expected != actual {
		t.Errorf("expected %t, got %t", expected, actual)
	}
}

func TestContainsDuplicate2(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 5}
	expected := true

	actual := dups.ContainsDuplicate(input)

	if expected != actual {
		t.Errorf("expected %t, got %t", expected, actual)
	}
}

func TestContainsDuplicate3(t *testing.T) {
	input := []int{1}
	expected := false

	actual := dups.ContainsDuplicate(input)

	if expected != actual {
		t.Errorf("expected %t, got %t", expected, actual)
	}
}
