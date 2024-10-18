package day_08

import (
	"regexp"
	"strconv"
	"strings"
)

type instruction struct {
	register string
	action   string
	value    int
	condReg  string
	operator string
	compVal  int
}

func parse(input []byte) []instruction {
	lines := strings.Split(string(input), "\n")

	instr := make([]instruction, 0)

	for _, line := range lines {
		if line != "" {
			instr = append(instr, newInstruction(line))
		}
	}

	return instr
}

func newInstruction(s string) instruction {
	wordsRe := regexp.MustCompile(`[a-z]+`)
	words := wordsRe.FindAllString(s, -1)

	intRe := regexp.MustCompile(`-?[0-9]+`)
	intsS := intRe.FindAllString(s, -1) // these are still strings!

	ints := make([]int, 0)
	for _, s := range intsS {
		i, err := strconv.Atoi(s)
		if err == nil {
			ints = append(ints, i)
		}
	}

	opRe := regexp.MustCompile(`[=<>!]+`)
	op := opRe.FindAllString(s, -1)[0]

	return instruction{
		register: words[0],
		action:   words[1],
		value:    ints[0],
		condReg:  words[3],
		operator: op,
		compVal:  ints[1],
	}
}

func executeAll(instr []instruction) (int, int) {
	regs := make(map[string]int)

	var highestEver int

	for _, i := range instr {
		switch i.operator {
		case "==":
			if regs[i.condReg] == i.compVal {
				if i.action == "inc" {
					regs[i.register] += i.value
				} else {
					regs[i.register] -= i.value
				}
			}
		case "!=":
			if regs[i.condReg] != i.compVal {
				if i.action == "inc" {
					regs[i.register] += i.value
				} else {
					regs[i.register] -= i.value
				}
			}
		case ">=":
			if regs[i.condReg] >= i.compVal {
				if i.action == "inc" {
					regs[i.register] += i.value
				} else {
					regs[i.register] -= i.value
				}
			}
		case "<=":
			if regs[i.condReg] <= i.compVal {
				if i.action == "inc" {
					regs[i.register] += i.value
				} else {
					regs[i.register] -= i.value
				}
			}
		case ">":
			if regs[i.condReg] > i.compVal {
				if i.action == "inc" {
					regs[i.register] += i.value
				} else {
					regs[i.register] -= i.value
				}
			}
		case "<":
			if regs[i.condReg] < i.compVal {
				if i.action == "inc" {
					regs[i.register] += i.value
				} else {
					regs[i.register] -= i.value
				}
			}
		}
		if regs[i.register] > highestEver {
			highestEver = regs[i.register]
		}
	}

	var highest int

	for _, v := range regs {
		if v > highest {
			highest = v
		}
	}

	return highest, highestEver
}

/*
func execute(inst instruction, reg map[string]int) {
	//
}
*/
