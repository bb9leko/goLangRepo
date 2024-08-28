package main

import (
	"encoding/json"
	"os"
)

// Omitindo valor no unmarshal com a tag "-".... como se fosse as Annotations ou Decorators no Go
type Operacao struct {
	Quantidade int     `json:"quantidade"`
	Valor      float64 `json:"-" validate:"vlr=0"`
}

func main() {

	operacao := Operacao{Quantidade: 100, Valor: 50.0}

	result, err := json.Marshal(operacao)
	if err != nil {
		println(err)
	}
	println(string(result))

	//Encoder - Não precisa guardar o resultado da variável. Já devolve o resutado.
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(operacao)
	//Encoder com apenas 1 linha
	err = json.NewEncoder(os.Stdout).Encode(operacao)
	if err != nil {
		println(err)
	}

	// No Go o Json puro é um slice de dados
	jsonPuro := []byte(`{"qtd":100,"vlr":50.0}`)
	//Criando uma variável de Conta
	var opercaoX Operacao
	err = json.Unmarshal(jsonPuro, &opercaoX)
	if err != nil {
		println(err)
	}
	println(opercaoX.Quantidade)

	//Decoder
	//json.NewDecoder(os.Stdout).Decode(result)

}
