package main

import "fmt"

func main() {

	//Declaration of fixed size array
	var a [3]int   //int array of size 3
	fmt.Println(a) // Will print a zero int array of size 3

	//Short hand declaration [size] type {values}
	b := [3]int{12, 10, 10}
	fmt.Println(b)

	//Short hand declaration will initializing the first element in the array.
	c := [3]int{12}
	fmt.Println(c)

	// Automatically inferencing the size of the array

	d := [...]int{12}
	fmt.Println("size of the array ", len(d))

	//arrays are value types . Arrays are copied when assigned to a new array variable
	employeesArr := [4]string{"Ramesh", "Suresh", "Umesh", "Kamlesh"}
	employeesNewArr := employeesArr
	employeesNewArr[0] = "Sudesh"
	fmt.Println("Employees arr", employeesArr)
	fmt.Println("Employees new arr", employeesNewArr)

	//iterating arrays using range
	for i,v := range employeesNewArr {
		fmt.Printf("Employee at index %d is %s",i,v)
		fmt.Println()
	}


}
