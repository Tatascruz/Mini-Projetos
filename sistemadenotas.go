package main

import "fmt"

func media(notas []float64) float64 {
	soma := 0.0
	for _, n := range notas {
		soma += n
	}
	return soma / float64(len(notas))
}
func situacao(m float64) string {
	if m >= 7 {
		return "Aprovado"
	} else if m >= 5 {
		return "Recuperação"
	}
	return "Reprovado"
}

func main() {
	var qtdAlunos int
	fmt.Print("Quantos alunos deseja cadastrar? ")
	fmt.Scan(&qtdAlunos)

	for i := 0; i < qtdAlunos; i++ {
		var nome string
		var n1, n2 float64

		fmt.Printf("\nDigite o nome do aluno %d: ", i+1)
		fmt.Scan(&nome)

		fmt.Printf("Digite a primeira nota de %s: ", nome)
		fmt.Scan(&n1)

		fmt.Printf("Digite a segunda nota de %s: ", nome)
		fmt.Scan(&n2)

		notas := []float64{n1, n2}
		m := media(notas)
		s := situacao(m)

		fmt.Printf("\nAluno: %s - Média: %.2f - Situação: %s\n", nome, m, s)
	}
}
