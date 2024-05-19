package main

import (
	"fmt"
	"go-oop-study/clientes"
	"go-oop-study/contas"
)

func main() {

	contaCorrente1 := contas.ContaCorrente{Titular: clientes.Titular{Cpf: "111.254.254-89", Nome: "João", Profissao: "desenvolvedor"},
		NumeroAgencia: 123,
		NumeroConta:   123456,
	}

	fmt.Println(contaCorrente1)

	fmt.Println(contaCorrente1.Sacar(159.0))

	fmt.Println(contaCorrente1.Depositar(10))

	var contaCorrente2 *contas.ContaCorrente
	contaCorrente2 = new(contas.ContaCorrente)

	titular2 := clientes.Titular{Cpf: "124.588.524-56"}

	contaCorrente2.Titular = titular2

	fmt.Println(*contaCorrente2)
	fmt.Println(contaCorrente1.Transferir(10, contaCorrente2))
	fmt.Println(contaCorrente1.ObterSaldo())
	fmt.Println(contaCorrente2.ObterSaldo())

	PagarBoleto(&contaCorrente1, 5.)
	fmt.Println(contaCorrente1.ObterSaldo())
	fmt.Println(contaCorrente2.ObterSaldo())

}

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}
