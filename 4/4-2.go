package main

import (
	"bytes"
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

func parseInput(data []byte) (password, password) {
	data2 := bytes.Split(data, []byte("-"))
	l, _ := strconv.Atoi(string(data2[0]))
	u, _ := strconv.Atoi(string(data2[1]))
	lower := password(l)
	upper := password(u)
	return lower, upper
}

type password int

func (p password) isSixDigits() bool {
	return len(strconv.Itoa(int(p))) == 6
}
func (p password) hasExactlyTwoAdjacentEqualDigits() bool {
	digits := strconv.Itoa(int(p))
	for i := 0; i < len(digits)-1; i++ {
		lneighbor := "X"[0]
		rneighbor := "Y"[0]
		if i > 0 {
			lneighbor = digits[i-1]
		}
		if i+2 < len(digits) {
			rneighbor = digits[i+2]
		}
		if digits[i] == digits[i+1] && lneighbor != digits[i] && rneighbor != digits[i] {
			return true
		}
	}
	return false
}
func (p password) digitsNeverDecrease() bool {
	digits := strconv.Itoa(int(p))
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
	}
	return true
}

func solve(inputFilename string) {
	data := readInput(inputFilename)
	lower, upper := parseInput(data)
	count := 0
	var p password
	for p = lower; p < upper+1; p++ {
		if p.isSixDigits() &&
			p.hasExactlyTwoAdjacentEqualDigits() &&
			p.digitsNeverDecrease() {
			count++
		}
	}
	fmt.Println(count)
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
