package main

func search(nums []int, target int) int {
	// 不可能找到的狀態，因為傳入沒東西
	if len(nums) == 0 {
		return -1
	}
	low, hight := 0, len(nums)-1 // 最前面的指標，跟最後的指標
	for low <= hight {
		mid := low + (hight-low)/2
		if nums[mid] == target {
			// 找到的剛好是中間數
			return mid
		}

		if nums[mid] > nums[low] {
			// 篩選中間以前的區域
			if nums[low] <= target && nums[mid] > target {
				hight = mid - 1
			} else {
				low = mid + 1
			}
		} else if nums[mid] < nums[hight] {
			// 篩選中間以後的區域
			if nums[hight] >= target && nums[mid] < target {
				low = mid + 1
			} else {
				hight = mid - 1

			}
		} else {
			// 有重複的數字
			if nums[mid] == nums[low] {
				low++
			}
			if nums[mid] == nums[hight] {
				hight--
			}
		}
	}
	return -1
}
