package main

import "fmt"

// funcao para gerar uma lista de numeros
func gerarLista(n int) []int {
    lista := make([]int, n)
    for i := 0; i < n; i++ {
        lista[i] = i + 1
    }
    return lista
}

func buscaBinaria(arr []int, valor int) int {
	// declarando variaveis
	low, high, etapas := 0, len(arr)-1, 0

	for low <= high {
		// fatia o array no meio
		mid := (low + high) / 2

		// se o valor for menor que o valor do meio
		if arr[mid] == valor {
			fmt.Printf("Numero de etapas: %d\n", etapas)
			return mid
		} else if(arr[mid] < valor) {
			low = mid + 1 // O valor está na metade direita
		} else {
			high = mid - 1 // O valor está na metade esquerda
		}
		etapas++
	}


	return -1
}

type BuscaBinaria struct {
	Valor int
	LengthArr int
}

func main() {

	objetos := []BuscaBinaria{
        {Valor: 64, LengthArr: 64},
        {Valor: 128, LengthArr: 128},
        {Valor: 256, LengthArr: 256},
        {Valor: 512, LengthArr: 512},
    }

	fmt.Println("Busca Binária")
    for _, obj := range objetos {
		fmt.Println("----------------------------------")
        fmt.Printf("Valor: %d, LengthArr: %d\n", obj.Valor, obj.LengthArr)

		lista := gerarLista(obj.LengthArr)
	
		resultado := buscaBinaria(lista, obj.Valor)
		if resultado != -1 {
			fmt.Printf("Elemento encontrado na posição %d\n", resultado)
		} else {
			fmt.Println("Elemento não encontrado")
		}
    }

}