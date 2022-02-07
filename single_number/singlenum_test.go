package singlenum_test

import (
	"testing"

	singlenum "github.com/d-jo/leetcode/single_number"
)

func TestSingleNum1(t *testing.T) {
	input := []int{2, 2, 1}
	expected := 1

	actual := singlenum.SingleNumber(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
