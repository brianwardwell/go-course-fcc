package main

import "fmt"

func main() {

	// #1 - Change the code to print "Starting Textio server...!" instead of "Hello, World!".
	fmt.Println("Starting Textio server...")

	// #2 - Intialize variables/types
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)

	// #3 - Assign values to variables and print message
	messageStart := "Happy birthday! You are now"
	age := 21
	messageEnd := "years old!"

	fmt.Println(messageStart, age, messageEnd)

	fmt.Println("test_author")
}
