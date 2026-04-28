package main

import "fmt"

// Payable is a custom interface that represents any entity that can calculate its pay.
type Payable interface {
	fmt.Stringer           // Interface embedding, For better output formatting
	CalculatePay() float64 // Calculate monthly pay
}

// SalariedEmployee represents an employee with a fixed annual salary.
type SalariedEmployee struct {
	Name         string
	AnnualSalary float64
}

// SalariedEmployee implements the Payable interface by providing a CalculatePay method that
// calculates the monthly pay based on the annual salary.
func (se *SalariedEmployee) CalculatePay() float64 {
	return se.AnnualSalary / 12.0 // Monthly pay
}

// SalariedEmployee also has to implement the fmt.Stringer interface by providing a String method that
// returns a formatted string for better output formatting.
// Note that fmt.Printf returns (int, error) but the String() method should
// return a string. Hence, you need to use fmt.Sprintf instead, which returns a formatted string.

func (se *SalariedEmployee) String() string {
	return fmt.Sprintf("Salaried: %s (Annual: $%.2f)", se.Name, se.AnnualSalary)
}

// HourlyEmployee represents an employee paid by the hour.
type HourlyEmployee struct {
	Name        string
	HourlyRate  float64
	HoursWorked float64 // Hours worked in the pay period
}

// Implement the Payable interface for HourlyEmployee
func (he *HourlyEmployee) CalculatePay() float64 {
	return he.HourlyRate * he.HoursWorked
}

// HourlyEmployee also has to implement the fmt.Stringer interface by providing a String method that
// returns a formatted string for better output formatting.
func (he *HourlyEmployee) String() string {
	return fmt.Sprintf("Hourly: %s (Rate: $%.2f, Hours: %.2f)", he.Name, he.HourlyRate, he.HoursWorked)
}

// CommissionEmployee represents an employee who earns a base salary plus a commission on sales.
type CommissionEmployee struct {
	Name           string
	BaseSalary     float64
	CommissionRate float64
	SalesAmount    float64 // Total sales in the pay period
}

// Implement the Payable interface for CommissionEmployee
func (ce *CommissionEmployee) CalculatePay() float64 {
	return ce.BaseSalary + (ce.CommissionRate * ce.SalesAmount)
}

// CommissionEmployee also has to implement the fmt.Stringer interface by providing a String method that
// returns a formatted string for better output formatting.
// Again Note that fmt.Printf returns (int, error) but the String() method should
// return a string. Hence, you need to use fmt.Sprintf instead, which returns a formatted string.

func (ce *CommissionEmployee) String() string {
	return fmt.Sprintf("Commission: %s (Base: $%.2f, Rate: %.2f, Sales: $%.2f)", ce.Name, ce.BaseSalary, ce.CommissionRate, ce.SalesAmount)
}

// printEmployeeDetails is a generic function that prints the details
// of any employee that implements the fmt.Stringer interface.
// Since employee satisfies fmt.Stringer, the fmt package will automatically invoke String()
// when you use it in formatted output.
func printEmployeeDetails[P fmt.Stringer](employee P) {
	fmt.Printf(" -- Processing: %s\n", employee) // Go checks if employee implements fmt.Stringer and calls the String() method to get the string representation of the employee for printing.
}

// processPayroll function processes the payroll for any number of employees
// that implement the Payable interface.
func processPayroll(employee []Payable) {
	totalPayroll := 0.0
	for _, e := range employee {
		printEmployeeDetails(e) // Generic function to print employee details, demonstrating the use of interfaces and polymorphism.
		pay := e.CalculatePay() //Depending on the type of employee, the appropriate CalculatePay method will be called, demonstrating polymorphism.
		fmt.Printf("Monthly Pay : %f\n", pay)
		fmt.Printf(" -- Payroll processed: $%.2f\n", pay)
		totalPayroll += pay
	}
	fmt.Printf(" -- Total payroll: $%.2f\n", totalPayroll)
}

// Since Payable is an interface, it is an abstract type. Any concrete type that implements the
// Payable interface can be used as a Payable.
// Since the receivers in the struct methods are pointer receivers, we need to use pointers to the struct
// when creating instances of the employees. This allows us to call the CalculatePay method
// on the employee instances, which is required by the Payable interface.

// The processPayroll function can then process each employee in the slice,
// regardless of their specific type (SalariedEmployee, HourlyEmployee, or CommissionEmployee), as long as they
// implement the Payable interface. In this example, we create a slice of Payable that contains
// different types of employees (SalariedEmployee, HourlyEmployee, and CommissionEmployee).
// The processPayroll function can then process each employee in the slice, regardless of their specific type, as long as they implement the Payable interface.

func main() {
	fmt.Println("--- Welcome to the Payroll Processor! ---")

	employees := []Payable{ // Slice of pointers to structs that implement the Payable interface
		&SalariedEmployee{Name: "Alice", AnnualSalary: 60000},
		&HourlyEmployee{Name: "Bob", HourlyRate: 15, HoursWorked: 160},
		&CommissionEmployee{Name: "Charlie", BaseSalary: 30000, CommissionRate: 0.1, SalesAmount: 50000},
	}

	processPayroll(employees)
}
