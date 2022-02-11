package dining_philosophers_test

import (
	"testing"
	"time"

	"github.com/d-jo/leetcode/dining_philosophers"
)

func TestMod0(t *testing.T) {
	input := 0
	expected := 0

	actual := (input + 5) % 5

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestMod1(t *testing.T) {
	input := -1
	expected := 4

	actual := (input + 5) % 5

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}
func TestMod2(t *testing.T) {
	input := 5
	expected := 0

	actual := (input + 5) % 5

	if actual != expected {
		t.Errorf("Expected %d, but got %d", expected, actual)
	}
}

func TestGetForkInds0(t *testing.T) {
	input := 0
	expectedLeft := 0
	expectedRight := 4

	actualLeft, actualRight := dining_philosophers.GetForkIndiciesForPhil(input)

	if actualLeft != expectedLeft {
		t.Errorf("Expected left %d, but got %d", expectedLeft, actualLeft)
	}

	if actualRight != expectedRight {
		t.Errorf("Expected right %d, but got %d", expectedRight, actualRight)
	}
}

func TestGetForkInds1(t *testing.T) {
	input := 1
	expectedLeft := 1
	expectedRight := 0

	actualLeft, actualRight := dining_philosophers.GetForkIndiciesForPhil(input)

	if actualLeft != expectedLeft {
		t.Errorf("Expected left %d, but got %d", expectedLeft, actualLeft)
	}

	if actualRight != expectedRight {
		t.Errorf("Expected right %d, but got %d", expectedRight, actualRight)
	}
}

func TestGetForkInds2(t *testing.T) {
	input := 2
	expectedLeft := 2
	expectedRight := 1

	actualLeft, actualRight := dining_philosophers.GetForkIndiciesForPhil(input)

	if actualLeft != expectedLeft {
		t.Errorf("Expected left %d, but got %d", expectedLeft, actualLeft)
	}

	if actualRight != expectedRight {
		t.Errorf("Expected right %d, but got %d", expectedRight, actualRight)
	}
}

func TestGetForkInds3(t *testing.T) {
	input := 3
	expectedLeft := 3
	expectedRight := 2

	actualLeft, actualRight := dining_philosophers.GetForkIndiciesForPhil(input)

	if actualLeft != expectedLeft {
		t.Errorf("Expected left %d, but got %d", expectedLeft, actualLeft)
	}

	if actualRight != expectedRight {
		t.Errorf("Expected right %d, but got %d", expectedRight, actualRight)
	}
}

func TestGetForkInds4(t *testing.T) {
	input := 4
	expectedLeft := 4
	expectedRight := 3

	actualLeft, actualRight := dining_philosophers.GetForkIndiciesForPhil(input)

	if actualLeft != expectedLeft {
		t.Errorf("Expected left %d, but got %d", expectedLeft, actualLeft)
	}

	if actualRight != expectedRight {
		t.Errorf("Expected right %d, but got %d", expectedRight, actualRight)
	}
}

func TestStart(t *testing.T) {

	start := time.Now()
	dining_philosophers.Start()
	elapsed := time.Since(start)

	if elapsed > time.Second*5 {
		t.Errorf("Expected less than 5 second, but got %s", elapsed)
	}
}
