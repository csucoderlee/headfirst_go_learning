package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println(inRange(1, 10, 1, 2, 3, 4, 5, 11))
}

func maximum(numbers ...float64) float64 {

	max := math.Inf(-1)

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}

func inRange(min int, max int, numbers ...int) []int {

	var result []int

	for _, number := range numbers {
		if number > min && number < max {
			result = append(result, number)
		}
	}
	return result
}
