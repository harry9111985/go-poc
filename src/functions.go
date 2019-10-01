package main

import "fmt"

// Add function can only be used to add variables of type int and assigned the result of the operation to a variable of type int.
func Add(No1 int, No2 int) int {
	return No1 + No2
}

// Area and perimeter as return values
func calcRectProps(length, breadth int) (int, int) {

	return length * breadth, 2 * (length + breadth)
}

func main() {

	No1 := 8
	No2 := 9
	fmt.Println("Sum is ", Add(No1, No2))

	length := 8
	breadth := 9

	// assignment of the result from a function which returns multiple values.
	area, perimeter := calcRectProps(length, breadth)

	fmt.Println("The area and perimeter of the rectangle ", area, perimeter)
}
