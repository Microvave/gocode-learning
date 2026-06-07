package main

import "fmt"

func reverseString(s []byte) {
	// เขียนตรงนี้! (กลับ slice ในตัวมันเอง ไม่ต้อง return)
	left := 0
	right := len(s) - 1

	for left < right {
		// ตรงนี้ → แทนที่จะ "เปรียบเทียบ" ให้ "สลับ" s[left] กับ s[right]
		// ใช้ a, b = b, a
		s[left], s[right] = s[right], s[left]
		left++
		right--
	}

}

func main() {
	s1 := []byte("hello")
	reverseString(s1)
	fmt.Println(string(s1)) // olleh

	s2 := []byte("Go")
	reverseString(s2)
	fmt.Println(string(s2)) // oG
}
