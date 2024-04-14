package main

import (
	"fmt"
	"math"
)

func main() {
	var h, a, ab, v float64

	fmt.Scan(&h, &a)

	ab = 3 * a * a * math.Sqrt(3) / 2
	v = ab * h / 3

	fmt.Print("O VOLUME DA PIRÂMIDE E =", v, " METROS CÚBICOS\n")

}
