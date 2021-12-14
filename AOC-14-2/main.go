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

func parseInputFromFile(file *os.File) ([]byte, map[[2]byte]byte, error) {

	scanner := bufio.NewScanner(file)
	insertions := map[[2]byte]byte{}

	scanner.Scan()
	template := scanner.Bytes()
	scanner.Scan()

	for scanner.Scan() {
		elems := strings.Split(scanner.Text(), " -> ")
		if len(elems) != 2 || len(elems[0]) != 2 || len(elems[1]) != 1 {
			return nil, nil, fmt.Errorf("improperly formated string \"%v\"", scanner.Text())
		}

		insertions[[2]byte{elems[0][0], elems[0][1]}] = elems[1][0]
	}

	return template, insertions, nil
}

func sliceToPairMap(slice []byte) map[[2]byte]int {
	pairIncidence := map[[2]byte]int{}

	for i := 0; i+1 < len(slice); i++ {
		pairIncidence[[2]byte{slice[i], slice[i+1]}]++
	}

	return pairIncidence
}

func react(chemicalPairs map[[2]byte]int, insertions map[[2]byte]byte) map[[2]byte]int {

	product := make(map[[2]byte]int)

	for reactors, incidence := range chemicalPairs {
		ins, reacts := insertions[reactors]

		if reacts {
			product[[2]byte{reactors[0], ins}] += incidence
			product[[2]byte{ins, reactors[1]}] += incidence
		} else {
			product[reactors] += incidence
		}
	}
	return product
}

func count(pairs map[[2]byte]int, start, end byte) map[byte]int {
	charIncidence := map[byte]int{start: 1}
	charIncidence[end]++

	for chars, incidence := range pairs {
		charIncidence[chars[0]] += incidence
		charIncidence[chars[1]] += incidence
	}

	for char := range charIncidence {
		charIncidence[char] /= 2
	}

	return charIncidence
}

func main() {

	if len(os.Args) < 3 {
		log.Fatalf("input file and iterations")
	}
	path := os.Args[1]

	iter, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal(fmt.Errorf("second argument must be interger:%v", err))
	}

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(fmt.Errorf("opening instrution file:%v", err))
	}
	defer file.Close()

	template, insertions, err := parseInputFromFile(file)
	if err != nil {
		log.Fatalf("error parsing file : %v", err)
	}

	pairs := sliceToPairMap(template)

	for i := 1; i <= iter; i++ {
		pairs = react(pairs, insertions)
	}

	max, min := 0, math.MaxInt

	for _, cont := range count(pairs, template[0], template[len(template)-1]) {
		if cont > max {
			max = cont
		}
		if cont < min {
			min = cont
		}
	}

	fmt.Printf("Final punctuation %v\n", max-min)
}
