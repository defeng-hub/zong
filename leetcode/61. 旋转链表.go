package main

func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}
	if head.Next == nil {
		return head
	}

	var slow *ListNode
	fNode := &ListNode{Val: 0}
	slow = fNode
	for i := 0; i < k-1; i++ {
		fNode.Next = &ListNode{Val: 0}
		fNode = fNode.Next
	}
	fNode.Next = head

	fast := head
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fast.Next = head
	kk := slow.Next

	return kk
}
