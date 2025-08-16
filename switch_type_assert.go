package main

import "fmt"

func typeAsserting(x any) {
	switch x.(type) {
	case bool:
		fmt.Println("Bool")
	case int:
		fmt.Println("Int")
	case string:
		fmt.Println("String")
	default:
		fmt.Println("None")
	}
}

type Person struct {
	Name string
}

func main() {
	var lol string = "LOL"
	typeAsserting(lol)
}
