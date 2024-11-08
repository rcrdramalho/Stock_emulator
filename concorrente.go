package main

import (
	"sync"
	"time"
)

var wg sync.WaitGroup

func CompraConc(indice int, ch_compras chan int, ch_vendas chan int) {
	defer wg.Done()
	comprador := compras[indice][0]
	acao := vendas[indice][1]
	ch_compras <- comprador
	mu2.Lock()
	transacoes++
	mu2.Unlock()

	select {
	case vendedor := <-ch_vendas:
		// Se uma mensagem for recebida no canal ch_vendas
		if (relacao[vendedor][acao] == 0) || (saldo[comprador] < valores[acao]) {
			falhas++
		} else {
			transacao(comprador, vendedor, acao)
			valores[acao] *= 1.1
		}
	case <-time.After(3 * time.Second): // Timeout de 2 segundos
		// Código para o caso de timeout, caso não haja resposta de ch_vendas
		falhas++
	}
}

func VendaConc(indice int, ch_compras chan int, ch_vendas chan int) {
	defer wg.Done()
	vendedor := vendas[indice][0]
	acao := vendas[indice][1]
	ch_vendas <- vendedor
	mu2.Lock()
	transacoes++
	mu2.Unlock()

	select {
	case comprador := <-ch_compras:
		// Se uma mensagem for recebida no canal ch_compras
		if (relacao[vendedor][acao] == 0) || (saldo[comprador] < valores[acao]) {
			falhas++
		} else {
			transacao(comprador, vendedor, acao)
			valores[acao] /= 1.1
		}
	case <-time.After(3 * time.Second): // Timeout de 2 segundos
		// Código para o caso de timeout, caso não haja resposta de ch_compras
		falhas++
	}
}

func concorrente() {
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
