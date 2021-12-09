package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func decode(msg string, key map[int]string) int {
	for i, str := range key {
		if len(msg) == len(str) && len(remove(msg, str)) == 0 {
			return i
		}
	}
	return -1
}

func remove(source, afector string) string { // Remove all elements in affector from source
	fin := []byte{}

	for _, elem := range []byte(source) {
		if !strings.Contains(afector, string(elem)) {
			fin = append(fin, elem)
		}
	}
	return string(fin)
}

func intersect(source, afector string) string { // Remove all elements not in affector from source
	fin := []byte{}

	for _, elem := range []byte(source) {
		if strings.Contains(afector, string(elem)) {
			fin = append(fin, elem)
		}
	}
	return string(fin)
}

func generateDecoder(code []string) map[int]string {
	// This will be pretty dumb but I have no other clue on how to do it
	segments := map[int][]string{} // Equivalence between number of segments and number representations
	nToStr := map[int]string{}     // Equivalence between number and representation

	for _, str := range code {
		segments[len(str)] = append(segments[len(str)], str)
	}

	nToStr[1] = segments[2][0]
	nToStr[4] = segments[4][0]
	nToStr[7] = segments[3][0]
	nToStr[8] = segments[7][0]

	i := 0
	for i = range segments[5] {
		if len(intersect(segments[5][i], segments[5][(i+1)%3])) == len(intersect(segments[5][i], segments[5][(i+2)%3])) {
			nToStr[3] = segments[5][i]
			break
		}
	}
	segments[5][i] = segments[5][len(segments[5])-1]

	if len(remove(nToStr[4], segments[5][0])) == 2 {
		nToStr[2] = segments[5][0]
		nToStr[5] = segments[5][1]
	} else {
		nToStr[2] = segments[5][1]
		nToStr[5] = segments[5][0]
	}

	for i = range segments[6] {
		if len(remove(nToStr[5], segments[6][i])) == 1 {
			nToStr[0] = segments[6][i]
			break
		}
	}
	segments[6][i] = segments[6][len(segments[6])-1]

	if len(remove(segments[6][0], nToStr[1])) == 5 {
		nToStr[6] = segments[6][0]
		nToStr[9] = segments[6][1]
	} else {
		nToStr[6] = segments[6][1]
		nToStr[9] = segments[6][0]
	}

	return nToStr
}

func parseInputFromFile(file *os.File) ([]int, int) {

	scanner := bufio.NewScanner(file)
	decodedList := []int{}
	total := 0

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "|")
		if len(line) != 2 {
			log.Default().Printf("improperly formated line", scanner.Text())
			continue
		}
		decoder := strings.Fields(line[0])
		code := strings.Fields(line[1])

		if len(decoder) != 10 || len(code) != 4 {
			log.Default().Printf("erroneus number of digits in line %v - %v", decoder, code)
			continue
		}

		key := generateDecoder(decoder)
		value := 0
		multiplier := 1000

		for _, elem := range code {
			digit := decode(elem, key)
			value += multiplier * digit
			multiplier /= 10
		}
		decodedList = append(decodedList, value)
		total += value
	}
	return decodedList, total
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
	numberList, total := parseInputFromFile(file)

	fmt.Printf("Decoded numbers %v - totaling %v\n", numberList, total)

}
