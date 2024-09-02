package enderecos

import "strings"

// tipo de endereço
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}

	enderecoLetraMinuscula := strings.ToUpper((endereco))

	primeiraPalavraEndereco := strings.Split(enderecoLetraMinuscula, " ")[0]

	enderecoTemTipoValido := false

	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraEndereco {
			enderecoTemTipoValido = true
		}
	}

	if enderecoTemTipoValido {
		return primeiraPalavraEndereco
	}

	return "Tipo Inválido"
}
