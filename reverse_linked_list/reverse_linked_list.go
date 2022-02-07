package reverse_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	curr := head
	next := head.Next

	for {
		if next == nil {
			break
		}
		grand := next.Next
		next.Next = curr
		curr = next
		next = grand
	}

	head.Next = nil

	return curr
}
