package day_04

import (
	"slices"
	"strings"
)

func parse(input []byte) [][]string {
	out := make([][]string, 0)

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if line != "" {
			out = append(out, strings.Split(line, " "))
		}
	}

	return out
}

func valid(p []string) bool {
	present := make(map[string]bool) // We need a set

	for _, w := range p {
		if !present[w] {
			present[w] = true
		} else {
			return false
		}
	}

	return true
}

func mapReduce(pps [][]string, f func([]string) bool) int {
	var out int

	for _, pp := range pps {
		if f(pp) {
			out++
		}
	}

	return out
}

func valid2(p []string) bool {
	present := make(map[string]bool) // We need a set

	for _, w := range p {
		w = alphabetize(w)
		if !present[w] {
			present[w] = true
		} else {
			return false
		}
	}

	return true
}

func alphabetize(p string) string {
	runeSlice := strings.Split(p, "")

	slices.Sort(runeSlice)

	return strings.Join(runeSlice, "")
}
