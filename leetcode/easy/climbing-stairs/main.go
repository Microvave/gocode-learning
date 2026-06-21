package main

import "fmt"

func climbStairs(n int) int {
	// เขียนตรงนี้!
	a, b := 0, 1
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}
	return b
}

func main() {
	fmt.Println(climbStairs(2)) // 2
	fmt.Println(climbStairs(3)) // 3
	fmt.Println(climbStairs(5)) // 8
}
