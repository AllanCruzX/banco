package main

import (
	"banco/contas"
	"fmt"
)

func PagarBoleto(conta contas.Conta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {

	contaDaLuisa := contas.ContaPoupanca{}
	contaDaLuisa.Depositar(500)
	PagarBoleto(&contaDaLuisa, 200)
	fmt.Println(contaDaLuisa.ObterSaldo())
}

func testes() {

	// clienteBruno := clientes.Titular{"Bruno", "01559874595", "Desenvolvedor Golang"}
	// contaDoBruno := contas.ContaCorrente{Titular: clienteBruno, NumeroAgencia: 123, NumeroConta: 123456}

	// contaDoBruno.Depositar(100)
	// fmt.Println(contaDoBruno.ObterSaldo())

	// contaDoAllan := contas.ContaCorrente{
	// 	Titular:       clientes.Titular{Nome: "Allan", CPF: "01559874591", Profissao: "Desenvolvedor"},
	// 	NumeroAgencia: 589,
	// 	NumeroConta:   123456,
	// 	Saldo:         125.5}

	// fmt.Println(contaDoAllan)

	//contaDaBruna := contas.ContaCorrente{"Bruna", 222, 11222, 200}

	// var contaDaCris *ContaCorrente
	// contaDaCris = new(ContaCorrente)
	// contaDaCris.titular = "Cris"

	// fmt.Println(contaDoAllan)
	// fmt.Println(contaDaBruna)
	// fmt.Println(*contaDaCris)

	//fmt.Println(contaDoAllan.Sacar(50))

	// status, valor := contaDoAllan.Depositar(2000)
	// fmt.Println(status, valor)

	// status := contaDaBruna.Tranferir(50, &contaDoAllan)
	// fmt.Println(status)
	// fmt.Println(contaDoAllan)
	// fmt.Println(contaDaBruna)
}

func executaPonteiro() {
	a := 10
	//& -> o endereco de memoria
	fmt.Println(&a)

	// * -> nesse contexto inicia o ponteiro
	var ponteiro *int = &a

	// * -> nesse contexto exibo o valor  que está no endereço
	fmt.Println(*ponteiro)

	// * -> nesse contexto exibo o valor o endereço
	fmt.Println(ponteiro)
}
