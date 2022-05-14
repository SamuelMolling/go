package main

import (
	"context"
	"fmt"
	"mercado_de_energia/pkg/comprador"
	"mercado_de_energia/pkg/fornecedor"
	quadromensagens "mercado_de_energia/pkg/quadro_mensagens"
	"os"
	"time"

	screen "github.com/inancgumus/screen"
	"github.com/jedib0t/go-pretty/v6/table"
)

func main() {

	screen.Clear() //limpa tela

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
			fmt.Printf("\nId: %d\nCapacidade de carga [kW]: %.2f\nEnergia gerada [kW]: %.2f\nEnergia fornecida [kW]: %.2f\nPreço minimo desejável [R$/kW]: %.2f\nDemanda Interna [kW]: %.2f\nPreço desejável [R$/kW]: %.2f\n", fornecedor1.Id, fornecedor1.CapacidadeCarga, fornecedor1.EnergiaGerada, fornecedor1.EnergiaFornecida, fornecedor1.MenorPreco, fornecedor1.DemandaInterna, fornecedor1.PrecoDesejavel)
			fmt.Printf("\nId: %d\nCapacidade de carga [kW]: %.2f\nEnergia gerada [kW]: %.2f\nEnergia fornecida [kW]: %.2f\nPreço minimo desejável [R$/kW]: %.2f\nDemanda Interna [kW]: %.2f\nPreço desejável [R$/kW]: %.2f\n", fornecedor2.Id, fornecedor2.CapacidadeCarga, fornecedor2.EnergiaGerada, fornecedor2.EnergiaFornecida, fornecedor2.MenorPreco, fornecedor2.DemandaInterna, fornecedor2.PrecoDesejavel)
			fmt.Printf("\nId: %d\nCapacidade de carga [kW]: %.2f\nEnergia gerada [kW]: %.2f\nEnergia fornecida [kW]: %.2f\nPreço minimo desejável [R$/kW]: %.2f\nDemanda Interna [kW]: %.2f\nPreço desejável [R$/kW]: %.2f\n", fornecedor3.Id, fornecedor3.CapacidadeCarga, fornecedor3.EnergiaGerada, fornecedor3.EnergiaFornecida, fornecedor3.MenorPreco, fornecedor3.DemandaInterna, fornecedor3.PrecoDesejavel)
		case 2:
			screen.Clear()
			valida_existencia := consumidor1.Demanda //Verifica se já existe alguma demanda cadastrada, caso não ele solicita o cadastro
			if valida_existencia == 0 {
				fmt.Println("ERRO: Consumidor ainda não cadastrado!!!!\n\n")
				consumidor1.Inicia_EConsumidor()
				consumidor2.Inicia_EConsumidor()
			} else {
				fmt.Println("###########################")
				fmt.Println("Dados dos consumidores")
				fmt.Println("#############################")
				fmt.Printf("\nId: %d\nPrazo de contrato do Consumidor [s]: %d\nDemanda Consumidor [kW]: %.2f\nMáximo preço admissível [R$/kW]: %.2f\nTarifa desejável [R$/kW]: %.2f\n", consumidor1.Id, consumidor1.PrazoContrato, consumidor1.Demanda, consumidor1.PrecoMaximo, consumidor1.TarifaDesejavel)
				fmt.Printf("\nId: %d\nPrazo de contrato do Consumidor [s]: %d\nDemanda Consumidor [kW]: %.2f\nMáximo preço admissível [R$/kW]: %.2f\nTarifa desejável [R$/kW]: %.2f\n", consumidor2.Id, consumidor2.PrazoContrato, consumidor2.Demanda, consumidor2.PrecoMaximo, consumidor2.TarifaDesejavel)
			}
		case 3:
			screen.Clear()
			valida_existencia := consumidor1.Demanda //Verifica se já existe alguma demanda cadastrada, caso não ele solicita o cadastro
			if valida_existencia == 0 {
				fmt.Println("ERRO: Consumidor ainda não cadastrado!!!!\n\n")
				consumidor1.Inicia_EConsumidor()
				consumidor2.Inicia_EConsumidor()
			} else {
				fmt.Println("Iniciando simulação...")
				quadro := quadromensagens.QuadroMsg{}                               //Cria um quadro
				quadro.InicializaQmsg()                                             //Inicializa o quadro
				ctx, _ := context.WithTimeout(context.Background(), 60*time.Second) //Cria um contexto de 120 segunds
				// go func() {                                                         //Thread pra debug
				// 	for {
				// 		printDbg(
				// 			quadro,
				// 			[]comprador.EConsumidor{consumidor1, consumidor2},
				// 			[]fornecedor.Efornecedor{fornecedor1, fornecedor2, fornecedor3})
				// 		time.Sleep(1 * time.Second)
				// 	}
				// }()
				//Cria as threads
				go consumidor1.WorkConsumidor(ctx, quadro)
				go consumidor2.WorkConsumidor(ctx, quadro)
				go fornecedor1.WorkFornecedorOferta(ctx, quadro)
				go fornecedor2.WorkFornecedorOferta(ctx, quadro)
				go fornecedor3.WorkFornecedorOferta(ctx, quadro)

				<-ctx.Done()
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

func printDbg(quadro quadromensagens.QuadroMsg, consumidores []comprador.EConsumidor, fornecedores []fornecedor.Efornecedor) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Id", "Prazo de Contrato", "Demanda", "Preço Máximo", "Tarifa Desejável", "Oferta Aberta"})
	for _, consumidor := range consumidores {
		t.AppendRow([]interface{}{consumidor.Id, consumidor.PrazoContrato, consumidor.Demanda, consumidor.PrecoMaximo, consumidor.TarifaDesejavel, consumidor.OfertaAberta})
	}

	t.Render()

	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Id", "Capacidade de Carga", "Energia Gerada", "Energia Fornecida", "Demanda Interna", "Preço Desejável", "Faz Oferta"})
	for _, fornecedor := range fornecedores {
		t.AppendRow([]interface{}{fornecedor.Id, fornecedor.CapacidadeCarga, fornecedor.EnergiaGerada, fornecedor.EnergiaFornecida, fornecedor.DemandaInterna, fornecedor.PrecoDesejavel, fornecedor.FazOferta})
	}

	t.Render()

	t = table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	t.AppendHeader(table.Row{"Codigo Fornecedor", "Preco Venda", "Capacidade Fornecimento", "Codigo Comprador", "Demanda Solicitada", "Status"})
	for _, msg := range quadro.Mensagem {
		if msg == nil {
			continue
		}

		status := quadromensagens.MsgStatusString[msg.Status]
		t.AppendRow([]interface{}{msg.CodigoFornecedor, msg.PrecoVenda, msg.CapacidadeFornecimento, msg.CodigoComprador, msg.DemandaSolicitada, status})
	}

	t.Render()
}

func exibeIntroducao() { //Função para exibir informações de introdução
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

func exibeMenu() { //Função para exibir opções do menu
	fmt.Println("1 - Mostrar dados dos fornecedores")
	fmt.Println("2 - Mostrar dados dos consumidores")
	fmt.Println("3 - Executar simulação")
	fmt.Println("0 - Sair do programa")
	fmt.Print("\nOpção: ")
}

func leComando() int { //Função para salvar a opção desejada do menu
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("")
	return comandoLido
}
