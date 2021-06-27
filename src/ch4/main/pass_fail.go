package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

}

func getFLoat() (float64, error) {

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return 0, err
	}

	input = strings.TrimSpace(input)

	float, err := strconv.ParseFloat(input, 64)

	if err != nil {
		return 0, err
	}
	return float, nil
}
