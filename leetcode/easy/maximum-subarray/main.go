package main

import "fmt"

func maxSubArray(nums []int) int {
	current := 0
	best := nums[0]

	for _, num := range nums {
		if num > current+num {
			current = num
		} else {
			current = current + num
		}
		if current > best {
			best = current
		}
	}
	return best
}

func main() {
	fmt.Println(maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})) // 6  ([4,-1,2,1])
	fmt.Println(maxSubArray([]int{1}))                             // 1
	fmt.Println(maxSubArray([]int{5, 4, -1, 7, 8}))                // 23
}
