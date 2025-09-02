package main

import (
	"fmt"
)

func main() {
	// Número secreto definido pelo jogo
	var segredo int = 7
	// Variável que armazenar a tentativa do jagor
	var tentativa int

	fmt.Println("=== Jogo da Adivinhação ===")
	fmt.Println("Você tem 5 tentativas para acertar o número secreto!")

	// Loop para permitir até 5 tentativas
	for i := 1; i <= 5; i++ {
		fmt.Printf("Tentativa %d: Digite um número: ", i)
		fmt.Scanln(&tentativa)

		// Verificar se acertou
		if tentativa == segredo {
			fmt.Println("Parabéns! Você acertou o número secreto!")
			break
		} else if tentativa < segredo {
			fmt.Println("Tente um número maior.")
		} else {
			fmt.Println("Tente um número menor.")
		}

		// Caso seja a última tentativa
		if i == 5 {
			fmt.Printf("Fim do jogo! O número secreto era %d. \n", segredo)
		}

	}
}
