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

const (
	//POSITION Parameter mode
	POSITION = iota
	//IMMEDIATE Parameter mode
	IMMEDIATE = iota
	//RELATIVE Parameter mode
	RELATIVE = iota
)

func parseParamMode(value int) (int, int, int) {
	param1mode := value / 100 % 10
	param2mode := value / 1000 % 10
	param3mode := value / 10000 % 10
	return param1mode, param2mode, param3mode
}

func getParam(program []int, IP int, RelativeBase int, param int, paramMode int) *int {
	if paramMode == POSITION {
		return &(program[program[IP+param]])
	}
	if paramMode == IMMEDIATE {
		return &(program[IP+param])
	}
	if paramMode == RELATIVE {
		return &(program[RelativeBase+program[IP+param]])
	}
	fmt.Println("Unknown param mode!", paramMode)
	return nil
}

func getOneParameter(program []int, IP int, RelativeBase int) *int {
	param1mode, _, _ := parseParamMode(program[IP])
	return getParam(program, IP, RelativeBase, 1, param1mode)
}

func getTwoParameters(program []int, IP int, RelativeBase int) (*int, *int) {
	param1mode, param2mode, _ := parseParamMode(program[IP])
	return getParam(program, IP, RelativeBase, 1, param1mode), getParam(program, IP, RelativeBase, 2, param2mode)
}

func getThreeParameters(program []int, IP int, RelativeBase int) (*int, *int, *int) {
	param1mode, param2mode, param3mode := parseParamMode(program[IP])
	return getParam(program, IP, RelativeBase, 1, param1mode), getParam(program, IP, RelativeBase, 2, param2mode), getParam(program, IP, RelativeBase, 3, param3mode)

}

func runProgram(programInput []int, noun int, verb int) int {
	program := make([]int, len(programInput)+1000)
	copy(program, programInput)
	var IP, RelativeBase int
	IP = 0
	RelativeBase = 0
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
			param1, param2, param3 = getThreeParameters(program, IP, RelativeBase)
			*param3 = *param1 + *param2
			IP += 4
		case 2: // MULTIPLY
			param1, param2, param3 = getThreeParameters(program, IP, RelativeBase)
			*param3 = *param1 * *param2
			IP += 4
		case 3: // INPUT
			param1 = getOneParameter(program, IP, RelativeBase)
			*param1 = readInputStream()
			IP += 2
		case 4: // OUTPUT
			param1 = getOneParameter(program, IP, RelativeBase)
			writeOutputStream(*param1)
			IP += 2
		case 5: // JUMP IF TRUE
			param1, param2 = getTwoParameters(program, IP, RelativeBase)
			if *param1 == 0 {
				IP += 3
			} else {
				IP = *param2
			}
		case 6: // JUMP IF FALSE
			param1, param2 = getTwoParameters(program, IP, RelativeBase)
			if *param1 == 0 {
				IP = *param2
			} else {
				IP += 3
			}
		case 7: // LESS THAN
			param1, param2, param3 = getThreeParameters(program, IP, RelativeBase)
			if *param1 < *param2 {
				*param3 = 1
			} else {
				*param3 = 0
			}
			IP += 4
		case 8: // EQUALS
			param1, param2, param3 = getThreeParameters(program, IP, RelativeBase)
			if *param1 == *param2 {
				*param3 = 1
			} else {
				*param3 = 0
			}
			IP += 4
		case 9:
			param1 = getOneParameter(program, IP, RelativeBase)
			RelativeBase += *param1
			debugPrint("Set relative base to" + strconv.Itoa(RelativeBase))
			IP += 2
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
