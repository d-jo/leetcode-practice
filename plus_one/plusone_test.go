package plusone_test

import (
	"testing"

	plusone "github.com/d-jo/leetcode/plus_one"
)

func TestPlusOne1(t *testing.T) {
	input := []int{9}
	expected := []int{1, 0}

	actual := plusone.PlusOne(input)

	t.Logf("%+v", actual)

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %d got %d", expected[i], actual[i])
		}
	}
}

func TestPlusOne2(t *testing.T) {
	input := []int{4, 3, 2, 1}
	expected := []int{4, 3, 2, 2}

	actual := plusone.PlusOne(input)

	t.Logf("%+v", actual)

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %d got %d", expected[i], actual[i])
		}
	}
}

func TestPlusOne3(t *testing.T) {
	input := []int{9, 9, 9, 9, 9, 9, 9, 9}
	expected := []int{1, 0, 0, 0, 0, 0, 0, 0, 0}

	actual := plusone.PlusOne(input)

	t.Logf("%+v", actual)

	for i := 0; i < len(expected); i++ {
		if actual[i] != expected[i] {
			t.Errorf("expected %d got %d", expected[i], actual[i])
		}
	}
}
