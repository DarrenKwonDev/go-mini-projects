package main

import "fmt"

func main() {
	a := 5
	b := &a // b는 포인터다
	fmt.Println(b)
	fmt.Println(*b) // 역참조
}
