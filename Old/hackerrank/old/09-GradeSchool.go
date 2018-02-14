// Go supports time formatting and parsing via
// pattern-based layouts.

package main

import "fmt"

func main() {
	var nEstudantes, diferenca int
	print := fmt.Println
	scan := fmt.Scan
	//print("Digite o número de estudantes:")
	scan(&nEstudantes)

	gradeEscolar := make([]int, nEstudantes)

	for i := 0; i < nEstudantes; i++ {
		scan(&gradeEscolar[i])
		if gradeEscolar[i] >= 38 {
			diferenca = gradeEscolar[i]
			gradeEscolar[i] = gradeEscolar[i] / 5
			gradeEscolar[i] = (gradeEscolar[i] + 1) * 5
			diferenca = gradeEscolar[i] - diferenca
			if diferenca >= 3 {
				gradeEscolar[i] -= diferenca
			}
		}
		print(gradeEscolar[i])
	}
}
