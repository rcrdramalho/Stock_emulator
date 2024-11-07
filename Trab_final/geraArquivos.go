package main

import (
	"fmt"
	"math/rand"
	"os"
)

func GeraSaldos() []int {
	listaSaldos := make([]int, N)
	for i := 0; i < N; i++ {
		listaSaldos[i] = rand.Intn(100000) // Valores aleatórios de 0 a 99 (ajuste conforme necessário)
	}
	return listaSaldos
}

func GeraValores() []int {
	listaValores := make([]int, N)
	for i := 0; i < N; i++ {
		listaValores[i] = rand.Intn(200)
	}
	return listaValores
}

func GeraRelacoes() [][]int {
	relacoes := make([][]int, N)
	for i := range relacoes {
		relacoes[i] = make([]int, N)
		for j := range relacoes[i] {
			relacoes[i][j] = rand.Intn(500)
		}
	}
	return relacoes
}

func GeraTransacoes() [][]int {
	transacoes := make([][]int, qnt)
	for i := range transacoes {
		transacoes[i] = make([]int, 2)
		for j := range transacoes[i] {
			transacoes[i][j] = rand.Intn(N)
		}
	}
	return transacoes
}

func VetorArquivo(filename string, array []int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, value := range array {
		_, err := fmt.Fprintf(file, "%d ", value)
		if err != nil {
			return err
		}
	}
	return nil
}

func MatrizArquivo(filename string, matrix [][]int) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, row := range matrix {
		for j, value := range row {
			if j == len(row)-1 {
				fmt.Fprintf(file, "%d", value)
			} else {
				fmt.Fprintf(file, "%d ", value)
			}
		}
		fmt.Fprintln(file) // Pula para a próxima linha
	}
	return nil
}

func gera() {
	listaSaldos := GeraSaldos()
	listaValores := GeraValores()
	relacao1 := GeraRelacoes()
	compras1 := GeraTransacoes()
	vendas1 := GeraTransacoes()

	err = VetorArquivo("saldos.txt", listaSaldos)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
	} else {
		fmt.Println("Saldos salvos com sucesso em saldos.txt")
	}

	err = VetorArquivo("valores.txt", listaValores)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
	} else {
		fmt.Println("Valores salvos com sucesso em saldos.txt")
	}

	err = MatrizArquivo("relacao.txt", relacao1)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
	} else {
		fmt.Println("Relação salva com sucesso em relacoes.txt")
	}

	err = MatrizArquivo("compras.txt", compras1)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
	} else {
		fmt.Println("Compras salvas com sucesso em compras.txt")
	}

	err = MatrizArquivo("vendas.txt", vendas1)
	if err != nil {
		fmt.Println("Erro ao salvar o arquivo:", err)
	} else {
		fmt.Println("Vendas salvas com sucesso em vendas.txt")
	}
}
