package main

import "fmt"

func main() {
	// #1 - Change the code to print "Starting Textio server...!" instead of "Hello, World!".
	fmt.Println("Starting Textio server...")
	// #2 Intial variables/types
	// initialize variables here
	var smsSendingLimit int
	var costPerSMS float64
	var hasPermission bool
	var username string

	fmt.Printf("%v %.2f %v %q\n", smsSendingLimit, costPerSMS, hasPermission, username)
}
