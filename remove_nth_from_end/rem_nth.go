package remove_nth_from_end

type ListNode struct {
	Val  int
	Next *ListNode
}

type NodeQueue struct {
	Nodes   []*ListNode
	MaxSize int
}

func NewQueue(size int) *NodeQueue {
	return &NodeQueue{Nodes: make([]*ListNode, 0), MaxSize: size}
}

func (q *NodeQueue) Enqueue(node *ListNode) {
	q.Nodes = append(q.Nodes, node)
	if len(q.Nodes) > q.MaxSize {
		q.Nodes = q.Nodes[1:]
	}
}

func (q *NodeQueue) Dequeue() *ListNode {
	if len(q.Nodes) == 0 {
		return nil
	}
	return q.Nodes[0]
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	queue := NewQueue(n + 1)

	sz := 0
	newHead := head

	for head != nil {
		queue.Enqueue(head)
		head = head.Next
		sz += 1
	}

	target := queue.Dequeue()

	if target == nil || target.Next == nil {
		return nil
	}

	if sz == n {
		return newHead.Next
	}

	target.Next = target.Next.Next

	return newHead

}
