package main

import (
	"fmt"
	"mercado_de_energia/pkg/fornecedor"
)

func main() {

	fmt.Println("Bem vindo ao Mercado de Energia!")

	fmt.Println("\n------------------------")
	fmt.Println("Dados dos fornecedores")
	fmt.Println("------------------------")

	fornecedor1 := fornecedor.Efornecedor{}
	fornecedor1.Inicia_Efornecedor()
	fmt.Println("Id: 1\nCapacidade de carga [kW]:", fornecedor1.CapacidadeCarga)

	fornecedor2 := fornecedor.Efornecedor{}
	fornecedor2.Inicia_Efornecedor()

	fornecedor3 := fornecedor.Efornecedor{}
	fornecedor3.Inicia_Efornecedor()

	//fmt.Println(fornecedor1) //debug
	//fmt.Println(fornecedor2) //debug
	//fmt.Println(fornecedor3) //debug

	fmt.Println("\nEncerrando o Mercado de Energia!")
}
