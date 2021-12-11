package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInputFromFile(file *os.File) ([][]int, error) {

	scanner := bufio.NewScanner(file)
	ret := make([][]int, 0, 10)

	for scanner.Scan() {
		line := scanner.Bytes()
		numLine := make([]int, 0, 10)
		if len(line) != 10 {
			return nil, fmt.Errorf("line \"%v\" in input does not contain the correct number of characters", scanner.Text())
		}
		for _, ch := range line {
			n, err := strconv.Atoi(string(ch))
			if err != nil {
				return nil, fmt.Errorf("non-numeric value in numeric sequence: %v", err)
			}
			numLine = append(numLine, n)
		}

		ret = append(ret, numLine)
	}

	if len(ret) != cap(ret) {
		return nil, fmt.Errorf("file contains incorrect number of lines, expected %v got %v", cap(ret), len(ret))
	}

	return ret, nil
}

func passIteration(octoMap [][]int) ([][]int, int) {
	flashMap := make([][]bool, 10)
	for i := range flashMap {
		flashMap[i] = make([]bool, 10)
	}

	flashes := 0

	// Increase all cells by one
	for i := range octoMap {
		for j := range octoMap[i] {
			octoMap[i][j]++
		}
	}

	// Keep checking until no more flashs
mainloop:
	for {
		for i := range octoMap {
			for j := range octoMap[i] {
				// If hasn't exploded yet and can explode, explode
				if !flashMap[i][j] && octoMap[i][j] > 9 {
					flashMap[i][j] = true

					if i > 0 {
						octoMap[i-1][j]++
						if j > 0 {
							octoMap[i-1][j-1]++
						}
						if j < len(octoMap[i])-1 {
							octoMap[i-1][j+1]++
						}
					}
					if i < len(octoMap)-1 {
						octoMap[i+1][j]++
						if j > 0 {
							octoMap[i+1][j-1]++
						}
						if j < len(octoMap[i])-1 {
							octoMap[i+1][j+1]++
						}
					}
					if j > 0 {
						octoMap[i][j-1]++
					}
					if j < len(octoMap[i])-1 {
						octoMap[i][j+1]++
					}
					continue mainloop
				}
			}
		}

		break
	}

	// Set exploded to 0
	for i := range flashMap {
		for j := range flashMap[i] {
			if flashMap[i][j] {
				octoMap[i][j] = 0
				flashes++
			}
		}
	}

	return octoMap, flashes
}

func showMap(octoMap [][]int) {
	for _, line := range octoMap {
		for _, val := range line {
			if val == 0 {
				fmt.Printf("\033[35m%v\033[37m", val)
			} else {
				fmt.Print(val)
			}
		}
		fmt.Println()
	}
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

	octoMap, err := parseInputFromFile(file)
	if err != nil {
		log.Fatalf("error parsing file : %v", err)
	}

	showMap(octoMap)
	fmt.Println()

	i := 1
	for ; ; i++ {
		octoMap, flashs := passIteration(octoMap)

		showMap(octoMap)
		if flashs == 100 {
			break
		}
	}

	fmt.Printf("Last iteration - %v\n", i)
}
