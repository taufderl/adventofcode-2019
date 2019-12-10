package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func readInput() []byte {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	return data
}

func solve(value int) int {
	return value/3 - 2
}

func main() {
	data := readInput()
	program := bytes.Split(data, []byte(","))
	var intProgram []int
	intProgram = make([]int, len(program))
	for i := 0; i < len(program); i++ {
		intProgram[i], _ = strconv.Atoi(string(program[i]))
	}
	index := 0
	fmt.Println(intProgram)
	intProgram[1] = 12
	intProgram[2] = 2
	for {
		switch intProgram[index] {
		case 1: //ADD
			intProgram[intProgram[index+3]] = intProgram[intProgram[index+1]] + intProgram[intProgram[index+2]]
			index += 4
		case 2: //MULTIPLY
			intProgram[intProgram[index+3]] = intProgram[intProgram[index+1]] * intProgram[intProgram[index+2]]
			index += 4
		case 99: //HALT
			fmt.Println("Halting.", index)
			fmt.Println(intProgram[0])
			os.Exit(0)
		default:
			fmt.Println("Failure!")
			os.Exit(1)
		}
	}

}
