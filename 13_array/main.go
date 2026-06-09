package main

import "fmt"

func main() {
	// 3 = fixed size
	// int = data type
	gasPrices := [3]int{5, 10, 20} 

	fmt.Println("Low: ", gasPrices[0], "Gwei")
	fmt.Println("Medium: ", gasPrices[1], "Gwei")
	fmt.Println("High: ", gasPrices[2], "Gwei")
}