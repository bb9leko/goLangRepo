package main

import (
	"html/template"
	"os"
)

type Operacao struct {
	Quantidade int
	Valor      float64
}

// Create a func that add two values , one int, and othe float
func add(quantidade int, valor float64) float64 {
	return float64(quantidade) * valor
}

func main() {

	operacao := Operacao{100, 1.1}
	total := add(operacao.Quantidade, operacao.Valor)
	println("Total: %f", total)
	tmp := template.New("OperacaoTemplate")
	tmp, _ = tmp.Parse("Quantidade: {{.Quantidade}} x Valor: {{.Valor}}")
	err := tmp.Execute(os.Stdout, operacao)
	if err != nil {
		panic(err)
	}

}
