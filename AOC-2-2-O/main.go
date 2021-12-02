package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("input file")
		os.Exit(1)
	}
	path := os.Args[1]

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(fmt.Errorf("opening instrution file:%v", err))
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	depth, distance, aim := 0, 0, 0

	command := map[string]func(num int){
		"up":      func(num int) { aim -= num },
		"down":    func(num int) { aim += num },
		"forward": func(num int) { distance += num; depth += num * aim },
	}

	for scanner.Scan() {
		txt := strings.Fields(scanner.Text())
		if _, ok := command[txt[0]]; !ok || len(txt) < 2 {
			log.Default().Printf("not valid command %v", txt)
			continue
		}
		num, err := strconv.Atoi(txt[1])
		if err != nil {
			log.Default().Print(err)
			continue
		}
		command[txt[0]](num)
	}

	fmt.Printf("Distance %v, Depth %v, Value %v\n", distance, depth, distance*depth)
}
