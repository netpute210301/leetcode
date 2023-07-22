package main

func longestPalindrome(s string) string {
	// 沒有的話，直接返回
	if len(s) < 2 {
		return s
	}
	start, end := 0, 0
	// 從第一個字符開始遍歷 s，將當前字符的索引記為 i。
	for i := 0; i < len(s); i++ {
		// 呼叫 helper 函數兩次，分別計算以 i 為中心的奇數長度的回文子串長度 len1 和以 i 和 i+1 為中心的偶數長度的回文子串長度 len2。
		len1 := helper(s, i, i)   // case 奇數回文 [aba]vv
		len2 := helper(s, i, i+1) // case 偶數回文狀態 z[bb]x
		// 取 len1 和 len2 中的最大值 maxLen，表示以 i 為中心的最長回文子串長度。
		maxLen := max(len1, len2)

		// 如果 maxLen 大於 right - left，則更新 left 和 right 的值，使其表示最長回文子串的起始位置和結束位置。
		if maxLen > end-start {
			// 最大長度是有算左邊跟右邊的，我們現在只要單邊
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	// 返回邊界< Go 語言的切片中，左閉右開< 要處理
	return s[start : end+1]

}

func helper(strs string, left, right int) int {
	for left >= 0 && right < len(strs) && strs[left] == strs[right] {
		left--
		right++
	}
	//返回最大長度
	return right - left - 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
