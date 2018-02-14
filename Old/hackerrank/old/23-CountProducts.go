package main

import "fmt"

func main() {
	var s, n, m int
	var z int = -1
	scan := fmt.Scan
	print := fmt.Print

	scan(&s, &n, &m)
	k := make([]int, n)
	u := make([]int, m)

	for i := 0; i < n; i++ {
		scan(&k[i])
	}
	for i := 0; i < m; i++ {
		scan(&u[i])
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			//S = MONEY $
			//Z = LIMIT ||
			//Se a soma dos produtos for maior ou igual a s, e maior que z
			if k[i]+u[j] <= s && k[i]+u[j] > z {
				z = k[i] + u[j]
			}
		}
	}
	print(z)
}
