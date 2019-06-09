package main

import "fmt"

//Defines a person struct
type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

//Greeting method value receiver
func (p Person) greet() string {
	return "Hello , my name is " + p.firstName + " " + p.lastName
}

func main() {
	//Init person struct
	person1 := Person{
		firstName: "Fredd",
		lastName:  "Bezerra",
		city:      "RJ",
		gender:    "male",
		age:       29,
	}

	fmt.Println(person1)
	fmt.Println(person1.firstName)
	person1.age++
	fmt.Println(person1.age)

}
