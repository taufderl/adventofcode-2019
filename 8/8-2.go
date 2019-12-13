package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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

func parseInput(data []byte, x int, y int) ([][]int, [][]int) {
	nLayers := len(data) / (x * y)
	layers := make([][]int, nLayers)
	digitCounts := make([][]int, nLayers)
	for l := range layers {
		layers[l] = make([]int, x*y)
		digitCounts[l] = make([]int, 10)
	}

	for l := 0; l < nLayers; l++ {
		for i := 0; i < x*y; i++ {
			digit, _ := strconv.Atoi(string(data[l*x*y+i]))
			layers[l][i] = digit
			digitCounts[l][digit]++
		}
	}
	return layers, digitCounts
}

var x = 25
var y = 6

func solve(inputFilename string) {
	data := readInput(inputFilename)
	layers, digitCounts := parseInput(data, x, y)

	minimalZeroDigits := x * y
	minimalZeroDigitsLayer := 0
	for l := 0; l < len(layers); l++ {
		if digitCounts[l][0] < minimalZeroDigits {
			minimalZeroDigits = digitCounts[l][0]
			minimalZeroDigitsLayer = l
		}
	}
	fmt.Println("Layer ", minimalZeroDigitsLayer, "has the fewest zeros (", minimalZeroDigits, ").")
	fmt.Println("The checksum is", digitCounts[minimalZeroDigitsLayer][1]*digitCounts[minimalZeroDigitsLayer][2], ".")

	picture := make([]rune, x*y)
	for pixel := 0; pixel < x*y; pixel++ {
		for l := 0; l < len(layers); l++ {
			if layers[l][pixel] == 0 { // black
				picture[pixel] = ' '
				break
			} else if layers[l][pixel] == 1 { // white
				picture[pixel] = '█'
				break
			}
		}
		if pixel%x == 0 {
			fmt.Println()
		}
		fmt.Print(string(picture[pixel]))
	}
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
