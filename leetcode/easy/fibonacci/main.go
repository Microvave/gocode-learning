package main

import "fmt"

func fib(n int) int {
	// เขียนตรงนี้! คืนค่า Fibonacci ลำดับที่ n
	// F(0)=0, F(1)=1, F(n)=F(n-1)+F(n-2)
	a := 0
	b := 1

	// เลื่อนไปข้างหน้า n รอบ (ถ้า n==0 ไม่เข้า loop → คืน a=0 พอดี)
	for i := 0; i < n; i++ {
		a, b = b, a+b
	}

	return a
}

func main() {
	fmt.Println(fib(2))  // 1   (0,1,1)
	fmt.Println(fib(3))  // 2   (0,1,1,2)
	fmt.Println(fib(4))  // 3   (0,1,1,2,3)
	fmt.Println(fib(10)) // 55
}
