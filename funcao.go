package main

import "fmt"

func media(n1, n2 float64) float64 {
	return (n1 + n2) / 2
}

func situacao(n float64) string {
	if n >= 7 {
		return "Aprovado"
	} else if n >= 5 {
		return "Recuperação"
	} else {
		return "Reprovado"
	}
}
func main() {
	var nota1, nota2 float64
	fmt.Print("Digite a primeira nota: ")
	fmt.Scanln(&nota1)
	fmt.Print("Digite a segunda nota: ")
	fmt.Scanln(&nota2)

	m := media(nota1, nota2)
	fmt.Println("Média:", m)
	fmt.Println("Situação:", situacao(m))
}
