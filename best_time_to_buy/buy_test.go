package buy_test

import (
	"testing"

	buy "github.com/d-jo/leetcode/best_time_to_buy"
)

func TestMaxProfit1(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := 4
	actual := buy.MaxProfit(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestMaxProfit2(t *testing.T) {
	input := []int{7, 6, 4, 3, 1}
	expected := 0
	actual := buy.MaxProfit(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func TestMaxProfit3(t *testing.T) {
	input := []int{7, 1, 5, 3, 6, 4}
	expected := 7
	actual := buy.MaxProfit(input)

	if expected != actual {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
