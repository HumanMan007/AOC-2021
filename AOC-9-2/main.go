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

func removeDuplicates(array [][2]int) [][2]int {
	keys := make(map[[2]int]bool)
	list := [][2]int{}

	for _, entry := range array {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func growBasin(x, y int, heightMap [][]int) [][2]int {

	if heightMap[x][y] == 9 {
		return nil
	}

	basinList := [][2]int{[2]int{x, y}}
	// Grow leftwards
	if x != 0 && heightMap[x][y] < heightMap[x-1][y] {
		basinList = append(basinList, growBasin(x-1, y, heightMap)...)
	}
	// Grow rightwards
	if x != len(heightMap)-1 && heightMap[x][y] < heightMap[x+1][y] {
		basinList = append(basinList, growBasin(x+1, y, heightMap)...)
	}
	// Grow uptwards
	if y != 0 && heightMap[x][y] < heightMap[x][y-1] {
		basinList = append(basinList, growBasin(x, y-1, heightMap)...)
	}
	// Grow downwards
	if y != len(heightMap[x])-1 && heightMap[x][y] < heightMap[x][y+1] {
		basinList = append(basinList, growBasin(x, y+1, heightMap)...)
	}

	return removeDuplicates(basinList)
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

	basinList := map[[2]int][][2]int{}
	for i, line := range heightMap {
		for j, num := range line {
			if ((i == 0) || (num < heightMap[i-1][j])) &&
				((i == len(heightMap)-1) || (num < heightMap[i+1][j])) &&
				((j == 0) || (num < line[j-1])) &&
				((j == len(line)-1) || (num < line[j+1])) {
				basinList[[2]int{i, j}] = growBasin(i, j, heightMap)
			}
		}
	}

	biggest := [3]int{0, 0, 0}
	for origin, basin := range basinList {
		fmt.Printf("Basin in %v - %v - size %v\n", origin, basin, len(basin))

		switch {
		case len(basin) > biggest[0]:
			biggest[2] = biggest[1]
			biggest[1] = biggest[0]
			biggest[0] = len(basin)
		case len(basin) > biggest[1]:
			biggest[2] = biggest[1]
			biggest[1] = len(basin)
		case len(basin) > biggest[2]:
			biggest[2] = len(basin)
		}
	}

	fmt.Printf("Biggest basins %v, multiplied %v\n", biggest, biggest[0]*biggest[1]*biggest[2])
}
