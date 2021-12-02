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

	depth, distance := 0, 0

	for scanner.Scan() {
		txt := scanner.Text()
		switch {
		case strings.Contains(txt, "up "):
			txt = strings.Replace(txt, "up ", "", -1)
			num, err := strconv.Atoi(txt)
			if err != nil {
				log.Default().Print(err)
				continue
			}
			depth -= num
		case strings.Contains(txt, "down "):
			txt = strings.Replace(txt, "down ", "", -1)
			num, err := strconv.Atoi(txt)
			if err != nil {
				log.Default().Print(err)
				continue
			}
			depth += num
		case strings.Contains(txt, "forward "):
			txt = strings.Replace(txt, "forward ", "", -1)
			num, err := strconv.Atoi(txt)
			if err != nil {
				log.Default().Print(err)
				continue
			}
			distance += num
		default:
			log.Default().Printf("reading line \"%v\": unrecognized command", txt)
		}
	}

	fmt.Printf("Distance %v, Depth %v, Value %v\n", distance, depth, distance*depth)
}
