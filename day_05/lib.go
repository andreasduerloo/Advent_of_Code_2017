package day_05

func execute(inst []int) int {
	max := len(inst)
	stepCount := 0
	i := 0

	for i < max && i >= 0 {
		next := i + inst[i]
		inst[i]++
		i = next
		stepCount++
	}

	return stepCount
}

func execute2(inst []int) int {
	max := len(inst)
	stepCount := 0
	i := 0

	for i < max && i >= 0 {
		next := i + inst[i]

		if inst[i] >= 3 {
			inst[i]--
		} else {
			inst[i]++
		}

		i = next
		stepCount++
	}

	return stepCount
}
