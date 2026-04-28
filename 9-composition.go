package main

import "fmt"

// In Go Inheritance is not supported, but you can achieve similar functionality by using composition.
// Composition allows you to build complex types by combining simpler ones.
// Instead of inheriting from a base class, you can embed one struct within another,
// allowing the outer struct to access the fields and methods of the embedded struct.
// Compostion enables a Has-A relationship, where one struct "has" another struct as a field,
// while Inheritance is an Is-A relationship.

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

// Let us create a method FullAddress() attached to the Address struct that returns the full address as a formatted string.
// We are using a value receiver for the FullAddress method because it does not modify the Address struct
// and allows us to call it on both value and pointer instances of Address.
func (a Address) FullAddress() string {
	if a.Street == "" && a.City == "" && a.State == "" && a.Zip == "" {
		return "No address provided"
	}
	return a.Street + ", " + a.City + ", " + a.State + " " + a.Zip
}

// With composition, whatever methods the Address struct has,
// the Customer struct can also access those methods directly, as if they were its own.
type Customer struct {
	Name            string
	Email           string
	BillingAddress  Address // Embedding the Address struct to achieve composition
	ShippingAddress Address // Embedding the Address struct to achieve composition
}

func (c Customer) PrintDetails() {
	fmt.Printf("Customer Name: %s\n", c.Name)
	fmt.Printf("Customer Email: %s\n", c.Email)
	fmt.Printf("Billing Address: %s\n", c.BillingAddress.FullAddress())   // Accessing method of the Composed type Address through the Customer struct
	fmt.Printf("Shipping Address: %s\n", c.ShippingAddress.FullAddress()) // Accessing method of the Composed type Address through the Customer struct
}

func main() {
	fmt.Printf("--- Welcome to the Customer Management System! ---\n")

	cust1 := Customer{
		Name:  "John Doe",
		Email: "john.doe@example.com",
		BillingAddress: Address{
			Street: "123 Main St",
			City:   "Anytown",
			State:  "CA",
			Zip:    "12345",
		},
		ShippingAddress: Address{
			Street: "456 Elm St",
			City:   "Othertown",
			State:  "CA",
			Zip:    "67890",
		},
	}

	cust1.PrintDetails()

	fmt.Printf("---Customer with Same Billing and Shipping Address---\n")

	cust2_Mainaddress := Address{
		Street: "789 Oak St",
		City:   "Sometown",
		State:  "CA",
		Zip:    "11223",
	}

	cust2 := Customer{
		Name:            "Jane Smith",
		Email:           "jane.smith@example.com",
		BillingAddress:  cust2_Mainaddress,
		ShippingAddress: cust2_Mainaddress,
	}
	cust2.PrintDetails()
}
