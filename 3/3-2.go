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
var pathA []wire
var pathB []wire

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

func parseInput(data []byte) ([][]byte, [][]byte) {
	data2 := bytes.Split(data, []byte("\n"))
	dataA := bytes.Split(data2[0], []byte(","))
	dataB := bytes.Split(data2[1], []byte(","))
	return dataA, dataB
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

type byPathLength []wire

func (coordinates byPathLength) Len() int {
	return len(coordinates)
}
func (coordinates byPathLength) Swap(i, j int) {
	coordinates[i], coordinates[j] = coordinates[j], coordinates[i]
}
func (coordinates byPathLength) Less(i, j int) bool {
	return pathLength(coordinates[i]) < pathLength(coordinates[j])
}

func pathLength(w wire) int {
	return indexOf(pathA, w) + indexOf(pathB, w) + 2 // add two because the path's start with the first step at index 0
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
func indexOf(a []wire, x wire) int {
	for index, n := range a {
		if x == n {
			return index
		}
	}
	return -1
}

func getPath(inputPathData [][]byte) []wire {
	var coordinates []wire
	pointer := wire{0, 0}
	for _, data := range inputPathData {
		debugPrint(string(data))
		steps, _ := strconv.Atoi(string(data[1:len(data)]))
		switch data[0] {
		case 76: // L
			for x := pointer.x - 1; x > pointer.x-steps-1; x-- {
				coordinates = append(coordinates, wire{x, pointer.y})
			}
		case 82: // R
			for x := pointer.x + 1; x < pointer.x+steps+1; x++ {
				coordinates = append(coordinates, wire{x, pointer.y})
			}
		case 68: // D
			for y := pointer.y - 1; y > pointer.y-steps-1; y-- {
				coordinates = append(coordinates, wire{pointer.x, y})
			}
		case 85: // U
			for y := pointer.y + 1; y < pointer.y+steps+1; y++ {
				coordinates = append(coordinates, wire{pointer.x, y})
			}
		}
		pointer = coordinates[len(coordinates)-1]
	}
	return coordinates
}

func solve(inputFilename string) {
	data := readInput(inputFilename)
	wireA, wireB := parseInput(data)
	pathA = getPath(wireA)
	pathB = getPath(wireB)
	debugPrint(pathA)
	debugPrint(pathB)

	var intersect []wire
	for _, a := range pathA {
		if contains(pathB, a) {
			intersect = append(intersect, a)
		}
	}

	sort.Sort(byPathLength(intersect))
	debugPrint(intersect)
	fmt.Println(intersect[0], "distance: ", pathLength(intersect[0]))
}

func main() {
	debugPtr := flag.Bool("d", false, "Enable debug output")
	flag.Parse()
	args := flag.Args()
	inputFilename := args[0]
	debug = *debugPtr
	if debug {
		fmt.Println("Enabled debug output.")
	}
	solve(inputFilename)
}
