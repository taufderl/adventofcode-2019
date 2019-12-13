package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
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

func parseInput(data []byte) ([]orbitObject, []orbitRelation) {
	data2 := bytes.Split(data, []byte("\n"))
	var orbitObjects []orbitObject
	var orbitRelations []orbitRelation
	// create root object
	orbitObjects = append(orbitObjects, orbitObject{"COM", ""})
	for _, entry := range data2 {
		entry2 := bytes.Split(entry, []byte(")"))
		centerObjectName := string(entry2[0])
		orbitObjectName := string(entry2[1])
		orbitObjects = append(orbitObjects, orbitObject{orbitObjectName, centerObjectName})
		orbitRelations = append(orbitRelations, orbitRelation{centerObjectName, orbitObjectName})
	}
	return orbitObjects, orbitRelations
}

type orbitRelation struct {
	centerObject string
	orbitObject  string
}

type orbitObject struct {
	name   string
	parent string
}

func contains(slice []orbitObject, o orbitObject) bool {
	for _, elem := range slice {
		if o.name == elem.name {
			return true
		}
	}
	return false
}

func getOrbitObjectByName(orbitObjects []orbitObject, name string) orbitObject {
	for _, oo := range orbitObjects {
		if oo.name == name {
			return oo
		}
	}
	fmt.Println("Could not find orbit object with name ", name)
	os.Exit(1)
	return orbitObject{}
}

func getPathLength(orbitObjects []orbitObject, oo orbitObject) int {
	if oo.name == "COM" {
		return 0
	}
	oo = getOrbitObjectByName(orbitObjects, oo.parent)
	return getPathLength(orbitObjects, oo) + 1
}

func getPath(orbitObjects []orbitObject, oo orbitObject) []orbitObject {
	var path []orbitObject
	if oo.name == "COM" {
		return path
	}
	ooParent := getOrbitObjectByName(orbitObjects, oo.parent)
	return append(getPath(orbitObjects, ooParent), oo)
}

func solve(inputFilename string) {
	data := readInput(inputFilename)
	orbitObjects, _ := parseInput(data)
	debugPrint(orbitObjects)

	YouPath := getPath(orbitObjects, getOrbitObjectByName(orbitObjects, "YOU"))
	SanPath := getPath(orbitObjects, getOrbitObjectByName(orbitObjects, "SAN"))
	debugPrint(YouPath)
	debugPrint(SanPath)

	var disjunction []orbitObject
	for _, a := range YouPath {
		if !(contains(SanPath, a)) {
			disjunction = append(disjunction, a)
		}
	}
	for _, a := range SanPath {
		if !(contains(YouPath, a)) {
			disjunction = append(disjunction, a)
		}
	}

	fmt.Println(len(disjunction) - 2)
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
