package main

import "fmt"

// Declaring ID as New type whose underlying type is int
type ID int

type Employee struct {
	ID         ID
	firstName  string
	lastName   string
	Age        int
	Salary     int
	Department string
}

// Implementing the Stringer interface for the Employee struct and new type ID by defining the String method of the Stringer interface.
// Stringer interface definition is as follows:
//
//	type Stringer interface {
//	    String() string
//	}

func (e *Employee) String() string {
	return fmt.Sprintf("Employee[ID: %d, Name: %s %s, Age: %d, Salary: %d, Department: %s]", e.ID, e.firstName, e.lastName, e.Age, e.Salary, e.Department)
}

func (id ID) String() string {
	return fmt.Sprintf("EmployeeID[%d]", id)
}

func main() {
	// John is a pointer to an Employee struct, initialized with the ID 1, first name "John", last name "Doe", age 30, salary 50000, and department "Engineering".
	John := &Employee{ID: 1, firstName: "John", lastName: "Doe", Age: 30, Salary: 50000, Department: "Engineering"}

	//Go tries it best to print the struct, but it doesn't know how to format it,
	// so it prints the default struct representation, which includes the field values.
	// If the type Employee had not implemented the Stringer interface, the output would be something like:
	// Employee Details &{1 John Doe 30 50000 Engineering}
	// Printf() is expecting a String, so since Employee implements the Stringer interface,
	// we can pass Employee directly to Printf() and it will call the String() method.

	fmt.Println("Employee Details", John)

	// Since ID implements the Stringer interface, we can pass it directly to Printf()
	// and it will call the String() method.
	fmt.Println("Employee ID", John.ID)

}
