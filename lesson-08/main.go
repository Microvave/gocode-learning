package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func fullInfo(name string, birthYear int) string {
	currentYear := time.Now().Year()
	age := currentYear - birthYear
	return fmt.Sprintf("%s เกิดปี %d อายุ %d ปี", name, birthYear, age)
}

func main() {

	// ==========================================
	// 1. fmt — พิมพ์ข้อความ (ใช้มาตลอด!)
	// ==========================================
	fmt.Println("=== fmt ===")

	name := "Microwave"
	age := 25

	fmt.Println("ชื่อ:", name)
	fmt.Printf("ชื่อ: %s อายุ: %d ปี\n", name, age) // จัดรูปแบบ
	text := fmt.Sprintf("Hello, %s!", name)         // เก็บใส่ตัวแปร
	fmt.Println(text)

	// ==========================================
	// 2. strings — จัดการข้อความ
	// ==========================================
	fmt.Println("\n=== strings ===")

	sentence := "Hello, Go Programming!"

	fmt.Println(strings.ToUpper(sentence))                    // ตัวใหญ่
	fmt.Println(strings.ToLower(sentence))                    // ตัวเล็ก
	fmt.Println(strings.Contains(sentence, "Go"))             // มีคำนี้ไหม
	fmt.Println(strings.Replace(sentence, "Go", "Golang", 1)) // แทนคำ
	fmt.Println(strings.Split("a,b,c,d", ","))                // ตัดด้วย ,
	fmt.Println(strings.TrimSpace("  hello  "))               // ตัด space

	// ==========================================
	// 3. strconv — แปลง type
	// ==========================================
	fmt.Println("\n=== strconv ===")

	// string → int
	num, err := strconv.Atoi("42")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("string → int:", num)
	}

	// int → string
	str := strconv.Itoa(100)
	fmt.Println("int → string:", str)

	// ==========================================
	// 4. math — คณิตศาสตร์
	// ==========================================
	fmt.Println("\n=== math ===")

	fmt.Println("Pi:", math.Pi)
	fmt.Println("รากที่สอง 16:", math.Sqrt(16))
	fmt.Println("ยกกำลัง 2^10:", math.Pow(2, 10))
	fmt.Println("ค่าสูงสุด:", math.Max(10, 20))

	// ==========================================
	// 5. time — วันและเวลา
	// ==========================================
	fmt.Println("\n=== time ===")

	now := time.Now()
	fmt.Println("เวลาตอนนี้:", now.Format("2006-01-02 15:04:05"))
	fmt.Println("ปี:", now.Year())
	fmt.Println("เดือน:", now.Month())
	fmt.Println("วัน:", now.Day())

	message := fullInfo("Microwave", 2001)
	fmt.Println(message)
}
