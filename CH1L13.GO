// same line declaration 
// You can declare multiple variables on the same line:

// EX :  mileage, company := 80276, "Toyota"

// The above is the same as:

// mileage := 80276
// company := "Toyota"

//ASSIGNMENT
// At the top of the main function, declare a float called averageOpenRate and string called displayMessage on the same line.

// Initialize them to values:

// .23
// is the average open rate of your messages



package main

import "fmt"

func main() {

	averageOpenRate, displayMessage := .23, "is the average open rate of your messages"
	fmt.Println(averageOpenRate, displayMessage)
}

