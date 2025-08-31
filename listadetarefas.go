package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	tarefas := []string{}
	var opcao int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nLista de Tarefas")
		fmt.Println("1 - Adicionar tarefa")
		fmt.Println("2 - Ver tarefas")
		fmt.Println("3 - Sair")
		fmt.Println("Escolha uma opção: ")
		fmt.Scanln(&opcao)

		if opcao == 1 {
			fmt.Print("Digite a tarefa: ")
			entrada, _ := reader.ReadString('\n')
			novaTarefa := strings.TrimSpace(entrada)
			tarefas = append(tarefas, novaTarefa)
			fmt.Println("Tarefa adicionada!")
		} else if opcao == 2 {
			if len(tarefas) == 0 {
				fmt.Println("Nenhuma tarefa adicionada ainda.")
			} else {
				fmt.Println("Tarefas atuais:")
				for i, t := range tarefas {
					fmt.Printf("%d - %s\n", i+1, t)
				}

			}
		} else if opcao == 3 {
			fmt.Println("Saindo do programa...")
			break
		} else {
			fmt.Println("Opção inválida")
		}
	}
}
