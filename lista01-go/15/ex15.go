package main

import "fmt"

func main() {
	var n, y float64
	var x float64 = 2

	fmt.Scan(&n)

	for i := 1; i < int(n); i = i + 2 {
		y = x * x
		fmt.Println(x, "^2 = ", y)
		x = x + 2

	}

}
