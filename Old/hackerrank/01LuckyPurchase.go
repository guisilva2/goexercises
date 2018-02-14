package main

import (
	"fmt"
	"strconv"
)

func verifyPrice(price int) int {
	var four, seven int
	priceChar := []rune(strconv.Itoa(price))
	for i := 0; i < len(priceChar); i++ {
		if priceChar[i] == '4' {
			four++
		} else if priceChar[i] == '7' {
			seven++
		} else {
			return 0
		}
	}
	if four == seven && four != 0 {
		return price
	}
	return 0
}

func main() {
	var n, bestPrice, price int
	var brand, bestBrand string
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&brand)
		fmt.Scan(&price)
		price = verifyPrice(price)
		if price > 0 && (price < bestPrice || bestPrice == 0) {
			bestBrand = brand
			bestPrice = price
		}
	}
	if bestBrand != "" {
		fmt.Println(bestBrand)
	} else {
		fmt.Println(-1)
	}

}
