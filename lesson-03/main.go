package main

import "fmt"

func main() {

	// ==========================================
	// 1. IF / ELSE
	// ==========================================
	fmt.Println("=== If / Else ===")

	score := 75

	if score >= 95 {
		fmt.Println("เกรด A")
	} else if score >= 70 {
		fmt.Println("เกรด B")
	} else if score >= 55 {
		fmt.Println("เกรด C")
	} else {
		fmt.Println("เกรด F")
	}

	// ==========================================
	// 2. FOR LOOP — วนซ้ำแบบนับเลข
	// ==========================================
	fmt.Println("\n=== For Loop (นับ 1-5) ===")

	for i := 1; i <= 5; i++ {
		fmt.Println("รอบที่", i)
	}

	// ==========================================
	// 3. FOR LOOP — แบบ while (Go ไม่มี while ใช้ for แทน)
	// ==========================================
	fmt.Println("\n=== For Loop (while style) ===")

	count := 0
	for count < 3 {
		fmt.Println("count =", count)
		count++
	}

	// ==========================================
	// 4. FOR LOOP — วน array
	// ==========================================
	fmt.Println("\n=== For Loop (วน array) ===")

	fruits := []string{"แอปเปิ้ล", "กล้วย", "มะม่วง"}

	for i, fruit := range fruits {
		fmt.Println(i, "->", fruit)
	}

	// ==========================================
	for i := 2; i <= 10; i += 2 {
		fmt.Println(i)
	}

}
