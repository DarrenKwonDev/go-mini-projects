package main

import "fmt"

func addEvery(num ...int) int {
	var result int

	// num은 슬라이스입니다
	for _, val := range num {
		result += val
	}

	return result
}

func main() {
	num1, num2, num3, num4, num5 := 1, 2, 3, 4, 5
	nums := []int{10, 20, 30, 40}

	fmt.Println(addEvery(num1, num2, num3, num4, num5)) // 15
	fmt.Println(addEvery(nums...))
}
