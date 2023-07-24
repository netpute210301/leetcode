## 11. Container With Most Water

You are given an integer array height of length n. There are n vertical lines drawn such that the two endpoints of the ith line are (i, 0) and (i, height[i]).

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return the maximum amount of water a container can store.

Notice that you may not slant the container.

題目就是在兩條線的中間，找到最大容量的面積

### Example 1:

![](question_11.jpg)
Input: height = [1,8,6,2,5,4,8,3,7]

Output: 49

Explanation: 
    The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7].<br>
    In this case, the max area of water (blue section) the container can contain is 49.

### Example 2:

Input: height = [1,1]

Output: 1


想法
使用左右兩邊的指針，將長的一邊往內推，短的一邊留著，讓這短邊跟底相乘，得到一個面積紀錄起來跟原本的面積相比

想法的虛擬代碼

left = 0 
right  = len(input) -1
maxArea = 0

for left <= right{
    if input[left] >= input[right] {
        area = input[right] * (right - left)
        maxArea = max(maxArea, area)
        left++
    }else {
        area = input[left] * (right - left)
        maxArea = max(maxArea, area)
        right--
    }
}