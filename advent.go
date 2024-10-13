package main

import (
	"advent2017/day_01"
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("No argument passed - exiting.")
		return
	}

	day, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println("Argument should be an integer - exiting.")
		return
	}

	solved := []func() (int, int){
		day_01.Solve,
	}
}
