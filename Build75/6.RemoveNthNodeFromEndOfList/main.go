package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	preSlow, slow, fast := dummy, head, head
	for fast != nil {
		if n <= 0 {
			// n>0 表示還沒到可以啟動的位數
			preSlow = slow
			slow = slow.Next
		}
		n--
		fast = fast.Next
	}
	preSlow.Next = slow.Next
	return dummy.Next
}
