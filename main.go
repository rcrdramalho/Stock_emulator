package main

import (
	"fmt"
	"sync"
	"time"
)

// Dados constantes do programa
const (
	N        = 1000  //Número de usuários e tipos de ações no mercado
	qnt      = 10000 // Quantidade de transações de compra e de venda
	saldoMax = 10000 // Saldo máximo inicial de cada usuário
	precoMax = 100   // Preço máximo inicial de cada ação
	multAcao = 1.02  // Multiplicador de valorização e desvalorização de ações
)

// Dados variáveis do programas
var (
	falhas  int32     = 0 // Número de transações não concluídas
	saldos  []float64     // Vetor com o saldo de cada usuário
	valores []float64     // Vetor com o preço de cada ação
	compras [][]int       // Lista de ordens de compra da forma [comprador] quer comprar [ação]
	vendas  [][]int       // Lista de ordens de venda da forma [vendedor] quer vender [ação]
	relacao [][]int       // Relação da quantidade que cada vendedor tem de cada ação da forma [vendedor][ação] = quantidade
	// Listas _aux criadas para testar sequencial e concorrentemente sem alterar os valores originais
	saldo_aux   []float64
	valores_aux []float64
	compras_aux [][]int
	vendas_aux  [][]int
	relacao_aux [][]int
	mu          sync.Mutex // Variável de exclusão mútua para sincronização das transações
	mu2         sync.Mutex // Variável de exclusão mútua para sincronização das compras/vendas
	marcador    time.Time  // Marcador de tempo para medição
)

// Função que preenche as listas auxiliares com os valores que serão usados na simulação
func atribuiValores() {
	saldo_aux = GeraSaldos()
	valores_aux = GeraValores()
	compras_aux = GeraTransacoes()
	vendas_aux = GeraTransacoes()
	relacao_aux = GeraRelacoes()

}

// Função que copia os valores das listas auxiliares para as que serão usadas nas execuções
func atribuiAuxiliares() {
	saldos = make([]float64, len(saldo_aux))
	copy(saldos, saldo_aux)

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

// Função que soma elementos de um vetor
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

// Função que soma elementos de uma matriz
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

// Função resposável por realizar as transações
func transacao(comprador int, vendedor int, acao int) {
	mu.Lock()
	defer mu.Unlock()                  //Exclusão mútua para que não haja sobrescrita
	saldos[comprador] -= valores[acao] // Retira o saldo do comprador
	saldos[vendedor] += valores[acao]  // E adiciona ao saldo do vendedor
	relacao[vendedor][acao]--          // Retira a ação do vendedor
	relacao[comprador][acao]++         // E adiciona ao comprador
}

// Função que printa métricas de uma execução para medição
func imprimirMetricas(marcador time.Time) {
	fmt.Println("Tempo:", time.Since(marcador))                                       // Tempo de execução
	fmt.Println("Saldo total:", somaVetor(saldo_aux))                                 // Soma do saldo de todos os usuários
	fmt.Println("Quantidade de ações:", somaMatriz(relacao_aux))                      // Soma da quantidade de ações presente na simulação
	fmt.Println("Falhas:", falhas, "(", float64(falhas)*100/(2*float64(qnt)), "%)\n") // Número de transações não concluídas
}

func main() {
	marcador = time.Now() //Inicia o cronômetro
	atribuiValores()      // Gera os valores do universo simulado
	fmt.Println("Simulação com", N, "usuários,", N, "ações e", qnt*2, "transações de compra e venda.")
	fmt.Println("Usuário possuem um saldos máximo de", saldoMax, "rupias, as ações tem um valor inicial de no máximo", precoMax, "rupias e o multiplicador de valorização por ação é de", multAcao)
	fmt.Println("Dados antes da simulação:")
	imprimirMetricas(marcador) // Printa métricas pré-execução
	atribuiAuxiliares()        // Copia os valores para as listas que serão executadas
	marcador = time.Now()      // Reseta o cronometro
	sequencial()               // Execução sequencial
	fmt.Println("Dados após a simulação sequencial:")
	imprimirMetricas(marcador) // Printa métricas da execução sequencial
	atribuiAuxiliares()        // Copia novamente os valores para as listas que serão executadas
	marcador = time.Now()      // Reseta o cronometro
	concorrente()              // Execução concorrente
	fmt.Println("Dados após a simulação concorrente:")
	imprimirMetricas(marcador) // Printa métricas da execução concorrente
}
