package main

import (
	"fmt"
	"BancoGo/contaS"
)

func main() {

	contaGabriel := contas.ContaCorrente{
		Titular:       "GABRIEL CARRIJO",
		NumeroAgencia: 589,
		NumeroConta:   999,
		Saldo:         1309.00,
	}

	contaNayla := contas.ContaCorrente{
		Titular:       "Nayla Cintra",
		NumeroAgencia: 589,
		NumeroConta:   999,
		Saldo:         1309.00,
	}

	contaGabriel.Transferir(200, &contaNayla)

	fmt.Println(contaGabriel.Saldo)
	fmt.Println(contaNayla.Saldo)
}
