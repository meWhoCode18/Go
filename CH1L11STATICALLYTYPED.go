// Go Is Statically Typed
// Go enforces static typing meaning variable types are known before the code runs. That means your editor and the compiler can display type errors before the code is ever run, making development easier and faster.


package main

import "fmt"

func main() {
	var username string = "presidentSkroob"
	var password string = "12345"

	// don't edit below this line
	fmt.Println("Authorization: Basic", username+":"+password)
}

