package main

import (
	"fmt"
)

func main() {
	var nTests, nStudents, conditionOnTime, nStudentsOnTime int
	var students []int
	scan := fmt.Scan
	print := fmt.Println

	scan(&nTests)

	for i := 0; i < nTests; i++ {

		scan(&nStudents, &conditionOnTime)
		students = make([]int, nStudents)
		nStudentsOnTime = 0
		for j := 0; j < nStudents; j++ {
			scan(&students[j])
			if students[j] <= 0 {
				nStudentsOnTime++
			}
		}
		if nStudentsOnTime >= conditionOnTime {
			print("NO")
		} else {
			print("YES")
		}
	}
}
