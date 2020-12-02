package main

import (
	"fmt"
	"strings"
)

func main() {
	arr := []string{"hello", "world", "and", "code"}
	str := "     he  ll    "

	fmt.Println(strings.Trim(str, "l"))
	fmt.Println(strings.TrimSpace(str))
	fmt.Println(strings.Join(arr, "|"))
}
