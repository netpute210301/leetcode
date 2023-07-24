## Merge Two Sorted Lists

You are given the heads of two sorted linked lists list1 and list2.
Merge the two lists in a one sorted list. The list should be made by splicing together the nodes of the first two lists.
Return the head of the merged linked list.

### Example 1:
Input: list1 = [1,2,4], list2 = [1,3,4] <br>
Output: [1,1,2,3,4,4]

### Example 2:

Input: list1 = [], list2 = [] <br>
Output: []

### Example 3:

Input: list1 = [], list2 = [0] <br>
Output: [0]

要考慮的狀況是 
1. 兩個都沒有的邊界狀況
2. 有一邊沒有了，把另一邊接到目前答案的後面
3. 循環的條件是直到有一邊是空的
4. 這種指標的技巧可以使用 dummy head
5. cursor := dummy // 單純指標，移動用，有增減還是用 dummyHead 增減，不然移動掉了就只會輸出最後一個