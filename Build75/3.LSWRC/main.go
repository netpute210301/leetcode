package main

import "fmt"

func main() {
	var input = "abcabcbb"
	var output = 3
	res := lengthOfLongestSubstring(input)
	fmt.Println(res == output)
}
func lengthOfLongestSubstring(s string) int {
	var tmpRecord = make(map[byte]int, len(s))
	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		if index, ok := tmpRecord[s[right]]; ok {
			// 如果裡面有紀錄，更新左指針
			left = max(index, right+1)
		}
		maxLen = max(maxLen, right-left+1)
		tmpRecord[s[right]] = right
	}
	return maxLen
}

func max(n1, n2 int) int {
	if n1 > n2 {
		return n1
	}
	return n2
}
