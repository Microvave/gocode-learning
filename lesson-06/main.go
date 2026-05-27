package main

import "fmt"

// ==========================================
// INTERFACE — กำหนดว่า struct ต้องมี method อะไร
// เหมือน "สัญญา" ว่าจะต้องทำได้
// ==========================================

type Animal interface {
	Sound() string
	Name() string
}

// ==========================================
// หลาย struct ที่ทำตาม interface เดียวกัน
// ==========================================

type Dog struct {
	TheName string
}

func (d Dog) Sound() string { return "โฮ่ง!" }
func (d Dog) Name() string  { return d.TheName }

type Cat struct {
	TheName string
}

func (c Cat) Sound() string { return "เมี้ยว!" }
func (c Cat) Name() string  { return c.TheName }

type Duck struct {
	TheName string
}

func (d Duck) Sound() string { return "ก้าก!" }
func (d Duck) Name() string  { return d.TheName }

type Bird struct {
	TheName string
}

func (b Bird) Sound() string { return "จุกกรู" }
func (b Bird) Name() string  { return b.TheName }

// ==========================================
// ประโยชน์ของ Interface — รับ type อะไรก็ได้
// ที่ทำตาม interface
// ==========================================

func Describe(a Animal) {
	fmt.Printf("%s พูดว่า: %s\n", a.Name(), a.Sound())
}

// ==========================================
// MAIN
// ==========================================

func main() {

	fmt.Println("=== Interface ===")

	dog := Dog{TheName: "บัดดี้"}
	cat := Cat{TheName: "วิสกี้"}
	duck := Duck{TheName: "โดนัลด์"}
	brid := Bird{TheName: "ก้อน"}
	// เรียกตรงๆ
	Describe(dog)
	Describe(cat)
	Describe(duck)
	Describe(brid)

	// ใส่ใน slice แล้ววนลูป
	fmt.Println("\n=== วนลูป Animals ===")

	animals := []Animal{dog, cat, duck, brid}

	for _, animal := range animals {
		Describe(animal)
	}

}
