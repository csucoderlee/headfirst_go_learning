package main

import "fmt"

func main() {

	truth := true
	negate1(&truth)
	fmt.Println(truth)

	lies := false
	negate1(&lies)
	fmt.Println(lies)
}

func negate(myBoolean *bool) *bool {
	myBooleanType := !*myBoolean
	myBoolean = &myBooleanType
	return myBoolean
}

func negate1(myBoolean *bool) {
	*myBoolean = !*myBoolean
}

func negate2(myBoolean *bool) *bool {
	*myBoolean = !*myBoolean
	return myBoolean
}
