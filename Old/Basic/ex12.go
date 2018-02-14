package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Exemplo de algumas conversões numéricas:")
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*y + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)

}
