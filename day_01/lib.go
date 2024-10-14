package day_01

import (
	"strconv"
	"strings"
)

func solveCaptcha(c string) int {
	var out int

	c = strings.TrimSpace(c)

	for i, r := range c {
		if i == (len(c) - 1) {
			if r == rune(c[0]) {
				out = addRune(out, r)
			}
		} else {
			if r == rune(c[i+1]) {
				out = addRune(out, r)
			}
		}
	}

	return out
}

func solveCaptcha2(c string) int {
	var out int

	c = strings.TrimSpace(c)

	for i, r := range c {
		halfwayAround := rune(c[(i+len(c)/2)%len(c)])

		if r == halfwayAround {
			out = addRune(out, r)
		}
	}

	return out
}

func addRune(sum int, r rune) int {
	int, err := strconv.Atoi(string(r))
	if err != nil {
		return sum
	} else {
		return sum + int
	}
}
