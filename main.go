package main

import (
	"fmt"

	"github.com/DarrenKwonDev/learnGo/accounts"
)

func main() {
	account := accounts.NewAccount("darrenkwon")
	fmt.Println(account)
}
