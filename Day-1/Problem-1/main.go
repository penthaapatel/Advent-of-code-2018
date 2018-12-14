package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	fileHandle, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)

	frequency := 0

	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())

		if err != nil {
			log.Fatal(err)
		}
		frequency += num
	}

	if err := fileScanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Answer = %d", frequency)
}
