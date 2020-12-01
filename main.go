package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var num int = 10
	fmt.Println(unsafe.Sizeof(num))
	var change float32 = float32(num)
	fmt.Println(change)
	fmt.Println(unsafe.Sizeof(change))
}
