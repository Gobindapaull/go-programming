package main

import "fmt"

func main() {
	isOwner := true
	tradingStarted := false

	bothTrue := isOwner && !tradingStarted // true && true evaluates to true
	fmt.Println(bothTrue) // true

	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true) // false
}