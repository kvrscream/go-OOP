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

	fmt.Println(contaDoThor.Extrato())
	fmt.Println(contaDaCharlotte.Extrato())

}
