package main

import (
	"banco/contas" //Para a importação dar certo é necessário usar o comando go mod init NOMEPROJETO
	"fmt"
)

func main() {
	contaDoThor := contas.ContaCorrente{
		Titular:       "Thor",
		NumeroAgencia: 22,
		NumeroConta:   40,
		Saldo:         500.00,
	}

	contaDaCharlotte := contas.ContaCorrente{
		Titular:       "Charlotte",
		NumeroAgencia: 23,
		NumeroConta:   50,
		Saldo:         100.00,
	}

	contaDoThor.Transferir(200, &contaDaCharlotte)

	fmt.Println(contaDoThor)
	fmt.Println(contaDaCharlotte)

}
