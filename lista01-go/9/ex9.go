package main

import "fmt"

func main() {
	var a, b, c, delta int

	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)

	delta = b*b - 4*a*c

	fmt.Print("Delta = ", delta)
}
