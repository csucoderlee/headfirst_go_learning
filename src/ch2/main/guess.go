package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println(" i got a random num between 1 and 100, but it is always ", target)
	rand.Seed(time.Now().Unix())
	target1 := rand.Intn(100) + 1;
	fmt.Println(" i got a random num between 1 and 100")
	fmt.Println(" the random num is ", target1)
}
