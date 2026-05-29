# LeetCode Easy — Notes & Patterns

> อัปเดตล่าสุด: 2026-05-30

---

## #1 — Two Sum

**Link:** https://leetcode.com/problems/two-sum

### โจทย์
```
ให้ array ของตัวเลข และ target
หา index ของเลข 2 ตัว ที่บวกกันแล้วได้ target

Input:  nums = [2, 7, 11, 15], target = 9
Output: [0, 1]
```

### วิธีคิด
```
1. สร้าง map ว่าง (seen)
2. วนทุกตัวใน nums
3. คำนวณ need = target - num
4. เช็คว่า need อยู่ใน seen ไหม?
   → ถ้ามี: return [index ที่เจอ, i]
   → ถ้าไม่มี: เก็บ seen[num] = i
```

### Solution
```go
func twoSum(nums []int, target int) []int {
    seen := map[int]int{}
    for i, num := range nums {
        need := target - num
        if index, exists := seen[need]; exists {
            return []int{index, i}
        }
        seen[num] = i
    }
    return []int{}
}
```

### Pattern ที่ใช้
> **HashMap Lookup** — เก็บค่าที่เคยเจอ แล้วเช็คว่ามีคู่ไหม

| | |
|--|--|
| Time  | O(n) — วนครั้งเดียว |
| Space | O(n) — เก็บใน map |

---

## #2 — Valid Anagram

**Link:** https://leetcode.com/problems/valid-anagram

### โจทย์
```
ให้ string 2 ตัว s และ t
เช็คว่าเป็น Anagram กันไหม
(Anagram = ใช้ตัวอักษรเหมือนกันทุกตัว แค่สลับที่)

Input:  s = "anagram", t = "nagaram"
Output: true

Input:  s = "rat", t = "car"
Output: false
```

### วิธีคิด
```
1. สร้าง map ว่าง (seen)
2. วน s → เพิ่มนับ seen[ch]++
3. วน t → ลดนับ seen[ch]--
4. เช็คทุกค่าใน seen
   → ถ้าค่าไหน != 0: return false
5. return true
```

### Solution
```go
func isAnagram(s string, t string) bool {
    seen := map[rune]int{}
    for _, ch := range s {
        seen[ch]++
    }
    for _, ch := range t {
        seen[ch]--
    }
    for _, count := range seen {
        if count != 0 {
            return false
        }
    }
    return true
}
```

### Pattern ที่ใช้
> **HashMap Count** — นับและเปรียบเทียบจำนวน

| | |
|--|--|
| Time  | O(n) — วน 3 รอบ |
| Space | O(n) — เก็บใน map |

---

## สรุป Pattern ที่เจอ

| Pattern | ใช้เมื่อ | ข้อที่เจอ |
|---------|---------|---------|
| HashMap Lookup | หาคู่/ตรวจสอบว่ามีอยู่ไหม | Two Sum |
| HashMap Count | นับและเปรียบเทียบ | Valid Anagram |

---

## สิ่งที่ต้องจำ

```go
// วน string ทีละตัวอักษร
for _, ch := range s { }   // ch เป็น rune

// map นับตัวอักษร
seen := map[rune]int{}
seen[ch]++   // เพิ่ม
seen[ch]--   // ลด

// map หา index
seen := map[int]int{}
seen[num] = i            // เก็บ
index, exists := seen[key]  // ดึง
```
