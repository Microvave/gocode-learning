package main

import "fmt"

func maxProfit(prices []int) int {
	minPrice := prices[0] // ราคาต่ำสุดที่เคยเห็น (เริ่มจากวันแรก)
	maxProfit := 0        // กำไรสูงสุดที่ทำได้

	for _, price := range prices {
		// 1. ถ้า price วันนี้ ถูกกว่า minPrice → อัปเดต minPrice
		// 2. คำนวณกำไรถ้าขายวันนี้ = price - minPrice
		//    ถ้ามากกว่า maxProfit → อัปเดต maxProfit
		if price < minPrice {
			minPrice = price
		}

		if price-minPrice > maxProfit {
			maxProfit = price - minPrice
		}
	}

	return maxProfit
}

func main() {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4})) // 5
	fmt.Println(maxProfit([]int{7, 6, 4, 3, 1}))    // 0
}
