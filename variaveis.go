package main

import "fmt"

func main() {
	var nome string
	var idade int
	var cidade string
	var nota float64
	var hobby string

	fmt.Print("Digite seu nome: ")
	fmt.Scanln(&nome)
	fmt.Print("Digite sua idade: ")
	fmt.Scanln(&idade)
	fmt.Print("Digite sua cidade: ")
	fmt.Scanln(&cidade)
	fmt.Print("Digite sua nota: ")
	fmt.Scanln(&nota)
	fmt.Print("Digite seu hobby: ")
	fmt.Scanln(&hobby)

	fmt.Println("\nResumo:")
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Cidade:", cidade)
	fmt.Println("Hobby:", hobby)

	if nota >= 7 {
		fmt.Println("Situação: Aprovado")
	} else if nota >= 5 {
		fmt.Println("Situação: Recuperação")
	} else {
		fmt.Println("Situação: Reprovado")
	}
}
