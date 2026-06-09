package main

import "fmt"

func isPalindrome(x int) bool {
	// เขียนตรงนี้! คืน true ถ้า x อ่านหน้า-หลังเหมือนกัน
	// 	เลขติดลบ → false ได้เลย
	// "ดึงหลักสุดท้าย" ออกมา: x % 10
	// "ตัดหลักสุดท้ายทิ้ง": x / 10
	// เอาหลักที่ดึงมาต่อกันเป็นเลขกลับด้าน: reversed = reversed*10 + digit
	// สุดท้ายเทียบ reversed กับเลขเดิม
	if x < 0 {
		return false
	}
	original := x // เก็บค่าเดิมไว้เทียบทีหลัง
	reversed := 0

	for x > 0 {
		digit := x % 10 // ดึงหลักสุดท้าย
		reversed = reversed*10 + digit
		x = x / 10 // ตัดหลักสุดท้ายทิ้ง
	}
	if original == reversed {
		return true
	}
	return false
}

func main() {
	fmt.Println(isPalindrome(121))  // true
	fmt.Println(isPalindrome(-121)) // false (ติดลบ อ่านกลับเป็น 121-)
	fmt.Println(isPalindrome(10))   // false
}
