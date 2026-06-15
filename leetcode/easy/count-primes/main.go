package main

import "fmt"

func isPrime(n int) bool {
	// วน j จาก 2 ถึง n-1
	// ถ้า n หาร j ลงตัว → return false
	// ถ้าวนจนครบแล้วไม่มีตัวไหนหารลงตัว → return true
	for j := 2; j < n; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}

func countPrimes(n int) int {
	// เขียนตรงนี้!
	count := 0
	for i := 2; i < n; i++ {

		if isPrime(i) {
			count++
		}
	}

	return count
}

func main() {
	fmt.Println(countPrimes(10)) // 4  (2, 3, 5, 7)
	fmt.Println(countPrimes(0))  // 0
	fmt.Println(countPrimes(1))  // 0
}
