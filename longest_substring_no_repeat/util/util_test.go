package util_test

import (
	"testing"

	"github.com/d-jo/leetcode/longest_substring_no_repeat/util"
)

func TestHasRepeat1(t *testing.T) {
	input := "asdf"
	expected := false

	actual := util.HasRepeat(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestHasRepeat2(t *testing.T) {
	input := "asdfasdf"
	expected := true

	actual := util.HasRepeat(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestHasRepeat3(t *testing.T) {
	input := ""
	expected := false

	actual := util.HasRepeat(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestRepeatArr1(t *testing.T) {
	input := "asdf"
	expected := 4

	actual := util.RepeatArr(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestRepeatArr2(t *testing.T) {
	input := "asdfasdf"
	expected := 4

	actual := util.RepeatArr(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestRepeatArr3(t *testing.T) {
	input := ""
	expected := 0

	actual := util.RepeatArr(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestRepeatArr4(t *testing.T) {
	input := "pwwkew"
	expected := 3

	actual := util.RepeatArr(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestRepeatArr5(t *testing.T) {
	input := "bbbbb"
	expected := 1

	actual := util.RepeatArr(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestRepeatArr6(t *testing.T) {
	input := "aab"
	expected := 2

	actual := util.RepeatArr(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}

func TestRepeatArr7(t *testing.T) {
	input := "dvdf"
	expected := 3

	actual := util.RepeatArr(input)

	if actual != expected {
		t.Errorf("expected: %v, actual: %v", expected, actual)
	}
}
