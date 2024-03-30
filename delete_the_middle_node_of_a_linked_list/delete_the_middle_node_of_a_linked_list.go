package deletethemiddlenodeofalinkedlist

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {

	var size int
	for n := head; n != nil; n = n.Next {
		size++
	}

	if size < 2 {
		return nil
	}

	n := head
	for i := 0; i < size/2-1; i++ {
		n = n.Next
	}

	n.Next = n.Next.Next

	return head
}
