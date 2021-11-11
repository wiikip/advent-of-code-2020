package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func sum_2() {
    // open the file
    file, err := os.Open("../inputs/day01/file")

    //handle errors while opening
    if err != nil {
        log.Fatalf("Error when opening file: %s", err)
    }

    fileScanner := bufio.NewScanner(file)

	numbers := make(map[int]bool)

    // read line by line
    for fileScanner.Scan() {
        num, err := strconv.Atoi(fileScanner.Text())

		if err != nil {
			fmt.Printf("Error converting string to int")
			return
		}
		if _, ok := numbers[num]; ok {

			fmt.Printf("Found: %v\n", num * (2020 - num ))
			return
		}
		numbers[2020-num] = true
    }
    // handle first encountered error while reading
    if err := fileScanner.Err(); err != nil {
        log.Fatalf("Error while reading file: %s", err)
    }

    file.Close()
}
func sum_3() {
	// open the file
	file, err := os.Open("../inputs/day01/file")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	// Create Array with all numbers
	var allNumbers []int
	numbers := make(map[int]int)
	fileScanner := bufio.NewScanner(file)

	// read line by line
	for fileScanner.Scan() {
		num, err := strconv.Atoi(fileScanner.Text())

		if err != nil {
			fmt.Printf("Error converting string to int")
			return
		}
		allNumbers = append(allNumbers, num)
	}
	for i:= 0; i < len(allNumbers)-1; i++ {
		for j:= i; j<len(allNumbers); j++ {
			if _, ok := numbers[allNumbers[i]]; ok {

				fmt.Printf("Found: %v \n", allNumbers[i]*numbers[allNumbers[i]])
				return
			}
			if _, ok := numbers[allNumbers[j]]; ok {

				fmt.Printf("Found: %v\n", allNumbers[j]*numbers[allNumbers[j]])
				return
			}
			numbers[2020 - allNumbers[i] - allNumbers[j]] = allNumbers[i] * allNumbers[j]
		}
	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}

func main(){
	sum_2()
	sum_3()
}
