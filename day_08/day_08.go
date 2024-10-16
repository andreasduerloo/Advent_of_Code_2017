package day_08

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/08.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	instructions := parse(input)

	first, second := executeAll(instructions)

	return first, second
}
