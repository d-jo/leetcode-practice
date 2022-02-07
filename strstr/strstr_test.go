package strstr_test

import (
	"testing"

	"github.com/d-jo/leetcode/strstr"
)

func TestStrStr1(t *testing.T) {
	inputHaystack := "hello"
	inputNeedle := "ll"

	expected := 2

	actual := strstr.StrStr(inputHaystack, inputNeedle)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}

func TestStrStr2(t *testing.T) {
	inputHaystack := "worldhello"
	inputNeedle := "ll"

	expected := 7

	actual := strstr.StrStr(inputHaystack, inputNeedle)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}

func TestStrStr3(t *testing.T) {
	inputHaystack := "worldhelo"
	inputNeedle := "ll"

	expected := -1

	actual := strstr.StrStr(inputHaystack, inputNeedle)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}

func TestStrStr4(t *testing.T) {
	inputHaystack := ""
	inputNeedle := "ll"

	expected := -1

	actual := strstr.StrStr(inputHaystack, inputNeedle)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}

func TestStrStr5(t *testing.T) {
	inputHaystack := "worldhelo"
	inputNeedle := ""

	expected := 0

	actual := strstr.StrStr(inputHaystack, inputNeedle)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}

func TestStrStr6(t *testing.T) {
	inputHaystack := "a"
	inputNeedle := "a"

	expected := 0

	actual := strstr.StrStr(inputHaystack, inputNeedle)

	if actual != expected {
		t.Errorf("expected %d gt %d", expected, actual)
	}
}
