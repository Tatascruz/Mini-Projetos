package main

import "fmt"

func main() {
	// Lista de produtos e preços
	produtos := []string{"Colar", "Brinco", "Pulseira", "Anel", "Tornozeleira", "Kit com todas as peças"}
	precos := []float64{50.0, 30.8, 45.0, 25.5, 38.0, 213.60}

	total := 0.0

	// Mostra cada produto com seu preço
	fmt.Println("Lista de Produtos:")
	for i := 0; i < len(produtos); i++ {
		fmt.Printf("%d. %s - R$ %.2f\n", i+1, produtos[i], precos[i])
		total += precos[i]
	}

	// Calcula o total
	fmt.Printf("\nTotal da compra: R$ %.2f\n", total)
}
