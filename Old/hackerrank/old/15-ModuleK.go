package main

import (
	"fmt"
)

func main() {
	var n, k, count int
	scan := fmt.Scan
	print := fmt.Println

	scan(&n, &k)
	v1 := make([]int, n)

	for i := 0; i < n; i++ {
		scan(&v1[i])
	}

	for i := 0; i < n; i++ {
		for y := i; y < n; y++ {
			if y != i {
				if (v1[i]+v1[y])%k == 0 {
					count++
				}
			}
		}
	}
	print(count)
}
