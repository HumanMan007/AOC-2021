package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func solve(list []int) int {

	cont := 0

	for i := 1; i < len(list); i++ {
		if list[i] > list[i-1] {
			cont++
		}
	}

	return cont

}

func readList(path string) ([]int, error) {

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening list file:%v", err)
	}
	defer file.Close()

	list := []int{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if cur, err := strconv.Atoi(scanner.Text()); err != nil {
			log.Default().Println(err)
		} else {
			list = append(list, cur)
		}
	}

	if len(list) == 0 {
		return nil, fmt.Errorf("file does not contain numbers")
	}

	return list, nil
}

func main() {

	if len(os.Args) < 2 {
		log.Fatal("input file")
		os.Exit(1)
	}
	path := os.Args[1]

	arr, err := readList(path)
	if err != nil {
		log.Fatalf("loading list: %v", err)
		os.Exit(1)
	}

	fmt.Printf("%v \n- solution %v\n", arr, solve(arr))

}
