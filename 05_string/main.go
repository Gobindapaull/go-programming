package main

import (
	"fmt"
	"strings"
)


func main() {
	userName := "clinton Plank"

	fmt.Println(strings.ToUpper(userName)) // to convert the string to uppercase
	fmt.Println(strings.Contains(userName, "lank")) // true
}