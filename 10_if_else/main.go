package main

import "fmt"

func main() {
	price := 5

	if price >= 6 {
		fmt.Println("Sell")
	} else if price <= 4 {
		fmt.Println("Buy")
	} else {
		fmt.Println("Hold")
	}
}