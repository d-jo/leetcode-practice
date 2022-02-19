package hamming_distance_test

import (
	"testing"

	"github.com/d-jo/leetcode/hamming_distance"
)

func TestHammingDistance1(t *testing.T) {
	input1, input2 := 1, 4
	expected := 2

	actual := hamming_distance.HammingDistance(input1, input2)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestHammingDistance2(t *testing.T) {
	input1, input2 := 3, 1
	expected := 1

	actual := hamming_distance.HammingDistance(input1, input2)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}

func TestHammingDistance3(t *testing.T) {
	input1, input2 := 1577962638, 1727613287
	expected := 16

	actual := hamming_distance.HammingDistance(input1, input2)

	if actual != expected {
		t.Errorf("Expected %d, got %d", expected, actual)
	}
}
