package contas

import (
	"banco/clientes"
	"fmt"
	"os"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	saldo                      float64
}

func (c *ContaCorrente) Sacar(valor float64) (bool, string) {
	ValidaValor(valor)

	mensagem := "saldo insuficiente"
	podeSacar := c.PodeSacar(valor)

	if podeSacar {
		c.saldo -= valor
		mensagem = "Saque realizado com sucesso"
	}
	return podeSacar, mensagem
}

func ValidaValor(valor float64) {
	if valor <= 0.0 {
		fmt.Println("Valor inválido. Necessário que valor seja maior que zero!", valor)
		os.Exit(-1)
	}
}

func (conta *ContaCorrente) PodeSacar(valor float64) bool {
	return conta.saldo >= valor
}

func (conta *ContaCorrente) Depositar(valor float64) {
	ValidaValor(valor)
	conta.saldo += valor
	fmt.Println("Depositor realizado com sucesso!")
}

func (conta *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) {
	ValidaValor(valor)

	if conta.PodeSacar(valor) {
		conta.Sacar(valor)
		contaDestino.Depositar(valor)
		fmt.Println("Tranferencia realizada com sucesso.")
	} else {
		fmt.Println("saldo insuficiente")
	}
}

func (conta *ContaCorrente) ObterSaldo() float64 {
	return conta.saldo
}
