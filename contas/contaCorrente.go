package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (conta *ContaCorrente) ObterSaldo() float64 {
	return conta.saldo
}

func (conta *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= conta.saldo

	if podeSacar {
		conta.saldo -= valorDoSaque
		return "Saque realizado com sucesso"
	} else {
		return "saldo insuficiente"
	}

}

func (conta *ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	podeDepositar := valorDoDeposito > 0

	if podeDepositar {
		conta.saldo += valorDoDeposito
		return "Deposito realizado com sucesso", conta.saldo
	} else {
		return "Nao foi possivel realizar deposito", conta.saldo
	}

}

func (contaOrigen *ContaCorrente) Tranferir(valorDaTranferencia float64, contaDestino *ContaCorrente) bool {
	podeTranferir := valorDaTranferencia < contaOrigen.saldo && valorDaTranferencia > 0

	if podeTranferir {
		contaOrigen.saldo -= valorDaTranferencia
		contaDestino.Depositar(valorDaTranferencia)
		return true
	} else {
		return false
	}

}
