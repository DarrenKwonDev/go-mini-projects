package main

import "fmt"

func amITen(age int) bool {
	switch koreanAge := age + 2; koreanAge {
	case 10:
		return true
	}
	return false
}

func main() {
	fmt.Println(amITen(16))
}
