package day_04

import (
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/04.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	passPhrases := parse(input)

	first := mapReduce(passPhrases, valid)
	second := mapReduce(passPhrases, valid2)

	return first, second
}
