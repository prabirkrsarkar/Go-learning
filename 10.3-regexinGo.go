package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	// Regular expressions in Go are implemented using the "regexp" package.
	/*
		Do you want to handle errors explicitly ? -> If yes, use Compile. Usually used for dynamic expressions from userinput
		OR should the program panic if the regex is invalid? -> If yes, use MustCompile, usually used for hardcoded regex
	*/

	// Take user input for a regex pattern. Regex pattern is Dynamic
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter regex pattern: ")
	pattern, _ := reader.ReadString('\n')

	// Trim the Whitespace and newline character from the end of the input
	pattern = strings.TrimSpace(pattern)
	fmt.Println("You entered:", pattern)

	// Compile the regex pattern. This will check if the pattern is valid and prepare it for matching.
	// It returns a compiled regex object and an error if the pattern is invalid.
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Invalid regex:", err)
		return
	}

	fmt.Println("Regex compiled successfully")
	if re.MatchString("The quick brown fox jumps over the lazy dog.") {
		fmt.Println("The regex matches the string 'The quick brown fox jumps over the lazy dog.'.")
	} else {
		fmt.Println("The regex does NOT match the string 'The quick brown fox jumps over the lazy dog.'.")
	}

	//re1 := regexp.MustCompile(`[a-zA-Z+`) // This will panic if the regex is invalid, so use it only for hardcoded patterns that you are sure are correct.
	re1 := regexp.MustCompile(`[a-z]+`)
	fmt.Println("Regex compiled successfully using MustCompile")
	if re1.MatchString("The quick brown fox jumps over the lazy dog.") {
		fmt.Println("The regex matches the string 'The quick brown fox jumps over the lazy dog.'.")
	} else {
		fmt.Println("The regex does NOT match the string 'The quick brown fox jumps over the lazy dog.'.")
	}

	ProductIDs := "Product ID: P12345, Product ID: P67890, Product ID: A54321"
	re2 := regexp.MustCompile(`P\d+`) // This will match IDs starting with P and then one or more digits in a string

	// FindString returns the first match of the regex in the input string.
	// If there are multiple matches, it will only return the first one.
	matche := re2.FindString(ProductIDs)
	fmt.Println("First matching product ID:", matche)

	// FindAllString returns a slice of all matches in the input string. The second argument -1 means to find all matches.
	matches := re2.FindAllString(ProductIDs, -1)
	fmt.Println("Found product IDs:", matches)
}
