## 5. Longest Palindromic Substring

Given a string s, return the longest palindromic substring in s.
給定一組字串，會回覆在裡面最長的回文字串

### Example 1:

Input: s = "babad"

Output: "bab"

Explanation: "aba" is also a valid answer.

### Example 2:

Input: s = "cbbd"

Output: "bb"


// 使用中心檢查法，也就是說由中心向外檢查，但由中心向外檢查有兩種情況

1 ex: b a b a b 由中心 b 向外檢查的話 b 的位置是 2 先檢查左邊跟右邊 0 號位置，然後做邊右邊一號位置
    step1 => left = center - 1 , right = center + 1
    step2 => left = center - 2 , right = center + 2
    .... 類推
2. ex cbbd 這樣的回文也算是一中回文 中心是 0 號位置時



// 核心算法，拿到目前的左邊右邊，當成中心，往外擴散比較是不是正確的

func helper(s string, left, right int64) int64 {
    for (left>=0 && right < len(s) && s[left] == s[right]){
        left--
        right++
    }
    // 返回回文最長字串
    return right - left -1
}