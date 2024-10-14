package day_02

import (
	"advent2017/helpers"
	"strings"
)

func parse(input []byte) [][]int {
	out := make([][]int, 0)

	lines := strings.Split(string(input), "\n")
	for _, line := range lines {
		if line != "" {
			out = append(out, helpers.ReGetInts(line))
		}
	}

	return out
}

func diff(in []int) int {
	var min, max int

	min = in[0]
	max = in[0]

	for _, num := range in {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return max - min
}

func cleanDiv(in []int) int {
	for i := 0; i < len(in)-1; i++ {
		for j := (i + 1); j < len(in); j++ {
			if in[i]%in[j] == 0 {
				return in[i] / in[j]
			} else if in[j]%in[i] == 0 {
				return in[j] / in[i]
			}
		}
	}
	return 0
}

func mapInts(inslice [][]int, f func([]int) int) []int {
	out := make([]int, 0)

	for _, elem := range inslice {
		out = append(out, f(elem))
	}

	return out
}

func sumInts(inslice []int) int {
	var out int

	for _, elem := range inslice {
		out += elem
	}

	return out
}
