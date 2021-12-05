package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Board struct {
	Elements [5][5]int
	Checked  [5][5]bool
}

func parseInputFromFile(file *os.File) map[[2]int]int {
	scanner := bufio.NewScanner(file)

	incidence := make(map[[2]int]int)

	for scanner.Scan() {
		// Get line
		line := [2][2]int{}
		_, err := fmt.Fscanf(strings.NewReader(scanner.Text()), "%v,%v -> %v,%v", &line[0][0], &line[0][1], &line[1][0], &line[1][1])
		if err != nil {
			log.Default().Printf("non-valid line \"%v\": %v", line, err)
		}

		// If horizontal
		if line[0][0] == line[1][0] {
			x := line[0][0]
			from, to := 0, 0
			if line[0][1] > line[1][1] {
				from, to = line[1][1], line[0][1]+1
			} else {
				from, to = line[0][1], line[1][1]+1
			}

			for ; from != to; from++ {
				incidence[[2]int{x, from}]++
			}

			continue
		}
		// If veritcal
		if line[0][1] == line[1][1] {
			y := line[0][1]
			from, to := 0, 0
			if line[0][0] > line[1][0] {
				from, to = line[1][0], line[0][0]+1
			} else {
				from, to = line[0][0], line[1][0]+1
			}

			for ; from != to; from++ {
				incidence[[2]int{from, y}]++
			}
		}
	}

	return incidence
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
	incidenceMap := parseInputFromFile(file)

	// Calculate puntuation
	total := 0
	for _, in := range incidenceMap {
		if in > 1 {
			total++
		}
	}

	//fmt.Printf("Incidence map %v\n", incidenceMap)
	fmt.Printf("Punctuation %v\n", total)
}
