package main

import "fmt"

func containsDuplicate(nums []int) bool {
	seen := map[int]bool{}

	for _, num := range nums {
		if seen[num] { // เคยเห็น num นี้แล้วไหม?
			return true // เคยแล้ว = ซ้ำ!
		}
		seen[num] = true // ยังไม่เคย → จดลงสมุด
	}
	return false
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))                   // true
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4}))                   // false
	fmt.Println(containsDuplicate([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2})) // true
}
