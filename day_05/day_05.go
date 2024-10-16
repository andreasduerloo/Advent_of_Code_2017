package day_05

import (
	"advent2017/helpers"
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/05.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	instructions := helpers.ReGetInts(string(input))

	first := execute(instructions)

	instructions = helpers.ReGetInts(string(input))

	second := execute2(instructions)

	return first, second
}
