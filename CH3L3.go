// Switch

// Switch statements are a way to compare a value against multiple options. They are similar to if-else statements but are more concise and readable when the number of options is more than 2.


// assignment 

// Fix the bug in the billingCost function. The "basic" plan is set correctly, but we need matches for the "pro" and "enterprise" plans too. If the plan is:

// "pro", the cost should be 20.0
// "enterprise", the cost should be 50.0



// code



package main

import "fmt"

func billingCost(plan string) float64 { // just added cases and returned it.. fixed.
	switch plan {
	case "basic":
		return 10.0
	case "pro":
		return 20.0
	case "enterprise":
		return 50.0
	default:
		return 0.0
	}
}


func main() {
	plan := "basic"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "pro"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "enterprise"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "free"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
	plan = "unknown"
	fmt.Printf("The cost for a %s plan is $%.2f\n", plan, billingCost(plan))
}
