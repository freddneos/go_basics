package main

import "fmt"

func main() {
	//Define map

	emails := make(map[string]string)

	//Assing key, values
	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"
	emails["Mike"] = "mike@gmail.com"

	fmt.Println(emails)
	fmt.Println(len(emails))
	fmt.Println(emails["Bob"])

	//delete
	delete(emails, "Bob")
	fmt.Println(emails)

}
