package main

import "fmt"

type Usuario struct {
	Nome  string
	Idade int
	Email string
}

func main() {
	usuarios := []Usuario{}
	var opcao int

	for {
		fmt.Println("\nSitema de Usuários")
		fmt.Println("1 - Adicionar usuário")
		fmt.Println("2 - Listar usuários")
		fmt.Println("3 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&opcao)

		if opcao == 1 {
			var nome, email string
			var idade int
			fmt.Print("Nome: ")
			fmt.Scanln(&nome)
			fmt.Print("Idade: ")
			fmt.Scanln(&idade)
			fmt.Print("Email: ")
			fmt.Scanln(&email)
			usuarios = append(usuarios, Usuario{Nome: nome, Idade: idade, Email: email})
			fmt.Println("Usuário adicionado!")
		} else if opcao == 2 {
			fmt.Println("Lista de Usuários:")
			for i, u := range usuarios {
				fmt.Println(i+1, "-", u.Nome, u.Idade, u.Email)
			}
		} else if opcao == 3 {
			fmt.Println("Saindo...")
			break
		} else {
			fmt.Println("Opção inválida")
		}
	}
}
