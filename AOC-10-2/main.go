package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func Bracket(input string) (bool, rune, []rune) {
	// Returns whether the string is not corrupted, the illegal character and the unclosed list of brackets if incomplete

	buffer := make([]rune, 0, len(input))

	for _, c := range input {
		switch c {
		case '>':
			if buffer[len(buffer)-1] == '<' {
				buffer = buffer[:len(buffer)-1]
			} else {
				return false, c, nil
			}
		case ']':
			if buffer[len(buffer)-1] == '[' {
				buffer = buffer[:len(buffer)-1]
			} else {
				return false, c, nil
			}
		case ')':
			if buffer[len(buffer)-1] == '(' {
				buffer = buffer[:len(buffer)-1]
			} else {
				return false, c, nil
			}
		case '}':
			if buffer[len(buffer)-1] == '{' {
				buffer = buffer[:len(buffer)-1]
			} else {
				return false, c, nil
			}
		default:
			buffer = append(buffer, c)
		}
	}

	return true, 0, buffer
}

func parseInputFromFile(file *os.File) (ret []string) {

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
}

func rating(list []rune) int {

	score := 0
	for i := len(list) - 1; i >= 0; i-- {
		score *= 5
		switch list[i] {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}
	return score
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

	lineList := parseInputFromFile(file)

	punctuations := make([]int, 0, len(lineList))
	for i, line := range lineList {
		fmt.Printf("%v - %v", i, line)
		if ok, ch, list := Bracket(line); !ok {
			fmt.Printf(" - Illegal character %v", string(ch))
		} else {
			if len(list) != 0 {
				punctuations = append(punctuations, rating(list))
				fmt.Printf(" - Unclosed brackets %v - %v", string(list), punctuations[len(punctuations)-1])
			}
		}
		fmt.Println()
	}

	sort.Ints(punctuations)

	fmt.Printf("Total - %v\n", punctuations)
	fmt.Printf("Winner - %v\n", punctuations[len(punctuations)/2])
}
