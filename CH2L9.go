// Format Practice

// assignment 

// Create a userLog variable on line 15. It should contain:
// Name: FNAME LNAME, Age: AGE, Rate: MESSAGERATE, Is Subscribed: ISSUBSCRIBED, Message: MESSAGE
// Where FNAME LNAME AGE MESSAGERATE ISSUBSCRIBED and MESSAGE correspond to the variables above.
// MESSAGERATE should be rounded to the tenths place.


// CODE 



package main

import "fmt"

func main() {
	fname := "Dalinar"
	lname := "Kholin"
	age := 45
	messageRate := 0.5
	isSubscribed := false
	message := "Sometimes a hypocrite is nothing more than a man in the process of changing."



	userLog := fmt.Sprintf(
		"Name: %s %s, Age: %d, Rate: %.1f, Is Subscribed: %t, Message: %s",
		fname,
		lname,
		age,
		messageRate,
		isSubscribed,
		message,
	)



	fmt.Println(userLog)
}

