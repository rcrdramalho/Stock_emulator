package main

func transacao(comprador int, vendedor int, acao int) {
	saldo[comprador] -= valores[acao]
	saldo[vendedor] += valores[acao]
	relacao[comprador][acao]++
	relacao[vendedor][acao]--
}

func CompraSeq(comprador int, acao int) bool {

	vendedor := -1
	id := -1
	for i := 0; i < qnt; i++ {
		if vendas[i][1] == acao {
			vendedor = vendas[i][0]
			id = i
			break
		}
	}
	if (vendedor == -1) || (relacao[vendedor][acao] == 0) || (saldo[comprador] < valores[acao]) {
		falhas++
		return false
	}

	transacao(comprador, vendedor, acao)
	vendas[id][0] = -1
	vendas[id][1] = -1
	valores[acao] *= 1

	return true
}

func VendaSeq(vendedor int, acao int) bool {

	comprador := -1
	id := -1
	for i := 0; i < qnt; i++ {
		if compras[i][1] == acao {
			comprador = compras[i][0]
			id = i
			break
		}
	}
	if (comprador == -1) || (relacao[comprador][acao] == 0) || (saldo[comprador] < valores[acao]) {
		return false
	}

	transacao(comprador, vendedor, acao)
	valores[acao] /= 1
	compras[id][0] = -1
	compras[id][1] = -1

	return true
}

func sequencial() {

	for i := 0; i < qnt; i++ {
		if compras[i][0] != -1 {
			if CompraSeq(compras[i][0], compras[i][1]) {
				compras[i][0] = -1
				compras[i][1] = -1
			}

		}
		if vendas[i][0] != -1 {
			if VendaSeq(vendas[i][0], vendas[i][1]) {
				vendas[i][0] = -1
				vendas[i][1] = -1
			}
		}

	}

}
