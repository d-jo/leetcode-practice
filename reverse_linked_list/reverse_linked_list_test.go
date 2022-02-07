package reverse_linked_list_test

import (
	"testing"

	"github.com/d-jo/leetcode/reverse_linked_list"
)

func MakeLinkedList(nums []int) *reverse_linked_list.ListNode {
	head := &reverse_linked_list.ListNode{}
	cur := head

	for _, n := range nums {
		cur.Next = &reverse_linked_list.ListNode{Val: n}
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

func TestReverseList1(t *testing.T) {
	input := MakeLinkedList([]int{1, 2, 3, 4, 5})
	expected := []int{5, 4, 3, 2, 1}

	actual := reverse_linked_list.ReverseList(input)

	for _, n := range expected {
		t.Logf("%+v", actual)
		if actual.Val != n {
			t.Errorf("expected %d gt %d", n, actual.Val)
		}
		actual = actual.Next
	}

	t.Fail()
}

func TestReverseList2(t *testing.T) {
	input := MakeLinkedList([]int{1})
	expected := []int{1}

	actual := reverse_linked_list.ReverseList(input)

	for _, n := range expected {
		t.Logf("%+v", actual)
		if actual.Val != n {
			t.Errorf("expected %d gt %d", n, actual.Val)
		}
		actual = actual.Next
	}
}

func TestReverseList3(t *testing.T) {
	input := MakeLinkedList([]int{})
	var expected *reverse_linked_list.ListNode

	actual := reverse_linked_list.ReverseList(input)

	if actual != expected {
		t.Errorf("expected %+v gt %+v", expected, actual)
	}
}
