package main

import "fmt"

func main() {

	// Declaration "var <var-name> <type>
	var age int

	// The age should be zero as its "int" zero value
	fmt.Println("The age is " ,age);

	// Declaration with an initial value
	var height int = 170

	fmt.Println(" The height is " ,height)

	// Declaration with an initial value doesnt need the type as Go inferences the type based on the initial value
	var size = "XL"
	fmt.Println("The size is ", size)


}
