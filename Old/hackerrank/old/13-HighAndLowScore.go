package main

import "fmt"

func main() {
	print := fmt.Println
	scan := fmt.Scan
	var n, highScore, lowScore, highCount, lowCount int
	scan(&n)

	v1 := make([]int, n)

	for i := 0; i < n; i++ {
		scan(&v1[i])
		if i == 0 {
			highScore = v1[i]
			lowScore = v1[i]
		}

		if v1[i] > highScore {
			highCount++
			highScore = v1[i]
		}
		if v1[i] < lowScore {
			lowCount++
			lowScore = v1[i]
		}

	}

	print(highCount, lowCount)

}
