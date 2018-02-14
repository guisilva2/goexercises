package main

import "fmt"

func main() {
	print := fmt.Println
	scan := fmt.Scan
	var n, m, divide, count int
	scan(&n, &m)

	v1 := make([]int, n)
	v2 := make([]int, m)

	for i := 0; i < n; i++ {
		scan(&v1[i])
	}

	for i := 0; i < m; i++ {
		scan(&v2[i])
	}

	for i := 0; i < n; i++ {
		for y := 0; y < m; y++ {
			if v2[y]%v1[i] == 0 {
				print("Divide")
				divide++
				if divide == n {
					count++
					divide = 0
				}
			}
		}
	}
	print(count)

}
