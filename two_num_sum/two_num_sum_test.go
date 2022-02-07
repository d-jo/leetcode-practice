package twonumsum_test

import (
	"testing"

	twonumsum "github.com/d-jo/leetcode/two_num_sum"
)

func TestBinSearch1(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 5
	expected := 4

	actual := twonumsum.BinSearch(&nums, target)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestBinSearch2(t *testing.T) {
	nums := []int{4, 5, 6, 7, 8, 9, 10}
	target := 5
	expected := 1

	actual := twonumsum.BinSearch(&nums, target)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestBinSearch3(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 11
	expected := -1

	actual := twonumsum.BinSearch(&nums, target)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestBinSearch4(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	target := 1
	expected := 0

	actual := twonumsum.BinSearch(&nums, target)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestBinSearch5(t *testing.T) {
	nums := []int{0, 2, 4}
	target := 3
	expected := -1

	actual := twonumsum.BinSearch(&nums, target)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestBinSearch6(t *testing.T) {
	nums := []int{0, 2, 4}
	target := 2
	expected := 1

	actual := twonumsum.BinSearch(&nums, target)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestTwoSum1(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	expected := []int{0, 1}

	actual := twonumsum.TwoSum(nums, target)

	if actual[0] != expected[0] || actual[1] != expected[1] {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestTwoSum2(t *testing.T) {
	nums := []int{3, 2, 4}
	target := 6
	expected := []int{1, 2}

	actual := twonumsum.TwoSum(nums, target)

	if actual[0] != expected[0] || actual[1] != expected[1] {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestTwoSum3(t *testing.T) {
	nums := []int{3, 2, 4, 1, 66, 7, 8, 9, 10}
	target := 76
	expected := []int{4, 8}

	actual := twonumsum.TwoSum(nums, target)

	if actual[0] != expected[0] || actual[1] != expected[1] {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
