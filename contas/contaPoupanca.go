package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operacao int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorSaque float64) string {
	if valorSaque < 0 {
		return "Você não pode sacar um valor negativo!"
	}

	if valorSaque > c.saldo {
		return "Saldo insuficiente!"
	}
	c.saldo -= valorSaque
	return "Saque realizado com sucesso!"
}

func (c *ContaPoupanca) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Depósito não pode ser efetuado", c.saldo
}

func (c *ContaPoupanca) Extrato() float64 {
	return c.saldo
}
