package main

import (
	"fmt"
    "math"
	"github.com/vrowello/fibonacci/calc"
)

var (
	input     float64
	solution1 float64
	solution2 float64
)

func main() {
	outputs := make(chan float64, 2)

	fmt.Printf("Enter an Integer: ")
	fmt.Scanln(&input)

	go calc.Case1(input, outputs)
	go calc.Case2(input, outputs)
	solution1 = <-outputs
	solution2 = <-outputs

	if solution1 == math.Floor(solution1) { // checks for perfect square
		fmt.Println(input, "is a fibonacci number")
	} else if solution2 == math.Floor(solution2) { // checks for perfect square
		fmt.Println(input, "is a fibonacci number")
	} else {
		fmt.Println(input, "is not a fibonacci number")
	}
}
