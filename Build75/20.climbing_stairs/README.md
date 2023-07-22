## 70. Climbing Stairs


You are climbing a staircase. It takes n steps to reach the top.
Each time you can either climb 1 or 2 steps. In how many distinct ways can you climb to the top?



### Example 1:

Input: n = 2
Output: 2
Explanation: There are two ways to climb to the top.
1. 1 step + 1 step
2. 2 steps


### Example 2:

Input: n = 3
Output: 3
Explanation: There are three ways to climb to the top.
1. 1 step + 1 step + 1 step
2. 1 step + 2 steps
3. 2 steps + 1 step

爬樓梯，看有多少種可能的狀況
例如： 有五層樓梯
1. 1+1+1+1+1
2. 1+1+1+2
3. 1+2+1+1
4. 2+1+1+1
5. 1+1+2+1
6. 2+2+1
7. 1+2+2
8. 2+1+2

在這個前提下，感覺可以用分治而治
我有一個方法，是傳入的數字，如果這個數字小於等於2 表示是我一次可以走的步數，就不再繼續拆分
如果不是，就呼叫自己再繼續拆分
例如 5
 (5) 進入之後可以呼叫 (-1)走一步 (-2)走兩步 的個別狀態
 (4) (3) 走一步是4 走兩步是三 別忘了判斷是否小於等於二，否則就繼續呼叫自己
 (3) (2) (2) (1) 兩個數字都呼叫自己之後，小魚二的可以加起來了，還沒的繼續拆
 (2)(1)(2)(2)(1) ，把這裡數字加起來就可以了，2 的情形可以考慮說他可能在這個順序，走一步，或走兩步
  



