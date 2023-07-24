package main

//1. 先寫一個 merge 兩個的方法
//2. 將這個 k 個使用 receive 一次整合一半
//3. 最後合併成一個

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	listLen := len(lists)
	// 考慮邊界狀態
	if listLen == 0 {
		return nil
	}
	if listLen == 1 {
		return lists[0]
	}
	mid := listLen / 2
	left := mergeKLists(lists[:mid])
	right := mergeKLists(lists[mid:])
	return mergeTwoLists(left, right)

}

// 要考慮的狀況是
// 1. 兩個都沒有的邊界狀況
// 2. 有一邊沒有了，把另一邊接到目前答案的後面
// 3. 循環的條件是直到有一邊是空的
// 4. 這種指標的技巧可以使用 dummy head
// 5. 要使用指標的概念
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	if list1 == nil && list2 == nil {
		return dummy.Next
	}
	c := dummy
	for list1 != nil && list2 != nil {
		if list1.Val > list2.Val {
			c.Next = list2
			list2 = list2.Next
		} else {
			c.Next = list1
			list1 = list1.Next
		}
		c = c.Next
	}
	if list1 == nil {
		c.Next = list2
	}
	if list2 == nil {
		c.Next = list1
	}
	return dummy.Next
}
