package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func Bracket(input string) (bool, rune) { // Returns whether the string is not corrupted and the illegal character if applicable

	buffer := make([]rune, 0, len(input))

	for _, c := range input {
		switch c {
		case '>':
			if buffer[len(buffer)-1] == '<' {
				buffer = buffer[:len(buffer)-1]
			} else {
				return false, c
			}
		case ']':
			if buffer[len(buffer)-1] == '[' {
				buffer = buffer[:len(buffer)-1]
			} else {
				return false, c
			}
		case ')':
			if buffer[len(buffer)-1] == '(' {
				buffer = buffer[:len(buffer)-1]
			} else {
				return false, c
			}
		case '}':
			if buffer[len(buffer)-1] == '{' {
				buffer = buffer[:len(buffer)-1]
			} else {
				return false, c
			}
		default:
			buffer = append(buffer, c)
		}
	}

	return true, 0
}

func parseInputFromFile(file *os.File) (ret []string) {

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		ret = append(ret, scanner.Text())
	}

	return ret
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
	valueMap := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	punctuation := 0
	for i, line := range lineList {
		fmt.Printf("%v - %v", i, line)
		if ok, ch := Bracket(line); !ok {
			fmt.Printf(" - Illegal character %v", string(ch))
			punctuation += valueMap[ch]
		}
		fmt.Println()
	}

	fmt.Printf("Total - %v\n", punctuation)
}
