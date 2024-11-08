package main

import (
	"fmt"
	"sync"
	"time"
)

const N = 1000
const qnt = 1000000

var (
	falhas      int32 = 0
	transacoes        = 0
	saldo       []float32
	valores     []float32
	compras     [][]int
	vendas      [][]int
	relacao     [][]int
	saldo_aux   []float32
	valores_aux []float32
	compras_aux [][]int
	vendas_aux  [][]int
	relacao_aux [][]int
	mu          sync.Mutex
	mu2         sync.Mutex
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
	saldo = make([]float32, len(saldo_aux))
	copy(valores, valores_aux)

	valores = make([]float32, len(valores_aux))
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

func transacao(comprador int, vendedor int, acao int) {
	mu.Lock()
	defer mu.Unlock()
	saldo[comprador] -= valores[acao]
	saldo[vendedor] += valores[acao]
	relacao[comprador][acao]++
	relacao[vendedor][acao]--
}

func main() {
	marcador = time.Now()
	atribuiValores()
	atribuiAuxiliares()
	fmt.Println("Tempo de atribuição dos sequenciais: ", time.Since(marcador))
	marcador = time.Now()
	sequencial()
	fmt.Println("Tempo sequencial: ", time.Since(marcador))
	marcador = time.Now()
	atribuiAuxiliares()
	fmt.Println("Tempo de atribuição dos concorrentes: ", time.Since(marcador))
	marcador = time.Now()
	concorrente()
	fmt.Println("Tempo concorrente: ", time.Since(marcador))
}
