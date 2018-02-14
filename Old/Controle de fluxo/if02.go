package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	//Se o número for menor que o lim, será aŕesentado o mesmo, se não será apresentado o número da potência.
	if v := math.Pow(x, n); v < lim {
		return vgo:17:14: syntax error: unexpected newline, expecting comma or )
	}
	return lim
}

func main() {
	fmt.Println(pow(3, 2, 10), pow(3, 3, 20))
}
