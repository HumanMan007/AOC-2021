package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseInputFromFile(file *os.File) ([][2]int, []func([][2]int) [][2]int, error) {

	scanner := bufio.NewScanner(file)
	ret := [][2]int{}
	foldList := []func([][2]int) [][2]int{}

	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 {
			break
		}

		coordinates := strings.Split(scanner.Text(), ",")
		if len(coordinates) != 2 {
			return nil, nil, fmt.Errorf("line \"%v\" is improperly formated")
		}

		x, err1 := strconv.Atoi(coordinates[0])
		y, err2 := strconv.Atoi(coordinates[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("improperly formated line \"%v\" : %v - %v", scanner.Text(), err1, err2)
		}

		ret = append(ret, [2]int{x, y})

	}

	for scanner.Scan() {
		ordinate := 0

		// Didn't quite feel like using atoi
		if _, err := fmt.Sscanf(scanner.Text(), "fold along x=%d", &ordinate); err == nil {
			foldList = append(foldList, func(dots [][2]int) [][2]int {
				fmt.Printf("FOLDING ALONG X = %v\n", ordinate)
				ret := make([][2]int, 0, len(dots))
				for _, coord := range dots {
					if coord[0] >= ordinate {
						coord = [2]int{2*ordinate - coord[0], coord[1]}
					}
					if !contains(ret, coord) {
						ret = append(ret, coord)
					}
				}
				return ret
			})
		} else if _, err := fmt.Sscanf(scanner.Text(), "fold along y=%d", &ordinate); err == nil {
			foldList = append(foldList, func(dots [][2]int) [][2]int {
				fmt.Printf("FOLDING ALONG Y = %v\n", ordinate)
				ret := make([][2]int, 0, len(dots))
				for _, coord := range dots {
					if coord[1] >= ordinate {
						coord = [2]int{coord[0], 2*ordinate - coord[1]}
					}
					if !contains(ret, coord) {
						ret = append(ret, coord)
					}
				}
				return ret
			})
		} else {
			return nil, nil, fmt.Errorf("improperly formated line \"%v\" : %v", scanner.Text(), err)
		}
	}

	return ret, foldList, nil
}

func contains(haystack [][2]int, needle [2]int) bool {
	for _, str := range haystack {
		if str == needle {
			return true
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

	dots, foldList, err := parseInputFromFile(file)
	if err != nil {
		log.Fatalf("error parsing file : %v", err)
	}

	fmt.Printf("%v - %v\n", dots, len(dots))

	dots = foldList[0](dots)

	fmt.Printf("%v - %v\n", dots, len(dots))

}
