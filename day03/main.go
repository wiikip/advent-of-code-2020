package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func slope(right int, down int) int {
	// open the file
	file, err := os.Open("../inputs/day03/file")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	nTrees := 0
	// read line by line
	next := 0
	down_next := 0
	for fileScanner.Scan() {
		if down_next == 0 {
			line := fileScanner.Text()

			if string(line[next]) == "#" {
				nTrees++
			}
			next = (next + right) % len(line)
		}
		down_next = (down_next + 1) % down


	}
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
	return nTrees
}

func main() {
	fmt.Println(slope(1, 1))
	fmt.Println(slope(3, 1))
	fmt.Println(slope(5, 1))
	fmt.Println(slope(7, 1))
	fmt.Println(slope(1, 2))
	fmt.Print(slope(1, 1) * slope(3, 1) * slope(5, 1) * slope(7, 1) * slope(1, 2))
}
