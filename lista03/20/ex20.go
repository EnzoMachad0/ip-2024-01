package main

import (
	"fmt"
	"math"
)

func main() {
	var N int
	fmt.Scan(&N)

	var prevX, prevY, prevZ float64
	fmt.Scan(&prevX, &prevY, &prevZ)

	maxAbsCoord := 0.0

	for i := 1; i < N; i++ {
		var x, y, z float64
		fmt.Scan(&x, &y, &z)

		vx := math.Abs(x - prevX)
		vy := math.Abs(y - prevY)
		vz := math.Abs(z - prevZ)

		if vx > maxAbsCoord {
			maxAbsCoord = vx
		}
		if vy > maxAbsCoord {
			maxAbsCoord = vy
		}
		if vz > maxAbsCoord {
			maxAbsCoord = vz
		}

		prevX, prevY, prevZ = x, y, z
	}

	fmt.Printf("%.2f\n", maxAbsCoord)
}
