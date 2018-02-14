package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe    bool       = true
	MaxInt2 uint64     = 1<<64 - 1
	MaxInt  uintptr    = 1<<64 - 1
	z       complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value %v\n", MaxInt2, MaxInt2)
	fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value %v\n", z, z)
}
