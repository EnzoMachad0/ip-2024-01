package main

import (
	"fmt"
)

var n int

func main() {

	arr := make([]int, 0)
	for {
		_, err := fmt.Scan(&n)
		if err != nil {
			break
		}

		fmt.Scanf("%d", &n)

		arr = append(arr, n)

	}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("%b \n", arr[i])

	}

}
