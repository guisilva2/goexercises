// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	tempo := time.Now().String()
	fmt.Scan(&tempo)
	tempo =
		p(tempo.String("13:02:03"))

}
