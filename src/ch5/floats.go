package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	exPath, err := os.Getwd()
	if err != nil {
		log.Fatal(" os getwd error")
	}

	//open不是在当前路径下查找, 而是项目的根目录
	fmt.Println(exPath)

	var nums [3]float64
	nums, err = getFloats(exPath + "/src/ch5/data.txt")
	if err != nil {
		log.Fatal("getFloats error")
	}

	for k, v := range nums {
		fmt.Println(k, " : ", v)
	}
}

func getFloats(path string) ([3]float64, error) {
	var nums [3]float64
	file, err := os.Open(path)
	if err != nil {
		return nums, err
	}

	scanner := bufio.NewScanner(file)

	i := 0
	for scanner.Scan() {
		nums[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			return nums, err
		}
		i++
	}

	err = file.Close()

	if err != nil {
		return nums, err
	}

	if scanner.Err() != nil {
		return nums, scanner.Err()
	}

	return nums, nil
}
