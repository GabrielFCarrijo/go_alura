package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, float64) {
	sacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo

	if sacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso: ", c.Saldo
	} else {
		return "Saque indisponivel: ", c.Saldo
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	depositar := valorDoDeposito > 0

	if depositar {
		c.Saldo += valorDoDeposito
		return "Deposito realizado com sucesso: ", c.Saldo
	} else {
		return "Valor invalido: ", c.Saldo
	}
}

func (c *ContaCorrente) Transferir(valorDeTransferencia float64, contaDestino *ContaCorrente) string {
	if valorDeTransferencia < c.Saldo && valorDeTransferencia > 0 {
		c.Saldo -= valorDeTransferencia
		contaDestino.Depositar(valorDeTransferencia)
		return "Transferencia feita com sucesso"
	} else {
		return "Ouve algum problema na hora da transferencia"
	}
}
