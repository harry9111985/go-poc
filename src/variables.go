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

	//Multiple variables for the same type can be initialized in the same line. As they are initialized a type is not needed
	// and implicitly inferred by go based on the initialized value.
	var firstName,lastName = "Harish","Subramaniam"

	fmt.Println("the Full name is ",firstName,lastName)

}
