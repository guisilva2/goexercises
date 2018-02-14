package main

import (
	"fmt"
)

func main() {
	var nTests, nPrisioners, nSweets, prisionerId int
	scan := fmt.Scan
	print := fmt.Println

	scan(&nTests)

	for i := 0; i < nTests; i++ {
		scan(&nPrisioners, &nSweets, &prisionerId)
		print(nSweets + prisionerId - 1)
	}

}
