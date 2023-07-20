## 題目
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
You may assume that each input would have exactly one solution, and you may not use the same element twice.
You can return the answer in any order.

### Example 1:
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].
### Example 2:
Input: nums = [3,2,4], target = 6
Output: [1,2]
### Example 3:
Input: nums = [3,3], target = 6
Output: [0,1]

這題的想法是，用一個 map 存 key = target - 目前的值 , value = 目前的位置
下一輪的時候可以看看有沒有值跟目前的值 是一樣的，如果有，就把目前的位置，跟 map 裡面的位置回傳
