## 20. Valid Parentheses

Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Every close bracket has a corresponding open bracket of the same type.

題目，給定一些字串只會有 '(', ')', '{', '}', '[',']'，請判斷是否是成對。
本題的想法會是使用一個 Stack 先進先出


### Example 1:

Input: s = "()" <br>
Output: true 


### Example 2:

Input: s = "()[]{}"<br>
Output: true


### Example 3:

Input: s = "(]" <br>
Output: false

// 先寫虛擬碼
```golan
func isValid(s string) bool { 
    // 判斷邊緣 case 是不是存在
    // for s 被走訪完成
    // 如果是 ( 或 { 或 [ 則加入堆疊
    // 如果是 ) 或 } 或 [ 先判斷堆疊是否為零
        為零 return false
        不為零
          如果是 ） 的話，從堆疊，最後一個取出來，看看是不是( 如果是繼續比對直到否則return fasle
          如果是 } 的話，從堆疊，最後一個取出來，看看是不是{ 如果是繼續比對直到否則return fasle
           如果是 ] 的話，從堆疊，最後一個取出來，看看是不是[ 如果是繼續比對直到否則return fasle
}
```
