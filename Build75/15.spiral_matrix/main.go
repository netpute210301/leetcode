package main

func spiralOrder(matrix [][]int) []int {
	var result []int
	rowStart, rowEnd, colStart, colEnd := 0, len(matrix)-1, 0, len(matrix[0])-1

	for rowStart <= rowEnd && colStart <= colEnd {
		// 第一次先往右邊走訪，走訪完把第一行標記為走完了
		for i := colStart; i <= colEnd; i++ {
			result = append(result, matrix[rowStart][i])
		}
		rowStart++
		//向下走訪，走完把最後一列標記為做完了
		for i := rowStart; i <= rowEnd; i++ {
			result = append(result, matrix[i][colEnd])
		}
		colEnd--
		// 向左走訪，走訪完把這一列標記為走完了
		if rowStart <= rowEnd {
			for i := colEnd; i >= colStart; i-- {
				result = append(result, matrix[rowEnd][i])
			}
			rowEnd--
		}
		// 向上走訪，走訪完把這一列標記為走完了
		if colStart <= colEnd {
			for i := rowEnd; i >= rowStart; i-- {
				result = append(result, matrix[i][colStart])
			}
			colStart++
		}
	}
	return result
}
