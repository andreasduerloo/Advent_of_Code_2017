package day_07

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/07.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	nodes, nodeList := parse(input)
	linkNodes(nodes)

	// Grab the first node and keep going up until one has no parents (or order list by len(parents?))
	current := nodeList[0]

	for current.parent != nil {
		current = current.parent
	}

	fmt.Println(current.name)

	calculateWeights(current)

	fmt.Println(current.subWeight)
	for _, c := range current.children {
		fmt.Println(c.subWeight)
	}

	return 0, 0
}
