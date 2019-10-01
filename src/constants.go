package main

import "fmt"

func main() {

	fmt.Println("Go - Constants exercise : start")

	// Declaring constants with an initial value . Constants are always initialized with a value . Otherwise compiler will throw error.
	const MinAge = 30

	fmt.Println("Minimum age is ", MinAge)

	const DefaultName  = "XYZ"
	type myString string

	// This assignment is possible because DefaultName is untyped constant. It wont work if DefaultName is typed.
	var name myString = DefaultName
	someOtherName := "DEF"

	fmt.Println("Some other name", someOtherName)

	// This will be not be not be allowed as someOtherName has a different type and not a constant
	//name = someOtherName

	fmt.Println("Name with a custom string type initialized with Default Name", name)

	fmt.Println("Go - Constants exercise : end")
}
