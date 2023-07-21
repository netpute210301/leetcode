## 54. Spiral Matrix

Given an m x n matrix, return all elements of the matrix in spiral order.

### Example 1
![e1](spiral1.jpg)

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]

Output: [1,2,3,6,9,8,7,4,5]

### Example 2
![e1](spiral.jpg)

Input: matrix = [[1,2,3,4],[5,6,7,8],[9,10,11,12]]

Output: [1,2,3,4,8,12,11,10,9,5,6,7]

題目的要求是給定一個 m * n 的陣列，返回螺旋狀的排列。這題很簡單

只要按照左下右上的順序，將陣列中的數字放入最大的陣列當中，即可
時間複雜度為 O(m * n)