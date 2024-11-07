package main

func CompraConc(comprador int, acao int) {

}

func VendaConc(comprador int, acao int) {

}

func concorrente() {
	for i := 0; i < qnt; i++ {
		if compras[i][0] != -1 {
			go CompraConc(compras[i][0], compras[i][1])

		}
		if vendas[i][0] != -1 {
			go VendaConc(vendas[i][0], compras[i][1])

		}

	}
}
