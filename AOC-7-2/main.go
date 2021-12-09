package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseInputFromFile(file *os.File) map[int]int {
	scanner := bufio.NewScanner(file)

	crabs := map[int]int{}

	for scanner.Scan() {
		fishList := strings.Split(scanner.Text(), ",")

		for _, s_position := range fishList {
			position, err := strconv.Atoi(s_position)
			if err != nil {
				log.Default().Printf("non-numerical element %v in numerical list: %v", s_position, err)
				continue
			}
			crabs[position]++
		}
	}

	return crabs
}

func distanceToFuelUse(distance int) int {
	// Equivalent to the sum of all intergers in [0, distance]
	return distance * (distance + 1) / 2
}

func cumulativeFuelUse(crabs map[int]int, target int) int {

	totalFuel := 0

	for position, num := range crabs {
		if target-position > 0 {
			totalFuel += distanceToFuelUse(target-position) * num
		} else {
			totalFuel += distanceToFuelUse(position-target) * num
		}
	}

	return totalFuel
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
	crabMap := parseInputFromFile(file)

	from, to := math.MaxInt32, 0
	//Get min and max positions
	for pos := range crabMap {
		if pos < from {
			from = pos
		}
		if pos > to {
			to = pos
		}
	}

	// Pass through iterations
	min := math.MaxInt32
	minPos := 0
	for pos := from; pos <= to; pos++ {
		fuel := cumulativeFuelUse(crabMap, pos)
		if fuel < min {
			min = fuel
			minPos = pos
		}
	}

	fmt.Printf("Target %v - Fuel use %v\n", minPos, min)

}
