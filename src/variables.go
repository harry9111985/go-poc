package main

import "fmt"

func main() {

	fmt.Println("Go - Variables exercise : start")
	// Declaration "var <var-name> <type>
	var age int

	// The age should be zero as its "int" zero value
	fmt.Println("The age is ", age);

	// Declaration with an initial value
	var height int = 170

	fmt.Println(" The height is ", height)

	// Declaration with an initial value doesnt need the type as Go inferences the type based on the initial value
	var size = "XL"
	fmt.Println("The size is ", size)

	//Multiple variables for the same type can be initialized in the same line. As they are initialized a type is not needed
	// and implicitly inferred by go based on the initialized value.
	var firstName, lastName = "Harish", "Subramaniam"

	fmt.Println("the Full name is ", firstName, lastName)

	/**
	  multiple variables declaration with different types
	 */
	var (
		block  = 5
		plot   = 76
		street = "R.A Kidwai Road"
	)

	fmt.Println("The address is ", block, plot, street)

	/**
	   Short hand initialization
	 */
	result := "Good"

	fmt.Println("The result of the work done is ", result)

	// name, age := "Harish" will throw error as initialization needs to be done for name and age both

	// Multiple assignment on the variables
	// This kinda multiple assignment is only allowed if there atleast one new variable declared in the same
	a,b := 80,90 // a and b are both new
	fmt.Println(" the values for a and b" ,a , b)
	b,c := 90,100 // c is new
	fmt.Println(" the values for b and c" ,b , c)

	//Conversion (explicit) is needed to perform an operation between 2 variables of different data types
	i := 1
	var j float32 = 3.48
	k := j + float32(i)

	fmt.Println("Sum is ",k)

	fmt.Println("Go - Variables exercise : end")


}
