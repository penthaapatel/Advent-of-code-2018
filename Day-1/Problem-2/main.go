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

	var nums []int
	frequency := 0
	var m = make(map[int]bool)

	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, num)
	}

	for {
		for _, num := range nums {
			frequency += num

			if m[frequency] == true {
				fmt.Printf("Answer: %d", frequency)
				return
			}

			m[frequency] = true

		}
	}

}
