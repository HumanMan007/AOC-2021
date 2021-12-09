package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseInputFromFile(file *os.File) [][]int {
	scanner := bufio.NewScanner(file)

	heightMap := [][]int{}

	for scanner.Scan() {
		line := make([]int, 0, len(scanner.Text()))

		for _, char := range scanner.Bytes() {
			num, _ := strconv.Atoi(string(char))

			line = append(line, num)
		}

		heightMap = append(heightMap, line)
	}

	return heightMap
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

	heightMap := parseInputFromFile(file)
	riskLevelSum := 0

	for i, line := range heightMap {
		for j, num := range line {
			if ((i == 0) || (num < heightMap[i-1][j])) &&
				((i == len(heightMap)-1) || (num < heightMap[i+1][j])) &&
				((j == 0) || (num < line[j-1])) &&
				((j == len(line)-1) || (num < line[j+1])) {
				riskLevelSum += 1 + num
			}
		}
	}

	fmt.Printf("Map\n%v\n-Total risk level %v\n", heightMap, riskLevelSum)

}
