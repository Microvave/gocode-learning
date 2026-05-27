package main

import "fmt"

// ==========================================
// 1. Function พื้นฐาน — รับค่า และ return ค่า
// ==========================================

func greet(name string) string {
	return "สวัสดี " + name + "!"
}

// ==========================================
// 2. Function รับหลาย parameter
// ==========================================

func add(a int, b int) int {
	return a + b
}

// ==========================================
// 3. Function return หลายค่า (Go ทำได้ ต่างจากภาษาอื่น!)
// ==========================================

func minMax(a int, b int) (int, int) {
	if a < b {
		return a, b
	}
	return b, a
}

// ==========================================
// 4. Function ไม่ return อะไร (แค่ทำงาน)
// ==========================================

func printLine() {
	fmt.Println("--------------------")
}

// ==========================================
// MAIN
// ==========================================

func multiply(a int, b int) int {
	// เขียนเองตรงนี้!
	d := a * b
	return d
}

func main() {

	// เรียกใช้ greet
	message := greet("Microwave")
	fmt.Println(message)

	printLine()

	// เรียกใช้ add
	result := add(10, 20)
	fmt.Println("10 + 20 =", result)
	fmt.Println("5 + 7 =", add(5, 7))

	printLine()

	// เรียกใช้ minMax (รับ 2 ค่ากลับมา)
	small, big := minMax(42, 15)
	fmt.Println("ค่าน้อย:", small)
	fmt.Println("ค่ามาก:", big)

	printLine()

	fmt.Println("3 x 4 =", multiply(3, 4))
	fmt.Println("6 x 7 =", multiply(6, 7))

}
