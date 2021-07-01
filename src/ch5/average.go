package main

import "fmt"

func main() {
	var nums = [...]float64{1,2,3}
	i := float64(len(nums))

	var sum float64 = 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum / i)
}
