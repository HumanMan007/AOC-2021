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

func react(template []byte, insertions map[[2]byte]byte) []byte {
	needle := 0
	reaction := make([]byte, 0, len(template)*2-1)

	for needle < len(template)-1 {
		reactors := [2]byte{template[needle], template[needle+1]}
		ins, ok := insertions[reactors]

		reaction = append(reaction, reactors[0])
		if ok {
			reaction = append(reaction, ins)
		}
		needle++
	}
	reaction = append(reaction, template[len(template)-1])

	return reaction
}

func punctuation(polymer []byte) int {
	counters := map[byte]int{}
	max, min := 0, math.MaxInt

	for _, char := range polymer {
		counters[char]++
	}

	for _, cont := range counters {
		if cont > max {
			max = cont
		}
		if cont < min {
			min = cont
		}
	}

	return max - min
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

	fmt.Printf("Template %v\n", string(template))

	for i := 1; i <= iter; i++ {
		template = react(template, insertions)
		//fmt.Printf("After step %v: %v\n", i, string(template))
	}

	fmt.Printf("Final punctuation %v\n", punctuation(template))
}
