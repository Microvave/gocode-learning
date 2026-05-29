package main

import "fmt"

func isAnagram(s string, t string) bool {
	seen := map[rune]int{}

	// Step 1: วน s → เพิ่มนับ
	for _, ch := range s {
		seen[ch]++ // เพิ่ม 1 ทุกครั้งที่เจอตัวอักษร
	}

	// Step 2: วน t → ลดนับ
	for _, ch := range t {
		seen[ch]-- // ลด 1 ทุกครั้งที่เจอตัวอักษร
	}

	// Step 3: เช็คว่าทุกค่าเป็น 0
	for _, count := range seen {
		if count != 0 {
			return false // มีตัวไหนไม่เท่ากัน
		}
	}

	return true
}

func main() {
	fmt.Println(isAnagram("anagram", "nagaram")) // true
	fmt.Println(isAnagram("rat", "car"))         // false
	fmt.Println(isAnagram("hello", "world"))     // false
}
