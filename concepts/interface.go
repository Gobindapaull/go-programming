package main
import "fmt"

type I interface {}

func PrintType(i I) {
	switch i.(type) {
	case int:
		fmt.Println("It's an integer")
	case string:
		fmt.Println("It's a string")
	default:
		fmt.Println("Other types")
	}
}

func main() {
	PrintType(5)
	PrintType("Hello")
	PrintType(1.5)
}
