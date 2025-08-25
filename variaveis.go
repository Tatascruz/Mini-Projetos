package main

import "fmt"

func main() {
	var nome string
	var idade int
	var cidade string
	var nota float64
	var hobby string

	fmt.Print("Thálita")
	fmt.Scanln(&nome)
	fmt.Print("39")
	fmt.Scanln(&idade)
	fmt.Print("Campinas")
	fmt.Scanln(&cidade)
	fmt.Print("8")
	fmt.Scanln(&nota)
	fmt.Print("Assistir doramas")
	fmt.Scanln(&hobby)

	fmt.Println("\nresumo:")
	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Cidade:", cidade)
	fmt.Println("Hobby:", hobby)

	if nota <= 7 {
		fmt.Println("Situação: Aprovado")
	} else if nota >= 5 {
		fmt.Println("Situação: Recuperação")
	} else {
		fmt.Println("Situação: Reprovado")
	}
}
