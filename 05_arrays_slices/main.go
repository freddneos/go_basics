package main

import "fmt"

func main() {
	// Arrays
	var fruitArray [2]string

	//Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	//Declare and assing
	thingsArray := [2]string{"pen", "pencil"}
	//Undimensioned
	otherThingsArray := []string{"pen", "pencil"}

	fmt.Println(fruitArray)
	fmt.Println(fruitArray[1])
	fmt.Println(thingsArray)
	fmt.Println(otherThingsArray)
}
