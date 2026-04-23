package main

import "fmt"

func main() {
	studentGrade := map[string]int{} // Emtpy Map, key is string and value is int
	studentGrade["Alice"] = 90
	studentGrade["Bob"] = 85
	studentGrade["Charlie"] = 92

	fmt.Println("Student Grades:", studentGrade)

	alice_grade, ok := studentGrade["slice"]
	if ok {
		fmt.Println("Alice's grade:", alice_grade)
	} else {
		fmt.Println("Alice's grade not found", alice_grade)
	}

	key := "Bob"
	if grade, ok := studentGrade[key]; ok {
		fmt.Println("Key", key, "exists with value", grade)
	}

	delete(studentGrade, "Bob")
	fmt.Println("After deleting Bob:", studentGrade)

	configs := make(map[string]int)
	fmt.Printf("%T\n", configs) // Prints the datastructure. Use Printf here, not Println.

	var table map[string]int // Nil map, since it is just declared.
	if table == nil {
		fmt.Println("Table is nil")
	}

}
