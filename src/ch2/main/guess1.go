package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	rand.Seed(time.Now().Unix())
	randnum := rand.Intn(100) + 1

	fmt.Println(" i got one rand num")
	fmt.Println(" can you guess the rand num?")

	success := false
	for i := 0; i < 10; i++ {
		fmt.Println(" you have 10  chance guess, remain", 10 - i, " chance ")
		reader := bufio.NewReader(os.Stdin)

		input, err := reader.ReadString('\n')

		if err != nil {
			log.Fatal(" get the input fail")
		}

		input = strings.TrimSpace(input)

		inputnum, err := strconv.Atoi(input)

		if err != nil {
			log.Fatal(" input string convert fail")
		}

		if inputnum > randnum {
			fmt.Println(" you guess too high")
		} else if inputnum < randnum{
			fmt.Println(" you guess too low")
		} else {
			success = true
			fmt.Println(" you guess is right")
			break
		}
	}

	if !success {
		fmt.Println(" you didnot guess the right number, the randnum is ", randnum)
	}
}

