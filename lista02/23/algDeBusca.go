package main

import "fmt"

func main() {
	l := []int{5,3,4,2,1}
	var y int
	fmt.Scan(&y)
	fmt.Println(buscaSequencial(l, y))
}

func buscaSequencial(l []int, x int) int {
	for i := 0; i < len(l); i++ {
		if l[i] == x {
			return i
		}
	}
	return -1
}
