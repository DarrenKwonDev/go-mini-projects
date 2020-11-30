package main

import (
	"fmt"

	mydict "github.com/DarrenKwonDev/learnGo/dict"
)

func main() {
	dictionary := mydict.Dictionary{"first": "word"}
	definition, err := dictionary.Search("first")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(definition)
	}
	dictionary.Add("second", "whwhwh")
	fmt.Println(dictionary)
	dictionary.Delete("first")
	fmt.Println(dictionary)
}
