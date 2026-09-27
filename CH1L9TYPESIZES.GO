// integers, uints, floats, and complex numbers all have type sizes.

// Signed integers (no decimal)
// int  int8  int16  int32  int64

// Unsigned integers (non-negative numbers/no decimal)
// uint uint8 uint16 uint32 uint64 uintptr

// Signed decimal numbers
// float32 float64

// Complex numbers (a complex number has a real and imaginary part)
// complex64 complex128



package main

import "fmt"

func main() {
	accountAgeFloat := 2.6
	accountAgeInt := int(accountAgeFloat)
	fmt.Println("Your account has existed for", accountAgeInt, "years")
}



