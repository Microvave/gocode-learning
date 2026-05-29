package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// เขียนตรงนี้!
	seen := map[int]int{}

	for i, num := range nums {
		need := target - num

		index, exists := seen[need]
		if exists {
			return []int{index, i} // คืน index คู่ที่เจอ
		}

		seen[num] = i // จดลงสมุด

	}
	return []int{}
}

func main() {
	fmt.Println("=== Array ===")

	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9)) // [0 1]
	fmt.Println(twoSum([]int{3, 2, 4}, 6))      // [1 2]
	fmt.Println(twoSum([]int{3, 3}, 6))

}
