package main

import "fmt"

func main() {
	// Lista de produtos e preços
	produtos := []string{"Colar", "Brinco", "Pulseira"}
	precos := []float64{52.0, 30.9, 45.6}

	total := 0.0

	// Exibir  os produtos e somar total
	fmt.Println("Produtos disponíveis:")
	for i := 0; i < len(produtos); i++ {
		fmt.Printf("%d - %s: R$ %.2f\n", i+1, produtos[i], precos[i])
		total += precos[i]
	}

	// Aplicar desconto de 10%
	desconto := 0.1
	totalComDesconto := total * (1 - desconto)

	// Mostrar resultado
	fmt.Printf("\nTotal sem desconto: R$ %.2f\n", total)
	fmt.Printf("Total com desconto de 10%%: R$ %2f\n", totalComDesconto)

	// Extra: permitir adicionar novos produtos
	var adicionar string
	for {
		fmt.Printf("\nDeseja adicionar um produto? (s/n): ")
		fmt.Scan(&adicionar)

		if adicionar == "s" {
			var nome string
			var preco float64
			fmt.Print("Nome do produto: ")
			fmt.Scan(&nome)
			fmt.Print("Preço do produto: R$ ")
			fmt.Scan(&preco)

			produtos = append(produtos, nome)
			precos = append(precos, preco)
			total += preco

			fmt.Println("Produto adicionado com sucesso!")
			fmt.Printf("Novo total: R$ %.2f\n", total)
		} else {
			break
		}
	}

	// Recalcular desconto final
	totalComDesconto = total * (1 - desconto)
	fmt.Printf("\nTotal final sem desconto: R$ %.2f\n", total)
	fmt.Printf("Total final com desconto de 10%%: R$ %.2f\n", totalComDesconto)
}
