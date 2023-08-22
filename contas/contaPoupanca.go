package contas

import (
	"banco/clientes"
	"fmt"
	"os"
)

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valor float64) (bool, string) {
	validaValor(valor)

	mensagem := "saldo insuficiente"
	podeSacar := c.podeSacar(valor)

	if podeSacar {
		c.saldo -= valor
		mensagem = "Saque realizado com sucesso"
	}
	return podeSacar, mensagem
}

func validaValor(valor float64) {
	if valor <= 0.0 {
		fmt.Println("Valor inválido. Necessário que valor seja maior que zero!", valor)
		os.Exit(-1)
	}
}

func (conta *ContaPoupanca) podeSacar(valor float64) bool {
	return conta.saldo >= valor
}

func (conta *ContaPoupanca) Depositar(valor float64) {
	ValidaValor(valor)
	conta.saldo += valor
	fmt.Println("Depositor realizado com sucesso!")
}

func (conta *ContaPoupanca) ObterSaldo() float64 {
	return conta.saldo
}
