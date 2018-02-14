package main

import "fmt"

func main() {
	var n, p, pages int
	scan := fmt.Scan
	print := fmt.Println
	scan(&n)
	scan(&p)

	if p <= n/2 {
		for i := 1; i < p; i += 2 {
			pages++
		}
	} else {
		if p%2 == 0 {
			for i := n; i > p+1; i -= 2 {
				pages++
			}
		} else {
			for i := n; i > p; i -= 2 {
				pages++
			}
		}

	}
	if n%2 == 1 && n-p == 1 {
		pages = 0
	}
	print(pages)
}
