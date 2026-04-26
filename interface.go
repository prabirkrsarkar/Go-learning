package main

type Employee struct {
	ID         int
	firstName  string
	lastName   string
	Age        int
	Salary     int
	Department string
}

type BusinessPerson struct {
	firstName string
	lastName  string
	business  string
}

type Person interface {
	getFullName() string
}

// Both Employee and BusinessPerson implement the Person interface by defining the getFullName method.
func (e *Employee) getFullName() string {
	return e.firstName + " " + e.lastName
}

func (b *BusinessPerson) getFullName() string {
	return b.firstName + " " + b.lastName
}

func main() {
	employee := &Employee{ID: 1, firstName: "John", lastName: "Doe", Age: 30, Salary: 50000, Department: "Engineering"}
	businessPerson := &BusinessPerson{firstName: "Jane", lastName: "Smith", business: "Tech Startup"}

	var person1 Person = employee
	var person2 Person = businessPerson

	println("Employee Full Name:", person1.getFullName())
	println("Business Person Full Name:", person2.getFullName())
}
