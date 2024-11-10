package main

import (
	"math/rand"
)

func GeraSaldos() []float64 {
	listaSaldos := make([]float64, N)
	for i := 0; i < N; i++ {
		listaSaldos[i] = float64(rand.Intn(saldoMax)) // Valores aleatórios de 0 a 99 (ajuste conforme necessário)
	}
	return listaSaldos
}

func GeraValores() []float64 {
	listaValores := make([]float64, N)
	for i := 0; i < N; i++ {
		listaValores[i] = float64(rand.Intn(precoMax))
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
