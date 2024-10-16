package day_03

import "fmt"

func distance(i int, layers []*layer) int {
	l := findLayer(i, layers)
	axesDistance := distanceToAxis(i, findAxes(l))

	return l.distance + axesDistance
}

type layer struct {
	start    int
	end      int
	distance int
}

func newLayer(start, end, dist int) *layer {
	s := start
	e := end
	d := dist

	return &layer{
		start:    s,
		end:      e,
		distance: d,
	}
}

func buildLayers(in int) []*layer {
	out := make([]*layer, 0)

	var low, high, dist int

	// Add the access port
	out = append(out, newLayer(1, 1, dist))

	// The data for the first layer away from the port
	low = 2
	high = 9
	dist = 1

	out = append(out, newLayer(low, high, dist))

	for high < in { // As long as we don't have a layer containing the input point
		dist++
		low = high + 1
		high = high + (dist * 8)

		out = append(out, newLayer(low, high, dist))
	}

	return out
}

func findLayer(in int, layers []*layer) *layer {
	for _, l := range layers {
		if l.start <= in && l.end >= in {
			return l
		}
	}

	return nil
}

// Each layer contains four points that are 'dist' away from the access port (straight lines from the origin).
// Find the nearest one - the sum of the steps to that point and dist is the distance of the input point.
func findAxes(l *layer) [4]int {
	out := [4]int{0, 0, 0, 0}

	if l == nil {
		fmt.Println("Uh oh!")
		return out
	}

	// The right axis is (dist - 1) above 'low' for any layer
	out[0] = l.start + (l.distance - 1)
	out[1] = out[0] + l.distance*2
	out[2] = out[1] + l.distance*2
	out[3] = out[2] + l.distance*2

	return out
}

func distanceToAxis(in int, axes [4]int) int {
	dist := absDiff(in, axes[0])

	for i := 1; i < 4; i++ {
		if absDiff(in, axes[i]) < dist {
			dist = absDiff(in, axes[i])
		}
	}

	return dist
}

func absDiff(i, j int) int {
	if i >= j {
		return i - j
	} else {
		return j - i
	}
}
