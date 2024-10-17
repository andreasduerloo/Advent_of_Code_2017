package day_06

import (
	"advent2017/helpers"
	"fmt"
	"os"
)

func Solve() (int, int) {
	input, err := os.ReadFile("./inputs/06.txt")
	if err != nil {
		fmt.Println("Could not read the input file - exiting")
		return 0, 0
	}

	memBanks := helpers.ReGetInts(string(input))

	first := cycle(memBanks)

	return first, 0
}
