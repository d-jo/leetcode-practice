package reverse_string_test

import (
	"testing"

	"github.com/d-jo/leetcode/reverse_string"
)

func TestReverseString1(t *testing.T) {
	input := []byte("asdf")
	expected := []byte("fdsa")

	reverse_string.ReverseString(input)

	for i := 0; i < len(input); i++ {
		if input[i] != expected[i] {
			t.Errorf("expected %s got %s", string(input[i]), string(expected[i]))
		}
	}

}

func TestReverseString2(t *testing.T) {
	input := []byte("asdfg")
	expected := []byte("gfdsa")

	reverse_string.ReverseString(input)

	for i := 0; i < len(input); i++ {
		if input[i] != expected[i] {
			t.Errorf("expected %s got %s", string(input[i]), string(expected[i]))
		}
	}

}
