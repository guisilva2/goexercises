package main

import (
	"fmt"
)

func main() {
	var n, k, charged, sum int
	scan := fmt.Scan
	print := fmt.Println

	scan(&n, &k)
	c := make([]int, n)

	for i := 0; i < n; i++ {
		scan(&c[i])
		if i != k {
			sum += c[i]
		}
	}
	scan(&charged)
	dif := charged - (sum / 2)
	if dif == 0 || dif > (sum/2) {
		print("Bon Appetit")

	} else if dif < (sum / 2) {
		print(dif)
	}
}
