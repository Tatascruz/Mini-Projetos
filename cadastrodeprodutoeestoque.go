package main

import "fmt"

type Produto struct {
	Nome       string
	Preco      float64
	Quantidade int
}

func main() {
	produtos := []Produto{}
	var opcao int

	for {
		fmt.Println("\nSistema de Produtos")
		fmt.Println("1 - Adicionar produto")
		fmt.Println("2 - Listar produtos")
		fmt.Println("3 - Calcular valor total do estoque")
		fmt.Println("4 - Alterar quantidade de um produto")
		fmt.Println("5 - Sair")
		fmt.Print("Escolha uma opção: ")
		fmt.Scanln(&opcao)

		if opcao == 1 {
			var nome string
			var preco float64
			var qtd int

			fmt.Print("Nome: ")
			fmt.Scanln(&nome)
			fmt.Print("Preço: ")
			fmt.Scanln(&preco)
			fmt.Print("Quantidade: ")
			fmt.Scanln(&qtd)

			produtos = append(produtos, Produto{Nome: nome, Preco: preco, Quantidade: qtd})
			fmt.Println("Produto adicionado!")

		} else if opcao == 2 {
			fmt.Println("\nLista de Produtos:")
			for i, p := range produtos {
				fmt.Printf("%d - %s | Preço: R$%.2f | Quantidade: %d\n", i, p.Nome, p.Preco, p.Quantidade)
			}

		} else if opcao == 3 {
			total := 0.0
			for _, p := range produtos {
				total += p.Preco * float64(p.Quantidade)
			}
			fmt.Printf("Valor total do estoque: R$%.2f\n", total)

		} else if opcao == 4 {
			if len(produtos) == 0 {
				fmt.Println("Não há produtos cadastrados.")
				continue
			}

			var indice, novaQtd int
			fmt.Println("Escolha o índice do produto para alterar:")
			for i, p := range produtos {
				fmt.Printf("%d - %s (Qtd: %d)\n", i, p.Nome, p.Quantidade)
			}
			fmt.Print("Índice: ")
			fmt.Scanln(&indice)

			if indice >= 0 && indice < len(produtos) {
				fmt.Print("Nova quantidade: ")
				fmt.Scanln(&novaQtd)
				produtos[indice].Quantidade = novaQtd
				fmt.Println("Quantidade alterada com sucesso!")
			} else {
				fmt.Println("Índice inválido.")
			}

		} else if opcao == 5 {
			fmt.Println("Saindo...")
			break
		} else {
			fmt.Println("Opção inválida")
		}

	}
}
