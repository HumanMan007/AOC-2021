package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	Elements [5][5]int
	Checked  [5][5]bool
}

func parseInputFromFile(file *os.File) ([]int, []*Board, error) {
	scanner := bufio.NewScanner(file)

	// Get Called numbers
	scanner.Scan()
	numbersTxt := strings.Split(scanner.Text(), ",")
	numberCall := make([]int, len(numbersTxt))
	for i, num := range numbersTxt {
		n, err := strconv.Atoi(num)
		if err != nil {
			return nil, nil, fmt.Errorf("non numeric value found in list %v : %v", num, err)
		}
		numberCall[i] = n
	}

	//Get boards
	boardList := []*Board{}

	for scanner.Scan() {
		if scanner.Text() != "" {
			return nil, nil, fmt.Errorf("expected empty line found %v", scanner.Text())
		}

		newBoard := &Board{}
		for i := range [5]byte{} {
			if scanner.Scan() {
				numbersTxt := strings.Fields(scanner.Text())
				if len(numbersTxt) != 5 {
					return nil, nil, fmt.Errorf("line expected to have 5 fileds has %v elements : %v", len(numbersTxt), numbersTxt)
				}
				for j, num := range numbersTxt {
					n, err := strconv.Atoi(num)
					if err != nil {
						return nil, nil, fmt.Errorf("non numeric value found in list %v : %v", num, err)
					}
					newBoard.Elements[i][j] = n
				}
			} else {
				newBoard = nil
				break
			}
		}

		if newBoard == nil {
			break
		}
		boardList = append(boardList, newBoard)
	}

	return numberCall, boardList, nil
}

func winning(board [5][5]bool) bool {

	// Check horizontals
	for _, line := range board {
		for j, val := range line {
			if !val {
				break
			} else if j == len(line)-1 {
				return true
			}
		}
	}

	// Check verticals
	for j := range board {
		for i := range board {
			if !board[j][i] {
				break
			} else if i == len(board)-1 {
				return true
			}
		}
	}

	return false

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
	numbers, boards, err := parseInputFromFile(file)
	if err != nil {
		log.Fatalf("error reading file %v : %v", path, err)
	}

	if len(numbers) == 0 {
		log.Fatalf("incomplete input %v", numbers)

	}
	if len(boards) == 0 {
		log.Fatalf("incomplete input %v", boards)
	}

	// Solve
	var winner *Board = nil

	val := 0
main:
	for _, val = range numbers {
		for _, board := range boards {

			for i, line := range board.Elements {

				for j, elem := range line {

					if elem == val {
						board.Checked[i][j] = true
						if winning(board.Checked) {
							winner = board
							break main
						}
					}
				}
			}
		}

	}

	if winner == nil {
		log.Fatalf("no winner found")
	}

	// Calculate puntuation
	total := 0

	for i, line := range winner.Elements {
		for j, elem := range line {
			if !winner.Checked[i][j] {
				total += elem
			}
		}
	}

	fmt.Printf("Winner %v\n", winner.Elements)
	fmt.Printf("Sum %v, Last %v, Punctuation %v\n", total, val, total*val)
}
