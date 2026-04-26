package main

import "fmt"

// Employee represents an employee in the organization
// This kind of loosely represents a class in other object oriented languages
type Employee struct {
	ID         int
	firstName  string
	lastName   string
	Age        int
	Salary     int
	Department string
}

// Similar to a constructor in other object oriented languages
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

// We know that classes have methods.
// Here we use a method with a value receiver (e Employee)
// You can use a value receiver when you don't need to modify the receiver
func (e Employee) FullName() string {
	return e.firstName + " " + e.lastName
}

// Use a pointer receiver when you need to modify the receiver
// This will cause side effects since it will modify the original value
func (e *Employee) GiveRaise(amount int) {
	e.Salary += amount
}

// BUT, TO BE CONSISTENT, WE SHOULD USE POINTER RECEIVERS FOR ALL METHODS IRRESPECTIVE
// OF WHETHER THEY MODIFY THE RECEIVER OR NOT

func main() {

	//Note, employee_1 is a pointer to a particular value of Employee struct
	employee_1 := newEmployee(1001, "John", "Doe", 30, 50000, "Engineering")

	// Dereferencing is automatic when calling fields and methods on a pointer
	fmt.Printf("Employee ID: %d\n", employee_1.ID)
	fmt.Printf("Employee Full Name: %s\n", employee_1.FullName())

	// Creates an Employee with default zero values.
	// But if joe = nil then accessing fields or methods will cause a runtime panic
	joe := &Employee{}
	fmt.Printf("Employee ID: %d\n", joe.ID)
	fmt.Printf("Employee Full Name: %s\n", joe.FullName())
}
