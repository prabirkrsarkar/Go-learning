package main

import "fmt"

type Contact struct { // A custom, user-defined type to represent a contact in a contact management system
	ID    int
	Name  string
	Email string
	Phone string
}

// A slice of a user-defined type, in this case, a slice of Contact structs.
var contactList []Contact // A slice to hold the list of contacts in the contact management system

var contactIndexByName map[string]int // A map to index contacts in the slice contactList by their name for quick lookup
// The value in map[string]int is an int because it is being used as an index tracker.
// Since the primary data store is a slice ([]Contact), the int value in the map stores
// the "position" (the index) where that specific contact lives inside the slice.

var nextID int = 1 // A variable to keep track of the next unique ID to assign to a new contact when adding to the contact management system

// Initialize the contact management system by creating an empty slice for contacts and an empty map for indexing contacts by name.
// Go automatically calls the init() function before the main() function is executed,
// so this is a good place to set up our initial data structures for the
// contact management system.
func init() {
	contactList = []Contact{}                 // Initialize the contact list as an empty slice of Contact structs
	contactIndexByName = make(map[string]int) // Initialize the map to index contacts by name as an empty map
}

func addContact(name, email, phone string) {
	if _, exists := contactIndexByName[name]; exists { // Check if a contact with the same name already exists in the contact management system using the map for quick lookup
		fmt.Println("Contact with the name " + name + " already exists. Please use a different name.")
		return // If a contact with the same name exists, do not add a new contact and return early
	}

	newContact := Contact{ // Create a new Contact struct with the provided name, email, and phone, and assign it a unique ID using nextID
		ID:    nextID,
		Name:  name,
		Email: email,
		Phone: phone,
	}

	nextID++ // Increment the nextID variable to ensure the next contact added will have a unique ID
	contactList = append(contactList, newContact)

	// contactIndexByName[name] = newContact.ID - 1 // This would break when we delete a contact from the middle of the slice, since the index of the contact in the slice would change. So, instead, we should store the index of the contact in the slice.
	contactIndexByName[name] = len(contactList) - 1 // Store the index of the new contact in the slice in the map for quick lookup by name
	fmt.Println("New Contact, " + newContact.Name + " added to the Contact List")
}

// Note, we are returning a pointer to the Contact struct from the slice,
// so that we can modify the contact's details if needed.
// If we return a copy of the Contact struct, any modifications made to
// the returned struct would not affect the original contact in the slice.
func getContactByName(name string) *Contact {
	if index, exists := contactIndexByName[name]; exists { // Check if a contact with the given name exists in the contact management system using the map for quick lookup
		return &contactList[index] // If it exists, return a pointer to the Contact struct object from the slice using the index stored in the map
	}
	fmt.Println("Contact with the name " + name + " not found.")
	return nil // If no contact with the given name exists, return nil pointer
}

func listContacts() {
	fmt.Printf("Contact List:")
	if len(contactList) == 0 {
		fmt.Println("No contacts found.")
		return // If the contact list is empty, print a message and return early
	}
	for index, contact := range contactList {
		fmt.Printf("%d. ID: %d, Name: %s, Email: %s, Phone: %s\n", index+1, contact.ID, contact.Name, contact.Email, contact.Phone) // Print the details of each contact in the contact list with a numbered index
	}
}

func main() {

	listContacts() // List contacts when the program starts, which will show an empty contact list
	addContact("Alice", "alice@example.com", "123-456-7890")
	addContact("Alice", "alice2@example.com", "987-654-3210") // This will not be added since a contact with the name "Alice" already exists
	addContact("Bob", "bob@example.com", "555-555-5555")
	listContacts()

	contactptr := getContactByName("Prabir") // Get a contact by name, which will return a pointer to the Contact struct for "Prabir"
	if contactptr == nil {
		fmt.Println("Contact not found.")
	} else {
		// Note, that Go automatically dereferences the pointer when we access the fields of the Contact struct, so we can use contactptr.Name instead of (*contactptr).Name to access the Name field of the Contact struct.
		fmt.Printf("Contact found: ID: %d, Name: %s, Email: %s, Phone: %s\n", contactptr.ID, contactptr.Name, contactptr.Email, contactptr.Phone) // Print the details of the contact if found
	}
}
