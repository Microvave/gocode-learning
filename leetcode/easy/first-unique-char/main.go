package main

import "fmt"

func firstUniqChar(s string) int {
	// เขียนตรงนี้! คืน index ของตัวอักษรตัวแรกที่ไม่ซ้ำ
	// ถ้าไม่มีเลย → คืน -1
	count := map[rune]int{}

	// รอบ 1: นับ
	for _, ch := range s {
		count[ch]++
	}
	// รอบ 2: วนหา "ตัวแรกที่ไม่ซ้ำ"
	for i, ch := range s {
		if count[ch] == 1 { // ← เงื่อนไขตรงนี้ต่างจากตัวอย่าง
			return i
		}
	}

	return -1
}

func main() {
	fmt.Println(firstUniqChar("leetcode"))     // 0  (l)
	fmt.Println(firstUniqChar("loveleetcode")) // 2  (v)
	fmt.Println(firstUniqChar("aabb"))         // -1
}
