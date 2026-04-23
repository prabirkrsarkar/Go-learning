package main

import "fmt"

func main() {
	var numbers [3]int
	numbers[0] = 1
	numbers[1] = 2

	for i, num := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", i, num) //note, Go assigns zero value to uninitialized elements
	}

	// short hand initialization
	moreNumbers := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(moreNumbers); i++ {
		fmt.Printf("Index: %d, Value: %d\n", i, moreNumbers[i])
	}

	var matrix [2][3]int // 2-D array with 2 rows and 3 columns
	matrix[0][0] = 1
	matrix[0][1] = 2
	matrix[0][2] = 3

	fmt.Println("%v\n", matrix)
}
