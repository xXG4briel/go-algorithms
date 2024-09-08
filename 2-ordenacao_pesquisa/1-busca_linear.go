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

func buscaLinear(arr []int, valor int) int {
	etapas := 0
	for i, v := range arr {
		etapas++
		if v == valor {
			fmt.Println("Numero de etapas:", etapas)
			return i
		}
	}


	return -1
}

type BuscaLinear struct {
	Valor int
	LengthArr int
}

func main() {

	objetos := []BuscaLinear{
        {Valor: 64, LengthArr: 64},
        {Valor: 128, LengthArr: 128},
        {Valor: 256, LengthArr: 256},
        {Valor: 512, LengthArr: 512},
    }

	fmt.Println("Busca Linear")
	for _, obj := range objetos {
		fmt.Println("----------------------------------")
        fmt.Printf("Valor: %d, LengthArr: %d\n", obj.Valor, obj.LengthArr)

		lista := gerarLista(obj.LengthArr)
	
		resultado := buscaLinear(lista, obj.Valor)
	
		if resultado == -1 {
			fmt.Println("O valor", obj.Valor, "não foi encontrado no array.")
		} else {
			fmt.Println("O valor", obj.Valor, "foi encontrado no índice", resultado, "do array.")
		}
	}


}