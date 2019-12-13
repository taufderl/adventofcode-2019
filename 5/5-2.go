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

func getOneParamter(program []int, IP int) *int {
	param1mode, _, _ := parseParamMode(program[IP])
	if param1mode == 1 {
		return &(program[IP+1])
	}
	return &(program[program[IP+1]])
}

func getTwoParamters(program []int, IP int) (*int, *int) {
	param1mode, param2mode, _ := parseParamMode(program[IP])
	var param1 *int
	var param2 *int
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
	return param1, param2
}

func getThreeParamters(program []int, IP int) (*int, *int, *int) {
	param1mode, param2mode, param3mode := parseParamMode(program[IP])
	var param1 *int
	var param2 *int
	var param3 *int
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
	return param1, param2, param3
}

func runProgram(programInput []int, noun int, verb int) int {
	program := make([]int, len(programInput)+3)
	copy(program, programInput)
	var IP int
	IP = 0
	if noun > -1 {
		program[1] = noun
	}
	if verb > -1 {
		program[2] = verb
	}
	for {
		debugPrint(IP)
		debugPrint(program[IP:len(program)])
		var param1, param2, param3 *int
		instruction := program[IP] % 100
		switch instruction {
		case 1: // ADD
			param1, param2, param3 = getThreeParamters(program, IP)
			*param3 = *param1 + *param2
			IP += 4
		case 2: // MULTIPLY
			param1, param2, param3 = getThreeParamters(program, IP)
			*param3 = *param1 * *param2
			IP += 4
		case 3: // INPUT
			param1 = getOneParamter(program, IP)
			*param1 = readInputStream()
			IP += 2
		case 4: // OUTPUT
			param1 = getOneParamter(program, IP)
			writeOutputStream(*param1)
			IP += 2
		case 5: // JUMP IF TRUE
			param1, param2 = getTwoParamters(program, IP)
			if *param1 == 0 {
				IP += 3
			} else {
				IP = *param2
			}
		case 6: // JUMP IF FALSE
			param1, param2 = getTwoParamters(program, IP)
			if *param1 == 0 {
				IP = *param2
			} else {
				IP += 3
			}
		case 7: // LESS THAN
			param1, param2, param3 = getThreeParamters(program, IP)
			if *param1 < *param2 {
				*param3 = 1
			} else {
				*param3 = 0
			}
			IP += 4
		case 8: // EQUALS
			param1, param2, param3 = getThreeParamters(program, IP)
			if *param1 == *param2 {
				*param3 = 1
			} else {
				*param3 = 0
			}
			IP += 4
		case 99: // HALT
			debugPrint(program[0])
			return 0
		default: // UNKNOWN OPCODE
			fmt.Println("Got unknown OPCODE:", instruction)
			return instruction
		}
	}
}

func solve(inputFilename string) {
	program := readProgram(inputFilename)
	result := runProgram(program, -1, -1)
	if result != 0 || debug {
		fmt.Println("Program exited with code", result)
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
