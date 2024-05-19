package contas

import "go-oop-study/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	poderSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if poderSacar {
		c.saldo -= valorDoSaque
		return "Saque efetuado com sucesso"
	}
	return "saldo insuficiente"
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito efetuado com sucesso", c.saldo
	}

	return "Dpósito não efetuado", c.saldo

}

func (c *ContaCorrente) Transferir(valorTranferencia float64, contaDestino *ContaCorrente) bool {

	if valorTranferencia > 0 && c.saldo >= valorTranferencia {
		c.Sacar(valorTranferencia)
		contaDestino.Depositar(valorTranferencia)
		return true

	}
	return false
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo

}
