package main

import "fmt"

func main() {

	blocks := []string {
		"Block #1",
		"Block #2",
		"Block #3",
	}

	for i := 0; i < len(blocks); i++ {
		fmt.Println("Processing: ", blocks[i])
	}

}