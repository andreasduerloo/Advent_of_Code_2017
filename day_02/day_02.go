package day_02

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/02.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	lines := parse(input)

	first := sumInts(mapInts(lines, diff))
	second := sumInts(mapInts(lines, cleanDiv))

	return first, second
}
