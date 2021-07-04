package main

import "math"

func maximum(numbers ...float64) float64 {

	max := math.Inf(-1)

	for _, number := range numbers {
		if number > max {
			max = number
		}
	}

	return max
}
