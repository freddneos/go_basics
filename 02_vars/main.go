package main

import "fmt"

func main() {
	//var name = "Fredd"
	//shorthand
	name := "Fredd Bezerra"
	var age = 29
	const isCool = true
	var size = 2.3

	fmt.Println(name, age, isCool, size)
	//showing the type
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
}
