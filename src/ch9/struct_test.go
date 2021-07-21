package main

import (
	"fmt"
	"testing"
)

func TestStruct(t *testing.T) {
	type gallons float64
}

type Number int

func (n *Number) Double() {
	*n *= 2
}

func main() {
	n := Number(2)
	fmt.Println("origin value is ", n)
	n.Double()
	fmt.Println("the value is", n)
}
