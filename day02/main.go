package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func policy1() {
	// open the file
	file, err := os.Open("../inputs/day02/file")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	nValid := 0
	// read line by line
	for fileScanner.Scan() {

		line := fileScanner.Text()

		splitedLine := strings.Split(line, ": ")
		policy, password := splitedLine[0], splitedLine[1]

		letterMin, letterMax, letter :=
			strings.Split(strings.Split(policy, " ")[0], "-")[0],
			strings.Split(strings.Split(policy, " ")[0], "-")[1],
			strings.Split(policy, " ")[1]

		nLetterMin, err := strconv.Atoi(letterMin)
		if err != nil {
			fmt.Print("Error converting string to int")
			return
		}
		nLetterMax, err := strconv.Atoi(letterMax)
		if err != nil {
			fmt.Print("Error converting string to int")
			return
		}
		if strings.Count(password, letter) >= nLetterMin && strings.Count(password, letter) <= nLetterMax {
			nValid = nValid + 1
		}

	}
	fmt.Printf("%v valid passwords", nValid)
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()
}

func policy2() {
	file, err := os.Open("../inputs/day02/file")

	//handle errors while opening
	if err != nil {
		log.Fatalf("Error when opening file: %s", err)
	}

	fileScanner := bufio.NewScanner(file)
	nValid := 0
	// read line by line
	for fileScanner.Scan() {

		line := fileScanner.Text()

		splitedLine := strings.Split(line, ": ")
		policy, password := splitedLine[0], splitedLine[1]

		letterMin, letterMax, letter :=
			strings.Split(strings.Split(policy, " ")[0], "-")[0],
			strings.Split(strings.Split(policy, " ")[0], "-")[1],
			strings.Split(policy, " ")[1]

		nLetterMin, err := strconv.Atoi(letterMin)
		if err != nil {
			fmt.Print("Error converting string to int")
			return
		}
		nLetterMax, err := strconv.Atoi(letterMax)
		if err != nil {
			fmt.Print("Error converting string to int")
			return
		}
		if nLetterMax > 0 && nLetterMax <= len(password) && nLetterMin > 0 && nLetterMin <= len(password) {
			if string(password[nLetterMax-1]) == letter && string(password[nLetterMin-1]) != letter ||
				string(password[nLetterMin-1]) == letter && string(password[nLetterMax-1]) != letter {
				nValid = nValid + 1
			}
		}

	}
	fmt.Printf("%v valid passwords", nValid)
	// handle first encountered error while reading
	if err := fileScanner.Err(); err != nil {
		log.Fatalf("Error while reading file: %s", err)
	}

	file.Close()

}

func main() {
	policy1()
	policy2()
}
