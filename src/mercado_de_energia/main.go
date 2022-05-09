package main

import (
	"fmt"
	"mercado_de_energia/pkg/comprador"
	"mercado_de_energia/pkg/fornecedor"
	"os"
	"time"

	screen "github.com/inancgumus/screen"
)

func main() {

	screen.Clear()

	exibeIntroducao() //exibe informacões de introdução

	fornecedor1 := fornecedor.Efornecedor{Id: 1} //vincula o fornecedor1 como id 1
	fornecedor1.Inicia_Efornecedor()             //Inicia os valores do fornecedor
	fornecedor2 := fornecedor.Efornecedor{Id: 2} //vincula o fornecedor2 como id 2
	fornecedor2.Inicia_Efornecedor()             //Inicia os valores do fornecedor
	fornecedor3 := fornecedor.Efornecedor{Id: 3} //vincula o fornecedor3 como id 3
	fornecedor3.Inicia_Efornecedor()             //Inicia os valores do fornecedor

	consumidor1 := comprador.EConsumidor{Id: 1} //vincula o consumidor1 como id 1
	consumidor2 := comprador.EConsumidor{Id: 2} //vincula o consumidor2 como id 2

	for {
		fmt.Println("\n\n ############# Bem-vindo ao Mercado de Energia! #############")

		exibeMenu()            //Exibe o menu de opções
		comando := leComando() //grava a opção digitada

		switch comando {
		case 1:
			screen.Clear()
			fmt.Println("###########################")
			fmt.Println("Dados dos fornecedores")
			fmt.Println("#############################")
			fmt.Printf("\nId: %d\nCapacidade de carga [kW]: %.2f\nEnergia gerada [kW]: %.2f\nEnergia fornecida [kW]: %.2f\nPreço minimo desejável [R$/kW]: %.2f\nDemanda Interna [kW]: %.2f\nPreço desejável [R$/kW]: %.2f\n", fornecedor1.Id, fornecedor1.CapacidadeCarga, fornecedor1.EnergiaGerada, fornecedor1.Energia_Fornecida, fornecedor1.MenorPreco, fornecedor1.Demanda_Interna, fornecedor1.PrecoDesejavel)
			fmt.Printf("\nId: %d\nCapacidade de carga [kW]: %.2f\nEnergia gerada [kW]: %.2f\nEnergia fornecida [kW]: %.2f\nPreço minimo desejável [R$/kW]: %.2f\nDemanda Interna [kW]: %.2f\nPreço desejável [R$/kW]: %.2f\n", fornecedor2.Id, fornecedor2.CapacidadeCarga, fornecedor2.EnergiaGerada, fornecedor2.Energia_Fornecida, fornecedor2.MenorPreco, fornecedor2.Demanda_Interna, fornecedor2.PrecoDesejavel)
			fmt.Printf("\nId: %d\nCapacidade de carga [kW]: %.2f\nEnergia gerada [kW]: %.2f\nEnergia fornecida [kW]: %.2f\nPreço minimo desejável [R$/kW]: %.2f\nDemanda Interna [kW]: %.2f\nPreço desejável [R$/kW]: %.2f\n", fornecedor3.Id, fornecedor3.CapacidadeCarga, fornecedor3.EnergiaGerada, fornecedor3.Energia_Fornecida, fornecedor3.MenorPreco, fornecedor3.Demanda_Interna, fornecedor3.PrecoDesejavel)
		case 2:
			screen.Clear()
			valida_existencia := consumidor1.Demanda //Verifica se já existe alguma demanda cadastrada, caso não ele solicita o cadastro
			if valida_existencia == 0 {
				fmt.Println("ERRO: Consumidor ainda não cadastrado")
				consumidor1.Inicia_EConsumidor()
				consumidor2.Inicia_EConsumidor()
			} else {
				fmt.Println("###########################")
				fmt.Println("Dados dos consumidores")
				fmt.Println("#############################")
				fmt.Printf("\nId: %d\nPrazo de contrato do Consumidor [s]: %.2f\nDemanda Consumidor [kW]: %.2f\nMáximo preço admissível [R$/kW]: %.2f\nTarifa desejável [R$/kW]: %.2f\n", consumidor1.Id, consumidor1.PrazoContrato, consumidor1.Demanda, consumidor1.PrecoMaximo, consumidor1.TarifaDesejavel)
				fmt.Printf("\nId: %d\nPrazo de contrato do Consumidor [s]: %.2f\nDemanda Consumidor [kW]: %.2f\nMáximo preço admissível [R$/kW]: %.2f\nTarifa desejável [R$/kW]: %.2f\n", consumidor2.Id, consumidor2.PrazoContrato, consumidor2.Demanda, consumidor2.PrecoMaximo, consumidor2.TarifaDesejavel)
			}
		case 3:
			screen.Clear()
			valida_existencia := consumidor1.Demanda //Verifica se já existe alguma demanda cadastrada, caso não ele solicita o cadastro
			if valida_existencia == 0 {
				fmt.Println("ERRO: Consumidor ainda não cadastrado")
				consumidor1.Inicia_EConsumidor()
				consumidor2.Inicia_EConsumidor()
			} else {
				printDate()
				//go
				//CRIAR THREADS AQUI
				fmt.Println("Iniciando simulação...")

				//qmsg.InicializaQmsg()
			}
		case 0: //Encerra o programa
			screen.Clear()
			fmt.Print("Encerrando o Mercado de Energia...\n")
			fmt.Println("Bye!")
			os.Exit(0)
		default: //Caso nenhum dos comandos acima seja selecionado, ele retorna um erro
			screen.Clear()
			fmt.Println("ERROR: command not found")
			os.Exit(-1)
		}
	}
}

func exibeIntroducao() {
	titulo := "Mercado de Energia"
	disciplina := "Interfaceamento e Drivers"
	versao := 1.0
	professora := "Bruna Fernandes Flesch"
	nomes := "Gabriel, Mauricio e Samuel"

	fmt.Println("Titulo: ", titulo)
	fmt.Println("Disciplina: ", disciplina)
	fmt.Println("Professora: ", professora)
	fmt.Println("Nomes: ", nomes)
	fmt.Println("Programa na versão: ", versao)
}

func exibeMenu() {
	fmt.Println("1 - Mostrar dados dos fornecedores")
	fmt.Println("2 - Mostrar dados dos consumidores")
	fmt.Println("3 - Executar simulação")
	fmt.Println("0 - Sair do programa")
	fmt.Print("\nOpção: ")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("")
	return comandoLido
}

func printDate() {
	currentTime := time.Now()
	fmt.Println(currentTime.Format("02/01/2006 15:04:05"))
}

func printThreads(action string, id int) {
	fmt.Printf("Thread #%d is %s\n", id+1, action)
}
