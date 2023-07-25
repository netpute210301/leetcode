## 33. Search in Rotated Sorted Array


There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot 
index k (1 <= k < nums.length) such that the resulting array is 
[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed).

For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

給定一個有排序過的陣列，但是這個陣列可能在不知道第幾個序有被旋轉過一次，請在 O(log n) 的時間複雜度內，
尋找到 target 是否在陣列中，如果是，要知道在第幾個



### Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0 <br>
Output: 4


### Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3  <br>
Output: -1  <br>


### Example 3:

Input: nums = [1], target = 0 <br>
Output: -1


1. 考慮邊界，以及 corner case 
2. 雖然順序有點改變，但是還是可以使用二元搜尋法，因為他想要把演算法的複雜度限制在 O(log n)
3. 虛擬碼，想法步驟
4. 中間點很重要 mid := start + (end-start) / 2 每次都要記得是從過濾過的中間
5. 重點中的重點，判斷左邊有序還是右邊有序，因為這個在未知的點被旋轉過了，所以還是要先判斷哪一邊是有序的，先過濾有序那邊
   在考慮無序的那半邊，進了有序的那半邊之後，用有序的首跟尾來看 target 是否在中間，不是的話就移動下標，看保留左半邊還是右半邊

```golang
// 虛擬碼
func search(nums []int, target int) int {
    // 不可能找到的狀態，因為傳入沒東西
	// 最前面的指標，跟最後的指標
      for low <= hight{
         mid := low + (hight-low) / 2
         // 找到的剛好是中間數
         // 篩選中間以前的區域
         // 篩選中間以後的區域
         // 有重複的數字
         return -1
      }
}
```
