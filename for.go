package main

import "fmt"

func main() {
	numeros := []int{23, 256, 67, 82, 1013}
	soma := 0

	for i := 0; i < len(numeros); i++ {
		if numeros[i]%2 == 0 {
			fmt.Println(numeros[i], "é par")
		} else {
			fmt.Println(numeros[i], "é impar")
		}
		soma += numeros[i]
	}

	fmt.Println("Soma total:", soma)

}
