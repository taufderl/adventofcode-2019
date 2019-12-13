package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
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

var root *orbitObject

// func addOrbit(centerObjectName string, orbitObjectName string) {
// 	o := &orbitTree{name: centerObjectName, orbit: []*orbitTree{}}

// 	if centerObjectName == "COM" {
// 		root = o
// 	} else {
// 		centerObject, _ := objectTable[objectName]
// 		parent.children = append(parent.children, node)
// 	}

// 	nodeTable[id] = node
// }

func getOrbitObjectByName(orbitObjects []orbitObject, name string) orbitObject {
	for _, oo := range orbitObjects {
		if oo.name == name {
			return oo
		}
	}
	return orbitObject{}
}

func getPathLength(orbitObjects []orbitObject, oo orbitObject) int {
	if oo.name == "COM" {
		return 0
	}
	oo = getOrbitObjectByName(orbitObjects, oo.parent)
	return getPathLength(orbitObjects, oo) + 1
}

func solve(inputFilename string) {
	data := readInput(inputFilename)
	orbitObjects, _ := parseInput(data)
	sum := 0
	for _, oo := range orbitObjects {
		sum += getPathLength(orbitObjects, oo)
		debugPrint(oo.name + " " + oo.parent)
	}
	fmt.Println(sum)
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
