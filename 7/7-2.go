package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"sync"
	"time"
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

func readInputStream(InputStream []string, index int) int {
	input, _ := strconv.Atoi(InputStream[index])
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

func getOneParameter(program []int, IP int) *int {
	param1mode, _, _ := parseParamMode(program[IP])
	if param1mode == 1 {
		return &(program[IP+1])
	}
	return &(program[program[IP+1]])
}

func getTwoParameters(program []int, IP int) (*int, *int) {
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

func getThreeParameters(program []int, IP int) (*int, *int, *int) {
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

func runProgram(computer string, programCode []int, noun int, verb int, InputChannel chan int, OutputChannel chan int) []int {
	program := make([]int, len(programCode)+3)
	copy(program, programCode)
	IP := 0
	var OutputCache []int
	if noun > -1 {
		program[1] = noun
	}
	if verb > -1 {
		program[2] = verb
	}
	for {
		debugPrint(IP)
		var param1, param2, param3 *int
		instruction := program[IP] % 100
		switch instruction {
		case 1: // ADD
			param1, param2, param3 = getThreeParameters(program, IP)
			*param3 = *param1 + *param2
			IP += 4
		case 2: // MULTIPLY
			param1, param2, param3 = getThreeParameters(program, IP)
			*param3 = *param1 * *param2
			IP += 4
		case 3: // INPUT
			param1 = getOneParameter(program, IP)
			select {
			case *param1 = <-InputChannel:
				debugPrint(computer + " read value " + strconv.Itoa(*param1))
			case <-time.After(3 * time.Second):
				fmt.Println(computer, "received timeout reading from input channel")
			}
			IP += 2
		case 4: // OUTPUT
			param1 = getOneParameter(program, IP)
			OutputCache = append(OutputCache, *param1)
			OutputChannel <- *param1
			IP += 2
		case 5: // JUMP IF TRUE
			param1, param2 = getTwoParameters(program, IP)
			if *param1 == 0 {
				IP += 3
			} else {
				IP = *param2
			}
		case 6: // JUMP IF FALSE
			param1, param2 = getTwoParameters(program, IP)
			if *param1 == 0 {
				IP = *param2
			} else {
				IP += 3
			}
		case 7: // LESS THAN
			param1, param2, param3 = getThreeParameters(program, IP)
			if *param1 < *param2 {
				*param3 = 1
			} else {
				*param3 = 0
			}
			IP += 4
		case 8: // EQUALS
			param1, param2, param3 = getThreeParameters(program, IP)
			if *param1 == *param2 {
				*param3 = 1
			} else {
				*param3 = 0
			}
			IP += 4
		case 99: // HALT
			debugPrint(computer + " halted.")
			debugPrint(program[0])
			return OutputCache
		default: // UNKNOWN OPCODE
			fmt.Println("Got unknown OPCODE:", instruction)
			return OutputCache
		}
	}
}

// Permutation computes all permutations of a given int slice
func Permutation(slice []int, f func([]int)) {
	permutation(slice, f, 0)
}

func permutation(slice []int, f func([]int), i int) {
	if i > len(slice) {
		f(slice)
		return
	}
	permutation(slice, f, i+1)
	for j := i + 1; j < len(slice); j++ {
		slice[i], slice[j] = slice[j], slice[i]
		permutation(slice, f, i+1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func computeOutput(program []int, permutation []int) int {
	var lastOutput int
	AmpAIN := make(chan int, 2)
	AmpAB := make(chan int, 2)
	AmpBC := make(chan int, 2)
	AmpCD := make(chan int, 2)
	AmpDE := make(chan int, 2)
	AmpEOUT := make(chan int, 2)

	AmpAIN <- permutation[0]
	AmpAB <- permutation[1]
	AmpBC <- permutation[2]
	AmpCD <- permutation[3]
	AmpDE <- permutation[4]
	AmpAIN <- 0

	go runProgram("A", program, -1, -1, AmpAIN, AmpAB)  // A
	go runProgram("B", program, -1, -1, AmpAB, AmpBC)   // B
	go runProgram("C", program, -1, -1, AmpBC, AmpCD)   // C
	go runProgram("D", program, -1, -1, AmpCD, AmpDE)   // D
	go runProgram("E", program, -1, -1, AmpDE, AmpEOUT) // E

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			//time.Sleep(100 * time.Millisecond)
			select {
			case lastOutput = <-AmpEOUT:
				AmpAIN <- lastOutput
			case <-time.After(1 * time.Millisecond):
				wg.Done()
				return
			}
		}

	}()
	wg.Wait()

	return lastOutput
}

func solve(inputFilename string) {
	Program := readProgram(inputFilename)
	Phases := []int{5, 6, 7, 8, 9}
	max := 0
	Permutation(Phases, func(permutations []int) {
		thrusterSignal := computeOutput(Program, permutations)
		if max < thrusterSignal {
			max = thrusterSignal
		}
	})
	fmt.Println(max)
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
