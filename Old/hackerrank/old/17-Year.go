package main

import (
	"fmt"
	"time"
)

func main() {
	var year int
	scan := fmt.Scan
	print := fmt.Println

	print("Digite o ano:")
	scan(&year)

	dia := time.Now()
	year = dia.Day()

	print(year)
}
