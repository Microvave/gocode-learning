package main

import "fmt"

func main() {

	// ==========================================
	// 1. VARIABLES — ตัวแปร
	// ==========================================

	// วิธีที่ 1: ประกาศแบบบอก type ชัดเจน
	var name string = "Microwave"
	var age int = 25
	var height float64 = 175.5
	var isStudent bool = true

	fmt.Println("=== Variables ===")
	fmt.Println("ชื่อ:", name)
	fmt.Println("อายุ:", age)
	fmt.Println("ส่วนสูง:", height)
	fmt.Println("เป็นนักเรียน:", isStudent)

	// วิธีที่ 2: ให้ Go เดา type เอง (:=)
	city := "Lampang"
	score := 100
	fmt.Println("เมือง:", city)
	fmt.Println("คะแนน:", score)

	// ==========================================
	// 2. TYPES — ประเภทข้อมูล
	// ==========================================
	fmt.Println("\n=== Types ===")

	var a int = 10       // จำนวนเต็ม
	var b float64 = 3.14 // ทศนิยม
	var c string = "Go"  // ข้อความ
	var d bool = false   // true/false

	fmt.Println("int:", a)
	fmt.Println("float64:", b)
	fmt.Println("string:", c)
	fmt.Println("bool:", d)

	// ==========================================
	// 3. CONSTANTS — ค่าคงที่ (เปลี่ยนไม่ได้)
	// ==========================================
	fmt.Println("\n=== Constants ===")

	const PI float64 = 3.14159
	const AppName string = "MyApp"
	const MaxScore int = 100

	fmt.Println("PI:", PI)
	fmt.Println("App:", AppName)
	fmt.Println("Max Score:", MaxScore)

	// ==========================================
	// test
	// ==========================================
	var myGoal string = "เข้า Google และ Agoda"
	fmt.Println("เป้าหมาย:", myGoal)

}
