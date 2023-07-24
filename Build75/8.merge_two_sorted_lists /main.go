package main

//要考慮的狀況是
//1. 兩個都沒有的邊界狀況
//2. 有一邊沒有了，把另一邊接到目前答案的後面
//3. 循環的條件是直到有一邊是空的
//4. 這種指標的技巧可以使用 dummy head

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	if list1 == nil && list2 == nil {
		return dummy.Next
	}
	cursor := dummy // 單純指標，移動用，有增減還是用 dummyHead 增減，不然移動掉了就只會輸出最後一個
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			cursor.Next = list1
			list1 = list1.Next
		} else {
			cursor.Next = list2
			list2 = list2.Next
		}
		cursor = cursor.Next
	}

	if list1 == nil {
		cursor.Next = list2
	}
	if list2 == nil {
		cursor.Next = list1
	}
	return dummy.Next
}
