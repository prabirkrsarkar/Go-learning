package main

import "fmt"

// Struct is a composite type
type Employee struct {
	ID         int
	firstName  string
	lastName   string
	Age        int
	Salary     int
	Department string
}

// Constructor like function
// This is a factory/constructor pattern
// Go doesn’t have constructors like Java/C++, so we use functions.
// Create an Employee struct and return its address (pointer). Returning pointer helps us to
// Avoid copying large structs
// Allow modification of the same object
// More efficient in real systems
func newEmployee(id int, firstName, lastName string, age, salary int, department string) *Employee {
	return &Employee{
		ID:         id,
		firstName:  firstName,
		lastName:   lastName,
		Age:        age,
		Salary:     salary,
		Department: department,
	}
}

func main() {
	//employee is a pointer to an Employee struct object
	employee := newEmployee(1, "John", "Doe", 30, 50000, "Engineering")

	// Dereference the pointer to access the struct value
	fmt.Println("Employee:", *employee)
}
