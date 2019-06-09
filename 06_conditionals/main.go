package main

import "fmt"

func main() {
	x, y := 15, 10

	if x <= y {
		fmt.Printf(" %d is less than or equal %d \n", x, y)
	} else {
		fmt.Printf(" %d is less than  %d \n", y, x)
	}

	//switch
	color := "green"

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("color is bot blue or red")
	}

}
