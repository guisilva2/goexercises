package main

import (
	"fmt"
)

func main() {
	var n, c int
	scan := fmt.Scan
	print := fmt.Println
	scan(&n)

	v1 := make([]int, n)

	for i := 0; i < n; i++ {
		scan(&v1[i])
		c = v1[i]
		v1[i] = 1
		for j := 0; j < c; j++ {
			if v1[i]%2 == 0 {
				v1[i]++
			} else if v1[i]%2 == 1 {
				v1[i] *= 2
			}
		}
		print(v1[i])
	}
}
