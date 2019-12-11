package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
)

var debug bool

func debugPrint(args interface{}) {
	if debug {
		fmt.Println(args)
	}
}

func readInput(inputFilename string) []byte {
	data, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	return data
}

type wire struct {
	x int
	y int
}

func createWire(input string) wire {
	value, _ := strconv.Atoi(input[1:len(input)])
	switch input[0] {
	case 76: // L
		return wire{-value, 0}
	case 82: // R
		return wire{value, 0}
	case 68: // D
		return wire{0, -value}
	case 85: // U
		return wire{0, value}
	}
	return wire{0, 0}
}

func parseInput(data []byte) ([]wire, []wire) {
	data2 := bytes.Split(data, []byte("\n"))
	dataA := bytes.Split(data2[0], []byte(","))
	dataB := bytes.Split(data2[1], []byte(","))
	wireA := make([]wire, len(dataA))
	wireB := make([]wire, len(dataB))
	for i := 0; i < len(dataA); i++ {
		wireA[i] = createWire(string(dataA[i]))
	}
	for i := 0; i < len(dataB); i++ {
		wireB[i] = createWire(string(dataB[i]))
	}
	return wireA, wireB
}

type byManhattan []wire

func (coordinates byManhattan) Len() int {
	return len(coordinates)
}
func (coordinates byManhattan) Swap(i, j int) {
	coordinates[i], coordinates[j] = coordinates[j], coordinates[i]
}
func (coordinates byManhattan) Less(i, j int) bool {
	return manhattan(coordinates[i]) < manhattan(coordinates[j])
}

func manhattan(w wire) int {
	return abs(w.x) + abs(w.y)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func contains(a []wire, x wire) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func solve(inputFilename string) {
	data := readInput(inputFilename)
	wireA, wireB := parseInput(data)
	var coordinatesA []wire
	var coordinatesB []wire
	pointer := wire{0, 0}
	var destination wire
	for _, w := range wireA {
		destination = wire{pointer.x + w.x, pointer.y + w.y}
		if w.x > 0 {
			for x := pointer.x; x < destination.x; x++ {
				coordinatesA = append(coordinatesA, wire{x, pointer.y})
			}
		}
		if w.x < 0 {
			for x := pointer.x; x > destination.x; x-- {
				coordinatesA = append(coordinatesA, wire{x, pointer.y})
			}
		}
		if w.y > 0 {
			for y := pointer.y; y < destination.y; y++ {
				coordinatesA = append(coordinatesA, wire{pointer.x, y})
			}
		}
		if w.y < 0 {
			for y := pointer.y; y > destination.y; y-- {
				coordinatesA = append(coordinatesA, wire{pointer.x, y})
			}
		}
		pointer = wire{pointer.x + w.x, pointer.y + w.y}
	}
	pointer = wire{0, 0}
	for _, w := range wireB {
		destination = wire{pointer.x + w.x, pointer.y + w.y}
		if w.x > 0 {
			for x := pointer.x; x < destination.x; x++ {
				coordinatesB = append(coordinatesB, wire{x, pointer.y})
			}
		}
		if w.x < 0 {
			for x := pointer.x; x > destination.x; x-- {
				coordinatesB = append(coordinatesB, wire{x, pointer.y})
			}
		}
		if w.y > 0 {
			for y := pointer.y; y < destination.y; y++ {
				coordinatesB = append(coordinatesB, wire{pointer.x, y})
			}
		}
		if w.y < 0 {
			for y := pointer.y; y > destination.y; y-- {
				coordinatesB = append(coordinatesB, wire{pointer.x, y})
			}
		}
		pointer = wire{pointer.x + w.x, pointer.y + w.y}
	}
	// remove leading {0,0} elements
	coordinatesA = coordinatesA[1:len(coordinatesA)]
	coordinatesB = coordinatesB[1:len(coordinatesB)]
	debugPrint(coordinatesA)
	debugPrint(coordinatesB)

	var intersect []wire
	for _, a := range coordinatesA {
		if contains(coordinatesB, a) {
			intersect = append(intersect, a)
		}
	}

	sort.Sort(byManhattan(intersect))
	debugPrint(intersect)
	fmt.Println(intersect[0], "distance: ", manhattan(intersect[0]))
}

func main() {
	debugPtr := flag.Bool("d", false, "Debug output")
	flag.Parse()
	args := flag.Args()
	inputFilename := args[0]
	debug = *debugPtr
	solve(inputFilename)
}
