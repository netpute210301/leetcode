package main

import "fmt"

func maxArea(height []int) int {
	start, end := 0, len(height)-1
	ma := 0
	for start <= end {
		area := 0
		if height[start] < height[end] {
			area = height[start] * (end - start)
			fmt.Println(area, "a")
			start++
		} else {
			area = height[end] * (end - start)
			fmt.Println(area, "b")
			end--
		}
		ma = max(ma, area)

	}
	return ma
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(input))

}
