package main

import (
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func CompraConc(indice int, ch_compras chan int, ch_vendas chan int) {
	defer wg.Done()
	mu2.Lock()
	comprador := compras[indice][0]
	acao := compras[indice][1]
	mu2.Unlock()
	ch_compras <- comprador

	select {
	case vendedor := <-ch_vendas:
		// Se uma mensagem for recebida no canal ch_vendas
		if (relacao[vendedor][acao] == 0) || (saldos[comprador] < valores[acao]) {
			atomic.AddInt32(&falhas, 1)
		} else {
			transacao(comprador, vendedor, acao)
			valores[acao] *= multAcao
		}
	case <-time.After(3 * time.Second): // Timeout de 2 segundos
		// Código para o caso de timeout, caso não haja resposta de ch_vendas
		atomic.AddInt32(&falhas, 1)
	}
}

func VendaConc(indice int, ch_compras chan int, ch_vendas chan int) {
	defer wg.Done()
	mu2.Lock()
	vendedor := vendas[indice][0]
	acao := vendas[indice][1]
	mu2.Unlock()
	ch_vendas <- vendedor

	select {
	case comprador := <-ch_compras:
		// Se uma mensagem for recebida no canal ch_compras
		if (relacao[vendedor][acao] == 0) || (saldos[comprador] < valores[acao]) {
			atomic.AddInt32(&falhas, 1)
		} else {
			transacao(comprador, vendedor, acao)
			valores[acao] /= multAcao
		}
	case <-time.After(3 * time.Second): // Timeout de 2 segundos
		// Código para o caso de timeout, caso não haja resposta de ch_compras
		atomic.AddInt32(&falhas, 1)
	}
}

func concorrente() {
	falhas = 0
	canais_compra := make([]chan int, N)
	canais_venda := make([]chan int, N)
	for i := 0; i < N; i++ {
		canais_compra[i] = make(chan int, qnt)
		canais_venda[i] = make(chan int, qnt)
	}
	for i := 0; i < qnt; i++ {
		wg.Add(2)
		go CompraConc(i, canais_compra[compras[i][1]], canais_venda[compras[i][1]])
		go VendaConc(i, canais_compra[vendas[i][1]], canais_venda[vendas[i][1]])
	}
	wg.Wait()
}
