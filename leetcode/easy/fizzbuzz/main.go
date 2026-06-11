package main

import (
	"fmt"
	"strconv"
)

func fizzBuzz(n int) []string {

	result := []string{}
	for i := 0; i < n; i++ {
		if (i+1)%3 == 0 && (i+1)%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if (i+1)%3 == 0 {
			result = append(result, "Fizz")

		} else if (i+1)%5 == 0 {
			result = append(result, "Buzz")

		} else {
			result = append(result, strconv.Itoa(i+1))
		}
	}
	return result
}

func main() {
	fmt.Println(fizzBuzz(15))
	// ["1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"]
}
