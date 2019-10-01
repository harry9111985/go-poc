package main

import (
	"fmt"
	"calculator/employee"
)

func main() {

	fmt.Println("Calculator package")
	fmt.Println("Salary of the employee " , employee.CalculateSalary(67.50,7,20))

}
