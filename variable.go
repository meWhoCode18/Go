package main

import "fmt"

func main() {
	// declaring variable using integer int
	var smsSendingLimit int
	smsSendingLimit = 1000
	fmt.Println("Your SMS sending limit is", smsSendingLimit)

	// declaring a variable using boolean (fixed type assignment)
	var smssendinglimit bool
	smssendinglimit = true
	fmt.Println("Your SMS is True", smssendinglimit)

	// declaring a variable using string (fixed type assignment and quote mismatch)
	var name string
	name = "Krish"
	fmt.Println("You are : ", name)

	// declaring a float value in go : float64 (fixed type assignment)
	var percentage float64
	percentage = 86.56
	fmt.Println("Your percentage is : ", percentage)
}
