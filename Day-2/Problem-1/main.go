package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var inputStrings []string
	checkSum := 0
	countFrequencyTwo := 0
	countFrequencyThree := 0

	fileScanner := bufio.NewScanner(file)

	for fileScanner.Scan() {
		inputStrings = append(inputStrings, fileScanner.Text())
	}

	//fmt.Println(inputStrings)

	data := make(map[string]map[string]int)

	for _, str := range inputStrings {

		data[string(str)] = make(map[string]int)
		for _, char := range str {
			freq := strings.Count(str, string(char))
			data[string(str)][string(char)] = freq
		}
	}

	//fmt.Print(data)

	for _, value1 := range data {

		var boolFrequencyTwo, boolFrequencyThree bool = false, false

		for _, value2 := range value1 {

			if value2 == 2 {
				boolFrequencyTwo = true
			}
			if value2 == 3 {
				boolFrequencyThree = true
			}

		}

		if boolFrequencyTwo {
			countFrequencyTwo++
		}

		if boolFrequencyThree {
			countFrequencyThree++
		}
	}

	checkSum = countFrequencyThree * countFrequencyTwo

	fmt.Printf("Answer = %d", checkSum)
}
