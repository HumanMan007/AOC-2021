package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("input file")
		os.Exit(1)
	}
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(fmt.Errorf("opening instrution file:%v", err))
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	maxCount := []int{}

	for scanner.Scan() {
		line := scanner.Bytes()

		// If encounter line with more 'bits' inrease size of array
		if inc := len(line) - len(maxCount); inc > 0 {
			t := make([]int, inc)
			maxCount = append(maxCount, t...)
		}

		for i, c := range line {
			switch c {
			case '0':
				maxCount[i]--
			case '1':
				maxCount[i]++
			default:
				log.Default().Println("input contains non-binary element in line: " + string(line))
			}
		}

	}

	gamma, epsilon := 0, 0
	power := 1

	for i := len(maxCount) - 1; i >= 0; i-- {
		// In case of equal number I asume the majority of 0 by default
		if maxCount[i] > 0 {
			gamma += power
		} else {
			epsilon += power
		}
		power *= 2
	}

	fmt.Printf("Gamma = %v, Epsilon = %v, Solution = %v\n", gamma, epsilon, gamma*epsilon)
}
