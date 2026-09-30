package main

import "fmt"

func main() {
	const name = "Krish Panchal"
	const openRate = 10.523

	msg := fmt.Sprintf("Hi %s , your open rate is %.1f percent\n", name, openRate)

	fmt.Print(msg)
}
