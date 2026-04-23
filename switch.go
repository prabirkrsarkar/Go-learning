package main

import (
	"fmt"
	"time"
)

func main() {
	day := "Monday"
	switch day {

	case "Monday", "Tuesday":
		fmt.Println("Start of the week, lots of meetings")

	case "Sunday", "Saturday":
		fmt.Println("Weekend")

	default:
		fmt.Println("Midweek days")

	}

	switch now := time.Now().Hour(); { // switch on a dynamic value
	case now < 12:
		fmt.Println("Good morning!")
	case now < 18 && now > 12:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}

	checkType := func(i interface{}) {
		switch v := i.(type) {
		case int:
			fmt.Println("Integer:", v)
		case string:
			fmt.Println("String:", v)
		case bool:
			fmt.Println("Boolean:", v)
		default:
			fmt.Println("Unknown type")
		}
	}

	checkType(true)
}
