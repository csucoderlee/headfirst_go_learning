package main

import (
	"fmt"
	"github.com/csucoderlee/datafile"
)

func main() {

	strings, err := datafile.GetStrings("data.txt")
	if err != nil {
		fmt.Println("datafile GetStrings error", err)
	}

	for _, s := range strings {
		fmt.Println(s)
	}

}
