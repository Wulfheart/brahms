package viz

import "math"

func cartesian(r float64, angle float64) (x, y float64) {
	x = math.Cos(angle) * r
	y = math.Sin(angle) * r
	return x, y
}
