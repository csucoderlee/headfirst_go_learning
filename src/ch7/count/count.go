package main

import (
	"fmt"
	"github.com/csucoderlee/datafile"
	"log"
)

func main() {

	countWithSlice()

	countWithMap()
}

func countWithSlice() {
	lines, err := datafile.GetStrings("votes.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(lines))
	var names []string
	var counts []int
	for _, line := range lines {
		matched := false
		for i, name := range names {
			if name == line {
				counts[i]++
				matched = true
			}
		}

		if !matched {
			names = append(names, line)
			counts = append(counts, 1)
		}
	}

	for i, name := range names {
		fmt.Printf("%s:%d\n", name, counts[i])
	}
}

func countWithMap() {
	lines, err := datafile.GetStrings("votes.txt")

	if err != nil {
		log.Fatal("datafile GetStrings error", err)
	}

	names := make(map[string]int)
	for _, line := range lines {
		names[line]++
	}

	for s, i := range names {
		fmt.Printf("%s : %d\n", s, i)
	}
}
