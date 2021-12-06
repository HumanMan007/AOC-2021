package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInputFromFile(file *os.File) map[int]int {
	scanner := bufio.NewScanner(file)

	fishes := map[int]int{}

	for scanner.Scan() {
		fishList := strings.Split(scanner.Text(), ",")

		for _, s_time := range fishList {
			time, err := strconv.Atoi(s_time)
			if err != nil {
				log.Default().Printf("non-numerical element %v in numerical list: %v", s_time, err)
				continue
			}
			fishes[time]++
		}
	}

	return fishes
}

func passIteration(fishes map[int]int) map[int]int {

	forth := map[int]int{}

	for time, num := range fishes {
		if time == 0 {
			forth[6] += num
			forth[8] += num
		} else {
			forth[time-1] += num
		}
	}

	return forth
}

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("input file and iterations")
	}
	path := os.Args[1]
	iter, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatalf("seccond argument \"%v\"must be a number : %v", os.Args[2], err)
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(fmt.Errorf("opening instrution file:%v", err))
	}
	defer file.Close()

	// Gather input
	incidenceMap := parseInputFromFile(file)

	// Pass through iterations
	for i := 0; i < iter; i++ {
		fmt.Printf("iteration %v - %v\n", i, incidenceMap)
		incidenceMap = passIteration(incidenceMap)
	}

	tot := 0
	for _, num := range incidenceMap {
		tot += num
	}

	fmt.Printf("Final iteration %v\n", incidenceMap)
	fmt.Printf("Total fish population %v\n", tot)
}
