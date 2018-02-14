package main

import "fmt"

func main() {
	var n, stepCount int
	stepCount = 0
	scan := fmt.Scan
	print := fmt.Println
	scan(&n)

	steps := make([]byte, n)

	for i := 0; i < n; i++ {
		scan(&steps[i])
		if steps[i] == 'D' {
			stepCount--
		} else if steps[i] == 'U' {
			stepCount++
		}
	}
	print(stepCount)
}
