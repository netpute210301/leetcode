## 48. Rotate Image

You are given an n x n 2D matrix representing an image, 
rotate the image by 90 degrees (clockwise).
You have to rotate the image in-place, 
which means you have to modify the input 2D matrix directly. 
DO NOT allocate another 2D matrix and do the rotation.


### Example 1:

![](mat1.jpg)

Input: matrix = [[1,2,3],[4,5,6],[7,8,9]]

Output: [[7,4,1],[8,5,2],[9,6,3]]

### Example 2:

![](mat2.jpg)

Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]

Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

## 題目：
給定一個2D 陣列，請向右邊 90 度（順時針）旋轉他，請原地旋轉，不要額外新創陣列

## 思路

只要旋轉的題目都是透過對折再對折的解法來解題
這裡旋轉九十度的話可以透過先斜切一半，在左右交換來達成
像是這樣


    1,2,3
    4,5,6
    7,8,9


先斜切一半交換(1,5,9 不變，其餘對角線交換)變成

    1,4,7
    2,5,8
    3,6,9

（如果使用雙迴圈會是倒 L 型的交還） ，來看每一次交換的案例
ex i =1
    
    1,4,7
    2,x,x
    3,x,x

ex i = 2

    x,x,x
    x,5,8
    x,6,x

ex i = 3

    x,x,x
    x,x,x
    x,x,9




然後在左右交換

    7,4,1
    8,5,2
    9,6,3

就完成交換了
