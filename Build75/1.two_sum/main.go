package main

import "fmt"

func main() {
	test_case1 := []int{2, 7, 11, 15}
	target1 := 9
	result := []int{0, 1}

	res := twoSum(test_case1, target1)
	fmt.Println(result, res)

}

func twoSum(nums []int, target int) []int {
	posMap := make(map[int]int, len(nums))
	for key, value := range nums {
		if index, ok := posMap[target-value]; ok {
			return []int{index, key}
		}
		posMap[value] = key
	}
	return []int{}
}
