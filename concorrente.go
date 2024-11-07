package main

func CompraConc(comprador int, acao int, ch_compras chan int, ch_vendas chan int) {
	ch_compras <- comprador
	vendedor := <-ch_vendas

	if (relacao[vendedor][acao] == 0) || (saldo[comprador] < valores[acao]) {
		falhas++
	}else{
		transacao(comprador, vendedor, acao)
	}
	
}

func VendaConc(vendedor int, acao int, ch_compras chan int, ch_vendas chan int) {
	ch_vendas <- vendedor
	comprador := <-ch_compras

	if (relacao[vendedor][acao] == 0) || (saldo[comprador] < valores[acao]) {
		falhas++
	}else{
		transacao(comprador, vendedor, acao)
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
		go CompraConc(compras[i][0], compras[i][1], canais_compra[compras[i][1]], canais_venda[compras[i][1]])
		go VendaConc(vendas[i][0], compras[i][1],canais_compra[vendas[i][1]], canais_venda[vendas[i][1]])

	}
}
