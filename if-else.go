package main

func main() {
	temp := 25

	if temp > 30 {
		println("It's hot outside.")
	} else if temp < 15 {
		println("It's cold outside.")
	} else {
		println("The weather is nice.")
	}

	score := 95

	if score >= 90 {
		println("Grade: A")
	} else if score >= 80 {
		println("Grade: B")
	} else if score >= 70 {
		println("Grade: C")
	} else if score >= 60 {
		println("Grade: D")
	} else {
		println("Grade: F")
	}

	userAccess := map[string]bool{
		"Prabir": true,
		"John":   false,
	}

	if hasAccess, exists := userAccess["Prabir"]; exists && hasAccess { // first part before the ; is initialization the rest is the condition
		println("Prabir has access.")
	} else if exists {
		println("Prabir does not have access.") // If exists is true here, but the first check failed, the only remaining logical possibility is that hasAccess must be false.
	} else {
		println("Prabir does not exist in the access list.")
	}
}
