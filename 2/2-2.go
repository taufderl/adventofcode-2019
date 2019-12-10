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

func runProgram(programInput []int, noun int, verb int) int {
	program := make([]int, len(programInput))
	copy(program, programInput)
	IP := 0
	program[1] = noun
	program[2] = verb
	//fmt.Println(program)
	for {
		switch program[IP] {
		case 1: //ADD
			program[program[IP+3]] = program[program[IP+1]] + program[program[IP+2]]
			IP += 4
		case 2: //MULTIPLY
			program[program[IP+3]] = program[program[IP+1]] * program[program[IP+2]]
			IP += 4
		case 99: //HALT
			return program[0]
		default:
			fmt.Println("Failure!")
			return 0
		}
	}
}

func main() {
	data := readInput()
	program := bytes.Split(data, []byte(","))
	var intProgram []int
	intProgram = make([]int, len(program))
	for i := 0; i < len(program); i++ {
		intProgram[i], _ = strconv.Atoi(string(program[i]))
	}
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			result := runProgram(intProgram, i, j)
			if result == 19690720 {
				fmt.Println(i, j)
				os.Exit(0)
			}
		}
	}
}
