package contas

type ContaCorrente struct { // struct ou estrutura seria o mesmo conceito de model
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string { // O ponteiro antes do nome da função permite chamar direto da struct
	if valorSaque < 0 {
		return "Você não pode sacar um valor negativo!"
	}

	if valorSaque > c.Saldo {
		return "Saldo insuficiente!"
	}
	c.Saldo -= valorSaque
	return "Saque realizado com sucesso!"
}

func (c *ContaCorrente) Depositar(valor float64) (string, float64) {
	if valor > 0 {
		c.Saldo += valor
		return "Deposito realizado com sucesso", c.Saldo
	}
	return "Depósito não pode ser efetuado", c.Saldo
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	if valor < c.Saldo && valor > 0 {
		c.Saldo -= valor
		contaDestino.Depositar(valor)
		return true
	} else {
		return false
	}
}
