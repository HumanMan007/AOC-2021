package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func parseInputFromFile(file *os.File) (map[string][]string, error) {

	scanner := bufio.NewScanner(file)
	ret := map[string][]string{}

	for scanner.Scan() {
		connection := strings.Split(scanner.Text(), "-")
		if len(connection) != 2 {
			log.Default().Print("line \"%v\" is improperly formated")
		}

		ret[connection[0]] = append(ret[connection[0]], connection[1])
		ret[connection[1]] = append(ret[connection[1]], connection[0])
	}

	_, ok1 := ret["start"]
	_, ok2 := ret["end"]
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("input lacks connections from/to start or end nodes")
	}

	return ret, nil
}

func contains(haystack []string, needle string) bool {
	for _, str := range haystack {
		if str == needle {
			return true
		}
	}
	return false
}

func getPaths(path []string, connections map[string][]string) [][]string {

	options := [][]string{}

	for _, nxt := range connections[path[len(path)-1]] {
		if nxt == "end" {
			options = append(options, append(path, nxt))
		} else if (nxt[0] >= 'A' && nxt[0] <= 'Z') || !contains(path, nxt) {
			options = append(options, getPaths(append(path, nxt), connections)...)
		}
	}

	return options
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

	connections, err := parseInputFromFile(file)
	if err != nil {
		log.Fatalf("error parsing file : %v", err)
	}

	paths := getPaths([]string{"start"}, connections)

	for _, path := range paths {
		fmt.Println(path)
	}
}
