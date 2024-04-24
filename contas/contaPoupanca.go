package contas

import "banco/clientes"

type ContaPoupanca struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (conta *ContaPoupanca) ObterSaldo() float64 {
	return conta.saldo
}

func (conta *ContaPoupanca) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo

	if podeSacar {
		conta.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}

}

func (conta *ContaPoupanca) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		conta.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", conta.saldo
	} else {
		return "Nao foi possivel realizar deposito", conta.saldo
	}

}
