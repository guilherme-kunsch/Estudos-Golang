package enderecos

import "testing"

//teste unitário

func TestDeEndereco(t *testing.T) {
	enderecoParaTeste := "Avenida"
	enderecoEsperado := "avenida"

	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != enderecoEsperado {
		t.Error("O tipo esperado é inválido")
	}
}
