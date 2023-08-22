package main

import (
	"banco/contas"
	"fmt"
)

type ContaInterface interface {
	Sacar(valor float64) (bool, string)
}

func PagarBoleto(conta ContaInterface, valor float64) {
	conta.Sacar(valor)
}

func main() {

	contaEduardo := contas.ContaPoupanca{}
	contaEduardo.Depositar(1000)
	PagarBoleto(&contaEduardo, 500)
	fmt.Println(contaEduardo.ObterSaldo())
}
