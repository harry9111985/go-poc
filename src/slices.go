package main

import "fmt"

func main() {

	//creating a slice from an array
	a := [5]int{1, 2, 3, 4, 5}
	//Slice is created out of the first four elements a[0{start}] to a[3{end-1}]
	b := a[0:4]
	fmt.Println("Slice created out of an array ", b)

	//Creates an array and returns a slice.
	c := [] int{1, 2, 3}
	fmt.Println("Slice created out of short hand initialization", c)

	// A slice is a reference to the array . Changing the slice will change the underlying array.
	d := [5] string{"Ramesh", "Suresh", "kamlesh", "Sudesh", "Sukesh"}
	fmt.Println("Array before update :", d)
	e := d[0:2]
	e[0] = "Narayan"
	fmt.Println("Array after update :", d)

	//Create a slice using make function
	employeeIds := make([]int, 5, 5)
	fmt.Println("Initialized slice using make function ", employeeIds)

	//Copy function to garbage collect the underlying array
	employeeIdsNew := employeeIds[0:2]
	employeeIdsNew[0] = 1
	employeeIdsNewCopy := make([] int, len(employeeIdsNew), cap(employeeIdsNew))
	copy(employeeIdsNewCopy, employeeIdsNew)
	fmt.Println("Slice copy created out of the slice to free the underlying array of employeeIds ", employeeIdsNewCopy)

	//Variadic functions
	fmt.Println("Sum with 2 values ", sumWithDefaultValue(1, 1, 2))
	fmt.Println("Sum with 0 values", sumWithDefaultValue(1))

	//Passing a slice to a variadic function using ... notation
	fmt.Println("Sum of 3 values ",sumWithDefaultValue(1, []int {1,2,3}...))
}

func sumWithDefaultValue(defaultVal int, values ...int) int {
	if len(values) == 0 {
		return defaultVal
	} else {
		sum := 0
		// Note : The variable number of arguments in the variadic function is converted to a slice implicitly
		// Hence using variadic functions creation of an additional slice can be avoided.
		for _, v := range values {
			sum += v
		}
		return sum
	}
}
