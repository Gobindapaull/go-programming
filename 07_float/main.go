package main

import "fmt"

func main() {
	// Go infers them as float64
	score1 := 4.9
	score2 := 4.8

	averageScore := (score1 + score2) / 2
	fmt.Println(averageScore)
}