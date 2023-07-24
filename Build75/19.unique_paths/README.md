## 62. Unique Paths

There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). The robot can only move either down or right at any point in time.

Given the two integers m and n, return the number of possible unique paths that the robot can take to reach the bottom-right corner.

The test cases are generated so that the answer will be less than or equal to 2 * 109.

題目是給定一個 m*n 的網格，有一個機器人，想要起點走到終點，總共有幾總走法
限制是這個機器人只能走下，跟右邊，無法走回頭路

### example 1 
Input: m = 3, n = 7 <br>
Output: 28

![](robot_maze.png)

有兩種想法
1. 動態規劃，動態規劃常使用在解決這種，發現路徑的問題
2. 公式為 dp[i][j] = dp[i][j-1]+ dp[i-i][j]
3. 也就是這格等於來時的路的所有格子相加

舉例說明
先定義終點為 1
因為題目限制只能走左邊跟下面，那我們先算最下面的邊界的部分人選
最下面那一排，下面都是零，只能走左邊所以到這邊可以移動的步伐都為1
最右邊，因為左邊都不能走，只能走下面，所以整排也都為一

```
? -> ? -> ? -> ? -> ? -> ? -> 1
|    |    |    |    |    |    |
? -> ? -> ? -> ? -> ? -> ? -> 1
|    |    |    |    |    |    |
1 -> 1 -> 1 -> 1 -> 1 -> 1 -> 1
```




開始填倒數第二排....類推
也從後面開始填回來，因為這樣可以把前面累加上去
```
? -> ? -> ? -> ? -> ? -> ? -> 1
|    |    |    |    |    |    |
? -> ? -> ? -> ? -> ? -> 2 -> 1
|    |    |    |    |    |    |
1 -> 1 -> 1 -> 1 -> 1 -> 1 -> 1
```

```
? -> ? -> ? -> ? -> ? -> ? -> 1
|    |    |    |    |    |    |
? -> ? -> ? -> ? -> 3 -> 2 -> 1
|    |    |    |    |    |    |
1 -> 1 -> 1 -> 1 -> 1 -> 1 -> 1
```
.
.
.
```
? -> ? -> ? -> ? -> ? -> ? -> 1
|    |    |    |    |    |    |
7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1
|    |    |    |    |    |    |
1 -> 1 -> 1 -> 1 -> 1 -> 1 -> 1
```

往下一排
```
28 -> 21 -> 15 -> 10 -> 6 -> 3 -> 1
|    |    |    |    |    |    |
7 -> 6 -> 5 -> 4 -> 3 -> 2 -> 1
|    |    |    |    |    |    |
1 -> 1 -> 1 -> 1 -> 1 -> 1 -> 1
```

滿了就回28 表示有28種方法可以走到目的地

func unique(n,m int) int{
    // 產生一個空白且夠寬的陣列
    var roadMap = make([][]int,m)
    // 先將最後面那一排填滿
    for i:=0; i<m; i++{
        roadMap[i] = make([]int,n)
        roadMap[i][n] = 1
    }
    // 再將最下面那一排填滿
    for i:=0; i<n; i++{
        roadMap[m][i] = 1
    }
    // 實踐DP dp[i][j] = dp[i][j-1]+ dp[i-i][j]
    for i:= 

    return roadMap[0][0]

}