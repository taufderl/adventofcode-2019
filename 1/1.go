package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
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
	values := bytes.Split(data, []byte("\n"))
	sum := 0
	for _, element := range values {
		intElement, _ := strconv.Atoi(string(element))
		res := solve(intElement)
		//fmt.Println(string(element), res)
		sum += res
	}
	fmt.Println(sum)
}
