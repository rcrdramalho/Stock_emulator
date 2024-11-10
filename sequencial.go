package main

// Função que realiza uma ordem de compra
func CompraSeq(comprador int, acao int) bool {

	vendedor := -1             // Vendedor inicial inválido
	id := -1                   // Índice inicial inválido
	for i := 0; i < qnt; i++ { // Procura um vendedor que queira vender a ação
		if vendas[i][1] == acao {
			vendedor = vendas[i][0] // Caso haja, grava o vendedor
			id = i                  // Grava o índice da lista de ordens de venda no qual o vendedor é achado
			break
		}
	}
	if (vendedor == -1) || (relacao[vendedor][acao] == 0) || (saldos[comprador] < valores[acao]) {
		// Caso não tenha sido achado vendedor, ou o vendedor achado não tenha aquela ação, ou o comprador não tenha saldo para compra a ação
		falhas++     // Adiciona essa ordem à quantidade de falhas
		return false // E encerra a compra
	}

	transacao(comprador, vendedor, acao) // Realiza a transação
	vendas[id][0] = -1                   // Retira o vendedor
	vendas[id][1] = -1                   // E ação da lista de ordens de venda
	valores[acao] *= multAcao            // Valoriza a ação

	return true
}

func VendaSeq(vendedor int, acao int) bool {

	comprador := -1            // Comprador inicial inválido
	id := -1                   // Índice inicial inválido
	for i := 0; i < qnt; i++ { // Procura um comprador que queira comprar a ação
		if compras[i][1] == acao {
			comprador = compras[i][0] // Caso haja, grava o comprador
			id = i                    // Grava o índice da lista de ordens de compra no qual o comprador é achado
			break
		}
	}
	if (comprador == -1) || (relacao[vendedor][acao] == 0) || (saldos[comprador] < valores[acao]) {
		// Caso não tenha sido achado comprador, ou o vendedor não tenha aquela ação, ou o comprador não tenha saldo para compra a ação
		falhas++     // Adiciona essa ordem à quantidade de falhas
		return false // E encerra a compra
	}

	transacao(comprador, vendedor, acao) // Realiza a transação
	compras[id][0] = -1                  // Retira o comprador
	compras[id][1] = -1                  // E ação da lista de ordens de compra
	valores[acao] /= multAcao            // Desvaloriza a ação

	return true
}

func sequencial() {
	falhas = 0
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
