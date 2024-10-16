package day_03

func Solve() (int, int) {
	// INPUT GOES HERE
	input := 361527

	layers := buildLayers(input)

	first := distance(input, layers)
	second := spiral(input)
	return first, second
}
