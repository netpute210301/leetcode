package main

// 原地交換不需輸出
func rotate(matrix [][]int) {
	// 走對腳斜切
	for row := 0; row < len(matrix); row++ {
		for col := row; col < len(matrix[0]); col++ {
			tmp := matrix[row][col]
			matrix[row][col] = matrix[col][row]
			matrix[col][row] = tmp
		}
	}
	// 在左右交換
	for row := 0; row < len(matrix); row++ {
		for col := 0; col < len(matrix[0])/2; col++ {
			tmp := matrix[row][col]
			matrix[row][col] = matrix[row][len(matrix[0])-col-1]
			matrix[row][len(matrix[0])-col-1] = tmp
		}
	}
}
