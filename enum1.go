package main

import "fmt"

type orderStatus int

const (
	Pending orderStatus = iota + 1 // iota starts at 0, so we add 1 to start from 1
	Processing
	Shipped
	Delivered
	Canceled
)

func getOrderStatus(status orderStatus) string {
	switch status {
	case Pending:
		return "We have received your order and it is pending."
	case Processing:
		return "Your order is being processed."
	case Shipped:
		return "Your order has been shipped."
	case Delivered:
		return "Your order has been delivered."
	case Canceled:
		return "Your order has been canceled."
	default:
		return "Unknown order status."
	}
}

func main() {
	// Create a new order
	currentStatus := Pending //currentStatus is of type orderStatus and is initialized to Pending

	fmt.Println("Initial Status:", getOrderStatus(currentStatus))

	// Simulate the order moving forward
	currentStatus = Shipped
	fmt.Println("Update:", getOrderStatus(currentStatus))

	// Example of using the underlying integer value
	fmt.Printf("The numeric value of 'Shipped' is: %d\n", currentStatus)
}
