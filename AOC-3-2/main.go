package main

import (
	"bufio"
	"container/list"
	"fmt"
	"log"
	"os"
)

func bool2int(arr []bool) int {
	ret, power := 0, 1

	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] {
			ret += power
		}
		power *= 2
	}
	return ret
}

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
	bitList_O2 := list.New()
	bitList_CO2 := list.New()

	for scanner.Scan() {
		line := scanner.Bytes()

		bitline := make([]bool, len(line))
		for i, c := range line {
			switch c {
			case '0':
				bitline[i] = false
			case '1':
				bitline[i] = true
			default:
				log.Default().Println("input contains non-binary element in line: " + string(line))
			}
		}

		bitList_O2.PushBack(bitline)
		bitList_CO2.PushBack(bitline)
	}

	for ruleIndex := 0; bitList_O2.Len() > 1 && ruleIndex < len(bitList_O2.Front().Value.([]bool)); ruleIndex++ {

		//Get bit rule
		ruleDelta := 0
		for elem := bitList_O2.Front(); elem != nil; elem = elem.Next() {
			line := elem.Value.([]bool)
			if line[ruleIndex] {
				ruleDelta++
			} else {
				ruleDelta--
			}
		}
		rule := (ruleDelta >= 0)

		// Find rulebreakers
		for elem := bitList_O2.Front(); elem != nil; {
			line := elem.Value.([]bool)
			t := elem.Next()
			if line[ruleIndex] != rule {
				bitList_O2.Remove(elem)
			}
			elem = t
		}

	}

	for ruleIndex := 0; bitList_CO2.Len() > 1 && ruleIndex < len(bitList_CO2.Front().Value.([]bool)); ruleIndex++ {

		//Get bit rule
		ruleDelta := 0
		for elem := bitList_CO2.Front(); elem != nil; elem = elem.Next() {
			line := elem.Value.([]bool)
			if line[ruleIndex] {
				ruleDelta++
			} else {
				ruleDelta--
			}
		}
		rule := (ruleDelta < 0)

		// Find rulebreakers
		for elem := bitList_CO2.Front(); elem != nil; {
			line := elem.Value.([]bool)
			t := elem.Next()
			if line[ruleIndex] != rule {
				bitList_CO2.Remove(elem)
			}
			elem = t
		}

	}

	O2 := bool2int(bitList_O2.Front().Value.([]bool))
	CO2 := bool2int(bitList_CO2.Front().Value.([]bool))

	fmt.Printf("Surviving elements O2 - %v -%v\n", bitList_O2.Front().Value.([]bool), O2)
	fmt.Printf("Surviving elements CO2%v - %v\n", bitList_CO2.Front().Value.([]bool), CO2)
	fmt.Println(O2 * CO2)
}
