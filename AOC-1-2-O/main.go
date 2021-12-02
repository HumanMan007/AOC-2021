package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func windowSolve(path string) (int, error) {

	file, err := os.Open(path)
	if err != nil {
		return -1, fmt.Errorf("opening list file:%v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	prev := make([]int, 0, 3)
	total := 0

	for scanner.Scan() {
		if cur, err := strconv.Atoi(scanner.Text()); err != nil {
			log.Default().Println(err)
		} else {
			if len(prev) < cap(prev) {
				prev = append(prev, cur)
			} else {
				if cur > prev[0] {
					total++
				}
				prev[0] = prev[1]
				prev[1] = prev[2]
				prev[2] = cur
			}
		}
	}

	if len(prev) != cap(prev) {
		return -1, fmt.Errorf("insuficient numbers %v, minimum %v", len(prev), cap(prev))
	}

	return total, nil
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("input file")
		os.Exit(-1)
	}
	path := os.Args[1]

	sol, err := windowSolve(path)
	if err != nil {
		log.Fatalf("solving file %v: %v", path, err)
		os.Exit(1)
	}

	fmt.Printf("solution %v\n", sol)

}
