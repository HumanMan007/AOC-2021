package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInputFromFile(file *os.File) int {

	scanner := bufio.NewScanner(file)
	instances := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")[1]
		elements := strings.Fields(line)

		for _, elem := range elements {
			switch len(elem) {
			case 2: //1
				instances++
			case 4: //4
				instances++
			case 3: //7
				instances++
			case 7: //8
				instances++
			}
		}
	}
	return instances
}

func main() {

	if len(os.Args) < 2 {
		log.Fatalf("input file")
	}
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(fmt.Errorf("opening instrution file:%v", err))
	}
	defer file.Close()

	// Gather input
	number := parseInputFromFile(file)

	fmt.Printf("Number of recognizable numbers %v\n", number)

}
