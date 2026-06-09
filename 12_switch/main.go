package main

import "fmt"

func main() {
	chainID := 56

	switch chainID {
	case 1:
		fmt.Println("Ethereum Mainnet")
	case 56:
		fmt.Println("BNB Smart Chain")
	case 137:
		fmt.Println("Polygon")
	default:
		fmt.Println("Unknown Network")
	}
}