/**
 the src relative folder in this case calculator/employee will be the import to use this package
 */
package employee

import "fmt"

/**
 the func outside the main package should start with a capital letter to be exposed in a main function
 */
func CalculateSalary(wagePerHour float32, noOfHoursPerDay int, noOfDays int) float64 {

	return float64(wagePerHour * (float32(noOfHoursPerDay)) * float32(noOfDays))
}

/**
This is the first function called when this package is imported as part of main package.
 */
func init() {
	fmt.Println("Initializing employee calculator package")
}