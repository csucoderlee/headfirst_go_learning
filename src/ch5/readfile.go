package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	exPath, err := os.Getwd()
	if err != nil {
		log.Fatal(" os getwd error")
	}

	//open不是在当前路径下查找, 而是项目的根目录
	fmt.Println(exPath)

	//file, err := os.Open(exPath + "/src/ch5/data.txt")
	file, err := os.Open("data.txt")

	if err != nil {
		log.Fatal(" open file error", err)
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err = file.Close()
	if err != nil {
		log.Fatal(" close file error")
	}

	if scanner.Err() != nil {
		log.Fatal(" scanner file error,", scanner.Err())
	}
}
