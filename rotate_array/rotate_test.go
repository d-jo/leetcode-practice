package rotate_test

import (
	"testing"

	rotate "github.com/d-jo/leetcode/rotate_array"
)

func TestRotate1(t *testing.T) {
	inputArr := []int{1, 2, 3, 4, 5, 6, 7}
	inputK := 3

	expected := []int{5, 6, 7, 1, 2, 3, 4}

	rotate.Rotate(inputArr, inputK)

	for i := 0; i < len(inputArr); i++ {
		if inputArr[i] != expected[i] {
			t.Errorf("expected %d, got %d", expected[i], inputArr[i])
		}
	}
}
