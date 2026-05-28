package main

import (
	"errors"
	"fmt"
)

// ==========================================
// 1. ERROR พื้นฐาน — return error กลับมา
// ==========================================

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("หารด้วยศูนย์ไม่ได้!")
	}
	return a / b, nil // nil = ไม่มี error
}

// ==========================================
// 2. สร้าง Custom Error
// ==========================================

func getAge(age int) (int, error) {
	if age < 0 {
		return 0, errors.New("อายุติดลบไม่ได้")
	}
	if age > 150 {
		return 0, errors.New("อายุเกิน 150 ปีไม่น่าเป็นไปได้")
	}
	return age, nil
}

// ==========================================
// test
// ==========================================
func checkPassword(password string) (string, error) {
	if len(password) < 8 {
		return "", errors.New("รหัสผ่านต้องมีอย่างน้อย 8 ตัว")
	}
	return "รหัสผ่านผ่าน!", nil
}

// ==========================================
// 3. MAIN
// ==========================================

func main() {

	// เรียก function ที่อาจเกิด error
	fmt.Println("=== Error Handling ===")

	// กรณีสำเร็จ
	result, err := divide(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("10 / 2 =", result)
	}

	// กรณี error
	result2, err2 := divide(10, 0)
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println("10 / 0 =", result2)
	}

	fmt.Println("\n=== Custom Error ===")

	// อายุปกติ
	age, err3 := getAge(25)
	if err3 != nil {
		fmt.Println("Error:", err3)
	} else {
		fmt.Println("อายุ:", age)
	}

	// อายุติดลบ
	age2, err4 := getAge(-5)
	if err4 != nil {
		fmt.Println("Error:", err4)
	} else {
		fmt.Println("อายุ:", age2)
	}

	// อายุเกิน
	age3, err5 := getAge(200)
	if err5 != nil {
		fmt.Println("Error:", err5)
	} else {
		fmt.Println("อายุ:", age3)
	}

	msg, err := checkPassword("1234")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(msg)
	}

	msg2, err2 := checkPassword("mypassword123")
	if err2 != nil {
		fmt.Println("Error:", err2)
	} else {
		fmt.Println(msg2)
	}
}
