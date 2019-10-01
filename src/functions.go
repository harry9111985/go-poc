package main

import (
	"fmt"
)

// Add function can only be used to add variables of type int and assigned the result of the operation to a variable of type int.
func Add(No1 int, No2 int) int {
	return No1 + No2
}

// Area and perimeter as return values
func calcRectProps(length, breadth int) (int, int) {

	return length * breadth, 2 * (length + breadth)
}

// Area and perimeter as named return values
func calcSquareProps(side int) (area int, perimeter int) {

	area =  side * side
	perimeter = 4 * side
	return
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

	// Assignment to change the values of existing variables
	length = 10
	breadth = 18

	// This is blank identifier expression . This is used when the caller calls a multi return value function and is interested in one or few
	//return values (in this case only area and not perimeter)
	area, _ = calcRectProps(length, breadth)

	fmt.Println("The area is ", area)

	side := 5
	squareArea,squarePeri := calcSquareProps(side)

	fmt.Println("Area and perimeter of square ",squareArea,squarePeri)

}
