package main

import (
	"fmt"

	"github.com/DarrenKwonDev/learnGo/accounts"
)

func main() {
	account := accounts.NewAccount("darren")
	// account.Deposit(10)
	// fmt.Println(account.Balance())
	// account.NewOwner("lee")
	// fmt.Println(account)
	// fmt.Println(account.Owner())
	fmt.Println(account)
	// err := account.Withraw(20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
}
