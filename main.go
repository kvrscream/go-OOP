package main

import (
	"banco/clientes"
	"banco/contas" //Para a importação dar certo é necessário usar o comando go mod init NOMEPROJETO
	"fmt"
)

func main() {
	contaDoThor := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Thor",
			CPF:       "111.222.222-00",
			Profissao: "Cachorro de madame",
		},
		NumeroAgencia: 22,
		NumeroConta:   40,
		// saldo:         500.00,
	}
	contaDoThor.Depositar(500)

	contaDaCharlotte := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "Charlotte",
			CPF:       "333.444.555-01",
			Profissao: "Dona da madame",
		},
		NumeroAgencia: 23,
		NumeroConta:   50,
		// saldo:         100.00,
	}
	contaDaCharlotte.Depositar(100)

	contaDoThor.Transferir(200, &contaDaCharlotte)
	PagarBoleto(&contaDoThor, 100)

	fmt.Println(contaDoThor.Extrato())
	fmt.Println(contaDaCharlotte.Extrato())

	contaFelipe := contas.ContaPoupanca{
		Titular: clientes.Titular{
			Nome:      "Felipe",
			CPF:       "0000000000",
			Profissao: "Programador",
		},
		NumeroAgencia: 9,
		Operacao:      51,
		NumeroConta:   9090,
	}

	contaFelipe.Depositar(2000)
	PagarBoleto(&contaFelipe, 300)
	fmt.Println(contaFelipe)

}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
