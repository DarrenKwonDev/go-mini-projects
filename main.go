package main

import "fmt"

func sumUp(numbers ...int) int {
	sum := 0
	// 두번째 인자로 배열의 값을 차례대로 반환한다
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func main() {
	total := sumUp(1, 2, 3, 4)
	fmt.Println(total)
}
