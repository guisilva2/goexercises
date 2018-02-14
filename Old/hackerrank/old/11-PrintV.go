package main

import "fmt"

func main() {
	print := fmt.Println
	scan := fmt.Scan
	var n, m int
	scan(&n)
	scan(&m)
	v1 := make([]int, n)
	v2 := make([]int, m)

	for i := 0; i < n; i++ {
		scan(&v1[i])
	}
	for i := 0; i < m; i++ {
		scan(&v2[i])
	}
	print(v1)
	print(v2)

}
