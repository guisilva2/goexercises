package main

import (
	"fmt"
)

var (
	i int
	f float64
	b bool
	s string
)

func main() {
	fmt.Printf("As váriaveis iniciadas sem um valor inicial definido iniciam, sequencialmente como: \n INT ==> 0, Float==> 0, Bool==> true, String ==> '' ")
	fmt.Printf("%v %v %v %q\n", i, f, b, s)

}
