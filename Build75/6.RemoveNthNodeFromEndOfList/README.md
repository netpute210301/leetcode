### 19. Remove Nth Node From End of List

Given the head of a linked list, remove the nth node from the end of the list and return its head.

給定一個 連結串列，移除從尾巴數過來的地 N 個 node, 然後把鏈結串列的頭部回傳

## Example 1:

![](remove_ex1.jpg)

Input: head = [1,2,3,4,5], n = 2 <br>
Output: [1,2,3,5]

### Example 2:
Input: head = [1], n = 1  <br>
Output: []

### Example 3:
Input: head = [1,2], n = 1 <br>
Output: [1]

想法是快慢指針，我們會像同一個方向前進，然後先扣除n, 唯有當 n <=0 時才會開始
使用最開始的頭，可以減少重複，或者是不好的特殊情況
