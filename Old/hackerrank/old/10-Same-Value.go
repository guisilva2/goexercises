package main

import "fmt"

func main() {
	var x1, v1, x2, v2 int
	var result string = "NO"
	print := fmt.Println
	scan := fmt.Scan

	scan(&x1, &v1, &x2, &v2)

	for i := 0; i < 10000; i++ {
		x1 += v1
		x2 += v2
		if x1 == x2 {
			result = "YES"
			break
		}
	}
	print(result)
}
