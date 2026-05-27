package main

import "fmt"

// ==========================================
// 1. STRUCT — เหมือน blueprint ของ object
// ==========================================

type Person struct {
	Name string
	Age  int
	City string
}

// ==========================================
// 2. METHOD — function ที่ผูกกับ struct
// ==========================================

func (p Person) Greet() string {
	return "สวัสดี ฉันชื่อ " + p.Name
}

func (p Person) IsAdult() bool {
	return p.Age >= 18
}

// ==========================================
// 3. STRUCT ซ้อน STRUCT
// ==========================================

type Address struct {
	Street string
	City   string
}

type Employee struct {
	Name    string
	Age     int
	Address Address // ใส่ struct ซ้อนกัน
	Salary  float64
}

func (e Employee) Info() string {
	return fmt.Sprintf("%s อายุ %d ปี อยู่ที่ %s เงินเดือน %.0f",
		e.Name, e.Age, e.Address.City, e.Salary)
}

type Car struct {
	Brand string
	Model string
	Year  int
}

func (c Car) Info() string {
	return fmt.Sprintf("%s %s ปี %d",
		c.Brand, c.Model, c.Year)
}

// ==========================================
// MAIN
// ==========================================

func main() {

	// สร้าง Person
	fmt.Println("=== Struct ===")

	p1 := Person{
		Name: "Microwave",
		Age:  25,
		City: "Lampang",
	}

	fmt.Println("ชื่อ:", p1.Name)
	fmt.Println("อายุ:", p1.Age)
	fmt.Println("เมือง:", p1.City)

	// เรียกใช้ method
	fmt.Println("\n=== Methods ===")
	fmt.Println(p1.Greet())
	fmt.Println("เป็นผู้ใหญ่ไหม:", p1.IsAdult())

	// แก้ไขค่า
	p1.City = "Bangkok"
	fmt.Println("ย้ายไป:", p1.City)

	// สร้าง Employee (struct ซ้อน struct)
	fmt.Println("\n=== Struct ใน Struct ===")

	e1 := Employee{
		Name: "Wave",
		Age:  25,
		Address: Address{
			Street: "Sukhumvit",
			City:   "Bangkok",
		},
		Salary: 50000,
	}

	fmt.Println(e1.Info())
	fmt.Println("ถนน:", e1.Address.Street)

	fmt.Println("\n=== test ===")
	c1 := Car{
		Brand: "Toyota",
		Model: "Camry",
		Year:  2025,
	}

	fmt.Println(c1.Info())

}
