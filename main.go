package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 1000
const qnt = 1000000
const saldoMax = 10000
const precoMax = 100
const multAcao = 1.02

var (
	falhas      int32 = 0
	saldo       []float64
	valores     []float64
	compras     [][]int
	vendas      [][]int
	relacao     [][]int
	saldo_aux   []float64
	valores_aux []float64
	compras_aux [][]int
	vendas_aux  [][]int
	relacao_aux [][]int
	mu          sync.Mutex
	marcador    time.Time
)

func atribuiValores() {
	saldo_aux = GeraSaldos()
	valores_aux = GeraValores()
	compras_aux = GeraTransacoes()
	vendas_aux = GeraTransacoes()
	relacao_aux = GeraRelacoes()

}

func atribuiAuxiliares() {
	saldo = make([]float64, len(saldo_aux))
	copy(saldo, saldo_aux)

	valores = make([]float64, len(valores_aux))
	copy(valores, valores_aux)

	compras = make([][]int, len(compras_aux))
	for i := range compras_aux {
		compras[i] = make([]int, len(compras_aux[i]))
		copy(compras[i], compras_aux[i])
	}

	vendas = make([][]int, len(vendas_aux))
	for i := range vendas_aux {
		vendas[i] = make([]int, len(vendas_aux[i]))
		copy(vendas[i], vendas_aux[i])
	}

	relacao = make([][]int, len(relacao_aux))
	for i := range relacao_aux {
		relacao[i] = make([]int, len(relacao_aux[i]))
		copy(relacao[i], relacao_aux[i])
	}
}

func somaVetor(vetor interface{}) float64 {
	var soma float64 = 0

	switch v := vetor.(type) {
	case []int:
		for _, valor := range v {
			soma += float64(valor)
		}
	case []float64:
		for _, valor := range v {
			soma += valor
		}
	default:
		fmt.Println("Tipo de vetor não suportado")
	}

	return soma
}

func somaMatriz(matriz interface{}) float64 {
	var soma float64 = 0

	switch m := matriz.(type) {
	case [][]int:
		for _, linha := range m {
			for _, valor := range linha {
				soma += float64(valor)
			}
		}
	case [][]float64:
		for _, linha := range m {
			for _, valor := range linha {
				soma += valor
			}
		}
	default:
		fmt.Println("Tipo de matriz não suportado")
	}

	return soma
}

func transacao(comprador int, vendedor int, acao int) {
	mu.Lock()
	defer mu.Unlock()
	saldo[comprador] -= valores[acao]
	saldo[vendedor] += valores[acao]
	relacao[comprador][acao]++
	relacao[vendedor][acao]--
}

func imprimirMetricas(marcador time.Time) {
	fmt.Println("Tempo:", time.Since(marcador))
	fmt.Println("Saldo total:", somaVetor(saldo_aux))
	fmt.Println("Quantidade de ações:", somaMatriz(relacao_aux))
	fmt.Println("Falhas:", falhas, "(", float64(falhas)*100/(2*float64(qnt)), "%)")
	fmt.Println("")
}

func main() {
	marcador = time.Now()
	atribuiValores()
	fmt.Println("Simulação com", N, "usuários,", N, "ações e", qnt*2, "transações de compra e venda.")
	fmt.Println("Usuário possuem um saldo máximo de", saldoMax, "rupias, as ações tem um valor inicial de no máximo", precoMax, "rupias e o multiplicador de valorização por ação é de", multAcao)
	fmt.Println("Dados antes da simulação:")
	imprimirMetricas(marcador)
	atribuiAuxiliares()
	marcador = time.Now()
	sequencial()
	fmt.Println("Dados após a simulação sequencial:")
	imprimirMetricas(marcador)
	atribuiAuxiliares()
	marcador = time.Now()
	concorrente()
	fmt.Println("Dados após a simulação concorrente:")
	imprimirMetricas(marcador)
}
