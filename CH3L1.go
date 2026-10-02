//conditionals
// if and else statment 

// Assignment

// Fix the bug on line 21. If messageLen is less than or equal to the maxMessageLen the program should print "Message sent", else it should print "Message not sent".



// code

package main

import "fmt"

func main() {
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length:", messageLen, "and a max length of:", maxMessageLen)

	if messageLen < maxMessageLen { // it was error here and it is now fixed.
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}
}
