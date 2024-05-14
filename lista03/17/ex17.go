package main

import "fmt"

func main() {
	var n, num int
	fmt.Scan(&n)

	count := make(map[int]int)

	for i := 0; i < n; i++ {
		fmt.Scan(&num)
		count[num]++
	}

	unicoCount := 0
	for _, v := range count {
		if v == 1 {
			unicoCount++
		}
	}

	fmt.Println(unicoCount)
}
