package day_06

import (
	"strconv"
	"strings"
)

func cycle(mem []int) int {
	var cycles int
	states := make(map[string]bool)

	// Loop de loop
	for !states[buildState(mem)] {
		states[buildState(mem)] = true
		mem = shuffle(mem)
		cycles++
	}

	return cycles
}

func shuffle(mem []int) []int {
	// Make a local copy?
	highest := indexMax(mem)
	highestVal := mem[highest]

	mem[highest] = 0

	for i := 1; i <= highestVal; i++ {
		mem[(highest+i)%len(mem)]++
	}

	return mem
}

func indexMax(mem []int) int {
	var highest int

	for i, val := range mem {
		if val > mem[highest] {
			highest = i
		}
	}

	return highest
}

func buildState(mem []int) string {
	strmem := make([]string, 0)

	for _, i := range mem {
		strmem = append(strmem, strconv.Itoa(i))
	}

	return strings.Join(strmem, " ")
}
