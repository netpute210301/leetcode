## 23. Merge k Sorted Lists

You are given an array of k linked-lists lists, each linked-list is sorted in ascending order.

Merge all the linked-lists into one sorted linked-list and return it.

給定 k 個 linked list 每一個 由小到大排序，題目要 merge 所有的 linked list


想法，可以 merger 兩個就可以 merge n 個，寫一個 merge 兩個的來 merge n 個


### Example 1:

Input: lists = [[1,4,5],[1,3,4],[2,6]] <br>
Output: [1,1,2,3,4,4,5,6] <br>
Explanation: The linked-lists are: <br>
```
[
    1->4->5,
    1->3->4,
    2->6
]
```

merging them into one sorted list: <br>
1->1->2->3->4->4->5->6

### Example 2:
Input: lists = []<br>
Output: []<br>

### Example 3:

Input: lists = [[]]<br>
Output: []<br>

1. 先寫一個 merge 兩個的方法
2. 將這個 k 個使用 receive 一次整合一半
3. 最後合併成一個