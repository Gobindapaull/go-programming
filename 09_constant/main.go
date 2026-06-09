package main

import "fmt"

func main() {
	// Multiple constants
	const (
		appName = "Crypto Bot"
		maxUsers = 100000
	)
	
	fmt.Println(appName, maxUsers)
}