package remove_nth_from_end_test

import (
	"testing"

	rem_nth "github.com/d-jo/leetcode/remove_nth_from_end"
)

func MakeLinkedList(nums []int) *rem_nth.ListNode {
	head := &rem_nth.ListNode{}
	cur := head

	for _, n := range nums {
		cur.Next = &rem_nth.ListNode{Val: n}
		cur = cur.Next
	}

	return head.Next
}

func TestMakeLinkedList(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{1, 2, 3, 4, 5}

	actual := MakeLinkedList(input)

	for _, n := range expected {
		t.Logf("%+v", actual)
		if actual.Val != n {
			t.Errorf("expected %d gt %d", n, actual.Val)
		}
		actual = actual.Next
	}
}

func TestRmNthFromEnd1(t *testing.T) {
	input := MakeLinkedList([]int{1, 2, 3, 4, 5})
	input_n := 2

	expected := []int{1, 2, 3, 5}

	actual := rem_nth.RemoveNthFromEnd(input, input_n)

	for _, n := range expected {
		if actual.Val != n {
			t.Errorf("expected %d gt %d", n, actual.Val)
		}
		actual = actual.Next
	}
}

func TestRmNthFromEnd2(t *testing.T) {
	input := MakeLinkedList([]int{1, 2})
	input_n := 1

	expected := []int{1}

	t.Logf("%+v", input)

	actual := rem_nth.RemoveNthFromEnd(input, input_n)

	for _, n := range expected {
		if actual.Val != n {
			t.Errorf("expected %d gt %d", n, actual.Val)
		}
		actual = actual.Next
	}

	//t.Fail()

}

func TestRmNthFromEnd3(t *testing.T) {
	input := MakeLinkedList([]int{1})
	input_n := 1

	var expected *rem_nth.ListNode

	actual := rem_nth.RemoveNthFromEnd(input, input_n)

	if actual != expected {
		t.Fail()
	}

}

func TestRmNthFromEnd4(t *testing.T) {
	input := MakeLinkedList([]int{1, 2})
	input_n := 2

	expected := []int{2}

	t.Logf("%+v", input)

	actual := rem_nth.RemoveNthFromEnd(input, input_n)

	for _, n := range expected {
		if actual.Val != n {
			t.Errorf("expected %d gt %d", n, actual.Val)
		}
		actual = actual.Next
	}

	//t.Fail()

}
