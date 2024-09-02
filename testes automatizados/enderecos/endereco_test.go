package enderecos

import "testing"

//teste unitário

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestDeEndereco(t *testing.T) {

	cenarioDeTestes := []cenarioDeTeste{
		{"Rua 124", "Rua"},
		{"Avenida Beira Mar", "Avenida"},
		{"Rodovia Beira", "Rodovia"},
		{"Estrada Golang", "Estrada"},
		{"", "Tipo Inválido"},
		{"Praça 123", "Tipo Inválido"},
	}

	for _, cenario := range cenarioDeTestes {
		tipoEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if tipoEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

}
