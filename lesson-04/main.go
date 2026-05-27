package main

import "fmt"

func main() {

	// ==========================================
	// 1. ARRAY — ขนาดตายตัว เปลี่ยนขนาดไม่ได้
	// ==========================================
	fmt.Println("=== Array ===")

	var scores [5]int = [5]int{10, 20, 30, 40, 50}

	fmt.Println("ทั้งหมด:", scores)
	fmt.Println("ตัวแรก:", scores[0])
	fmt.Println("ตัวสุดท้าย:", scores[4])
	fmt.Println("จำนวน:", len(scores))

	// ==========================================
	// 2. SLICE — เหมือน array แต่ขยายได้ (ใช้บ่อยกว่ามาก!)
	// ==========================================
	fmt.Println("\n=== Slice ===")

	fruits := []string{"แอปเปิ้ล", "กล้วย", "มะม่วง"}
	fmt.Println("fruits:", fruits)
	fmt.Println("จำนวน:", len(fruits))

	// เพิ่มของเข้าไป
	fruits = append(fruits, "สตรอเบอร์รี่")
	fruits = append(fruits, "องุ่น")
	fmt.Println("หลัง append:", fruits)
	fmt.Println("จำนวนใหม่:", len(fruits))

	// ดึงบางส่วน (slice of slice)
	fmt.Println("ตัวที่ 1-2:", fruits[0:2])
	fmt.Println("ตัวที่ 3 เป็นต้นไป:", fruits[2:])

	// ==========================================
	// 3. MAP — เก็บข้อมูลแบบ key: value
	// ==========================================
	fmt.Println("\n=== Map ===")

	person := map[string]string{
		"name": "Microwave",
		"city": "Lampang",
		"job":  "Developer",
	}

	fmt.Println("ทั้งหมด:", person)
	fmt.Println("ชื่อ:", person["name"])
	fmt.Println("เมือง:", person["city"])

	// เพิ่ม / แก้ไข
	person["age"] = "25"
	person["city"] = "Bangkok"
	fmt.Println("หลังแก้:", person)

	// ลบ
	delete(person, "job")
	fmt.Println("หลังลบ job:", person)

	// เช็คว่า key มีอยู่ไหม
	value, exists := person["name"]
	fmt.Println("มี name ไหม:", exists, "->", value)

	value2, exists2 := person["salary"]
	fmt.Println("มี salary ไหม:", exists2, "->", value2)

	// ==========================================
	me := map[string]string{
		"name":  "wave",
		"color": "darkBlue",
		"food":  "noodle",
	}
	fmt.Println("ทั้งหมด:", me)
	fmt.Println("ชื่อ:", me["name"])
	fmt.Println("สี:", me["color"])
	fmt.Println("อาหาร:", me["food"])
}
