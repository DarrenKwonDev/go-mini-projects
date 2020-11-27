package main

import "fmt"

type person struct {
	name    string
	age     int
	favFood []string
}

func main() {
	me := map[string]string{"name": "darren", "job": "entrepreneur"}
	fmt.Println(me)

	yourFav := []string{"cheeze", "wine"}
	you := person{name: "mary", age: 18, favFood: yourFav}
	fmt.Println(you)
	fmt.Println(you.name)
}
