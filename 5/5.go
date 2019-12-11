package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

var debug bool

func debugPrint(args interface{}) {
	if debug {
		fmt.Println(args)
	}
}

func readProgram(inputFilename string) []int {
	data, err := ioutil.ReadFile(inputFilename)
	if err != nil {
		fmt.Println("File reading error", err)
		return nil
	}
	program := bytes.Split(data, []byte(","))
	var intProgram []int
	intProgram = make([]int, len(program))
	for i := 0; i < len(program); i++ {
		intProgram[i], _ = strconv.Atoi(string(program[i]))
	}
	debugPrint(intProgram)
	return intProgram
}

var inputStream []byte
var outputStream []byte

func readInputStream() int {
	fmt.Print("Enter input: ")
	var input int
	_, err := fmt.Scanf("%d", &input)
	if err != nil {
		os.Exit(1111)
	}
	return input
}

func writeOutputStream(value int) {
	fmt.Println(value)
}

func parseParamMode(value int) (int, int, int) {
	param1mode := value / 100 % 10
	param2mode := value / 1000 % 10
	param3mode := value / 10000 % 10
	return param1mode, param2mode, param3mode
}

func runProgram(programInput []int, noun int, verb int) int {
	program := make([]int, len(programInput)+3)
	copy(program, programInput)
	IP := 0
	if noun > -1 {
		program[1] = noun
	}
	if verb > -1 {
		program[2] = verb
	}
	for {
		debugPrint(program[IP:len(program)])
		param1mode, param2mode, param3mode := parseParamMode(program[IP])
		var param1 *int
		var param2 *int
		var param3 *int
		instruction := program[IP] % 100
		switch instruction {
		case 1: //ADD
			if param1mode == 1 {
				param1 = &(program[IP+1])
			} else {
				param1 = &(program[program[IP+1]])
			}
			if param2mode == 1 {
				param2 = &(program[IP+2])
			} else {
				param2 = &(program[program[IP+2]])
			}
			if param3mode == 1 {
				param3 = &(program[IP+3])
			} else {
				param3 = &(program[program[IP+3]])
			}
			*param3 = *param1 + *param2
			IP += 4
		case 2: //MULTIPLY
			if param1mode == 1 {
				param1 = &(program[IP+1])
			} else {
				param1 = &(program[program[IP+1]])
			}
			if param2mode == 1 {
				param2 = &(program[IP+2])
			} else {
				param2 = &(program[program[IP+2]])
			}
			if param3mode == 1 {
				param3 = &(program[IP+3])
			} else {
				param3 = &(program[program[IP+3]])
			}
			*param3 = *param1 * *param2
			IP += 4
		case 3: //INPUT
			if param1mode == 1 {
				param1 = &(program[IP+1])
			} else {
				param1 = &(program[program[IP+1]])
			}
			*param1 = readInputStream()
			IP += 2
		case 4: //OUTPUT
			if param1mode == 1 {
				param1 = &(program[IP+1])
			} else {
				param1 = &(program[program[IP+1]])
			}
			writeOutputStream(*param1)
			IP += 2
		case 99: //HALT
			return program[0]
		default:
			fmt.Println("Failure!")
			return 0
		}
	}
}

func solve(inputFilename string) {
	program := readProgram(inputFilename)
	result := runProgram(program, -1, -1)
	fmt.Println(result)
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
