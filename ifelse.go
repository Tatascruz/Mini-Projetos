package main

import "fmt"

func main() {
	var nome string
	var idade int
	var nota float64

	// Entrada de dados
	fmt.Print("Digite seu nome: ")
	fmt.Scan(&nome)

	fmt.Print("Digite sua idade: ")
	fmt.Scan(&idade)

	fmt.Print("Digite sua nota (0-10): ")
	fmt.Scan(&nota)

	// Resumo
	fmt.Println("\nResumo:")
	fmt.Println("Nome:", nome)

	// Verifica idade
	if idade >= 18 {
		fmt.Println("Idade:", idade, "_ Adulto")
	} else if idade >= 13 {
		fmt.Println("Idade:", idade, "_ Adolescente")
	} else {
		fmt.Println("Idade:", idade, "_ Criança")
	}

	//Verifica nota
	if nota >= 7 {
		fmt.Println("Nota:", nota, "_ Aprovado")
	} else if nota >= 5 {
		fmt.Println("Nota:", nota, "_ Recuperação")
	} else {
		fmt.Println("Nota:", nota, "_ Reprovado")
	}
}
