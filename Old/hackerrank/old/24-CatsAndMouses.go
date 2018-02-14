package main

import (
	"fmt"
)

func abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

func main() {
	var q int
	scan := fmt.Scan
	print := fmt.Println
	scan(&q)

	x, y, z := make([]int, q), make([]int, q), make([]int, q)
	for i := 0; i < q; i++ {
		scan(&x[i], &y[i], &z[i])
	}
	for i := 0; i < q; i++ {
		if ((x[i]+y[i])/2) == z[i] && ((x[i]+y[i])%2) != 1 || x[i] == y[i] {
			print("Mouse C")
		} else if abs((z[i] - x[i])) < abs((z[i] - y[i])) {
			print("Cat A")
		} else if abs(z[i]-y[i]) < abs(z[i]-x[i]) {
			print("Cat B")
		}
	}
}
