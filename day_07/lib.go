package day_07

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	name           string
	weight         int
	childrenString []string
	parent         *node
	children       []*node
	subWeight      int
}

func parse(input []byte) (map[string]*node, []*node) {
	nodeMap := make(map[string]*node)
	nodeList := make([]*node, 0)

	lines := strings.Split(string(input), "\n")

	for _, line := range lines {
		if line != "" {
			newNode := nodeFromRe(line)

			nodeMap[newNode.name] = newNode
			nodeList = append(nodeList, newNode)
		}
	}

	return nodeMap, nodeList
}

func nodeFromRe(in string) *node {
	namesRe := regexp.MustCompile(`[a-z]+`)
	weightRe := regexp.MustCompile(`[0-9]+`)

	names := namesRe.FindAllString(in, -1)
	weight, err := strconv.Atoi(weightRe.FindAllString(in, -1)[0])
	if err != nil {
		fmt.Println("We have a non-int weight, abort!")
		panic("PANIC!")
	}

	return &node{
		name:           names[0],
		weight:         weight,
		childrenString: names[1:],
	}
}

func linkNodes(m map[string]*node) {
	for _, v := range m {
		if len(v.childrenString) != 0 {
			for _, c := range v.childrenString {
				v.children = append(v.children, m[c])
				m[c].parent = v
			}
		}
	}
}

// Function that recursively calculates weights - use memoization
func calculateWeights(root *node) int {
	subWeight := root.weight

	for _, c := range root.children {
		subWeight += calculateWeights(c)
	}

	root.subWeight = subWeight
	return subWeight
}
