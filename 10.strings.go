package main

import (
	"fmt"
	"strings"
)

func main() {

	s1 := "      Hello, World!      "
	s2 := strings.ToUpper(s1)
	println(s2)
	s3 := strings.ToLower(s1)
	println(s3)

	//Trim space from both ends
	s4 := strings.TrimSpace(s1)
	println(s4)

	// Check if a string ends with a specific suffix
	fmt.Println(strings.HasSuffix("prabir@gmail.com", "gmail.com"))

	// Check if a string starts with a specific prefix
	fmt.Println(strings.HasPrefix("prabir@gmail.com", "prabir"))

	fmt.Println(strings.Replace("India is Great", "India", "Bharat", 1))

	// strings.Split() returns a slice of substrings
	parts := strings.Split("prabir@gmail.com", "@")
	username, domain := parts[0], parts[1]
	fmt.Println("Username:", username)
	fmt.Println("Domain:", domain)

	fmt.Println(strings.Join(parts, ","))

}
