package contas

import "banco/clientes"

type ContaCorrente struct { // struct ou estrutura seria o mesmo conceito de model
	Titular       clientes.Titular // Composição! No Go não temos eranças
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string { // O ponteiro antes do nome da função permite chamar direto da struct
	if valorSaque < 0 {
		return "Você não pode sacar um valor negativo!"
	}

	if valorSaque > c.saldo {
		return "Saldo insuficiente!"
	}
	c.saldo -= valorSaque
	return "Saque realizado com sucesso!"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.saldo += valor
		return "Deposito realizado com sucesso", c.saldo
	}
	return "Depósito não pode ser efetuado", c.saldo
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor < c.saldo && valor > 0 {
		c.saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) Extrato() float64 {
	return c.saldo
}
