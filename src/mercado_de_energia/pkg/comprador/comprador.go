package comprador

import (
	"fmt"
)

type EConsumidor struct {
	Id              int     //id
	TarifaDesejavel float64 // tarifa desejavel
	PrecoMaximo     float64 // maximo preco admissivel
	PrazoContrato   float64 // prazo do contrato do cliente
	Demanda         float64 // demanda do cliente
}

// Inicialização da estrutura de dados
func (c *EConsumidor) Inicia_EConsumidor() {
	fmt.Printf("\nCadastrar dados Consumidor: %d\n", c.Id)
	fmt.Print("Prazo de contrato do Consumidor [s]:")
	valor := setValores()
	c.PrazoContrato = valor
	fmt.Print("Demanda Consumidor [kW]:")
	valor = setValores()
	c.Demanda = valor
	fmt.Print("Máximo preço admissível [R$/kW]:")
	valor = setValores()
	c.PrecoMaximo = valor
	fmt.Print("Tarifa desejável [R$/kW]:")
	valor = setValores()
	c.TarifaDesejavel = valor
}

func (c *EConsumidor) AtualizaPA() {
	if c.PrazoContrato <= 15 {
		c.PrecoMaximo += (c.PrecoMaximo * 0.22)
	}
}

func setValores() float64 {
	var valor float64
	fmt.Scan(&valor)
	fmt.Println("")
	return valor
}

//Atualiza tA
//void atualiza_pa(struct_EConsumidor *, double, double, double)
