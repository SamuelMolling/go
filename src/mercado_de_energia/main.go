package main

import (
	"context"
	"fmt"
	"math"
	"mercado_de_energia/pkg/comprador"
	"mercado_de_energia/pkg/fornecedor"
	"mercado_de_energia/pkg/gopherdance"
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
		exibeMenu()            //Exibe o menu de opções
		comando := leComando() //grava a opção digitada

		switch comando {
		case 1:
			screen.Clear()
			printFornecedor([]fornecedor.Efornecedor{fornecedor1, fornecedor2, fornecedor3})
		case 2:
			screen.Clear()
			valida_existencia := consumidor1.Demanda //Verifica se já existe alguma demanda cadastrada, caso não ele solicita o cadastro
			if valida_existencia == 0 {
				fmt.Println("ERRO: Consumidor ainda não cadastrado!!!!\n\n")
				consumidor1.Inicia_EConsumidor()
				consumidor2.Inicia_EConsumidor()
			} else {
				printConsumidor([]comprador.EConsumidor{consumidor1, consumidor2})
			}
		case 3:
			screen.Clear()
			valida_existencia := consumidor1.Demanda //Verifica se já existe alguma demanda cadastrada, caso não ele solicita o cadastro
			if valida_existencia == 0 {
				fmt.Println("ERRO: Consumidor ainda não cadastrado!!!!\n\n")
				consumidor1.Inicia_EConsumidor()
				consumidor2.Inicia_EConsumidor()
			}
			fmt.Println("Iniciando simulação...")
			quadro := quadromensagens.QuadroMsg{}                                //Cria um quadro
			quadro.InicializaQmsg()                                              //Inicializa o quadro
			ctx, _ := context.WithTimeout(context.Background(), 120*time.Second) //Cria um contexto de 120 segundos
			// go func() {                                                          //Thread pra debug
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
		case 99: //Easter egg
			gopherdance.Main()
			return
		case 0: //Encerra o programa
			screen.Clear()
			fmt.Print("Encerrando o Mercado de Energia...\n")
			fmt.Println("Bye!")
			os.Exit(0)
		default: //Caso nenhum dos comandos acima seja selecionado, ele retorna um erro
			screen.Clear()
			fmt.Println("ERROR: command not found")
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
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Titulo", "Disciplina", "Versão", "Professora", "Nomes"})

	t.AppendRow([]interface{}{"Mercado Livre de Energia", "Interfaceamento e Drivers", "1.0", "Bruna Fernandes Flesch", "Gabriel, Mauricio e Samuel"})
	t.Render()
}

func exibeMenu() { //Função para exibir opções do menu
	descrição := [5]string{
		"Sair do programa",
		"Mostrar dados dos fornecedores",
		"Mostrar dados dos consumidores",
		"Executar simulação"}
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Opção", "Descrição"})
	for i := 0; i < 4; i++ {
		t.AppendRow([]interface{}{i, descrição[i]})
	}
	t.Render()
	fmt.Print("\nOpção: ")
}

func leComando() int { //Função para salvar a opção desejada do menu
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("")
	return comandoLido
}

func printFornecedor(fornecedores []fornecedor.Efornecedor) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Capacidade de Carga", "Energia Gerada", "Energia Fornecida", "Demanda Interna", "Preço Desejável"})
	for _, fornecedor := range fornecedores {
		t.AppendRow([]interface{}{fornecedor.Id, toFixed(fornecedor.CapacidadeCarga, 2), toFixed(fornecedor.EnergiaGerada, 2), toFixed(fornecedor.EnergiaFornecida, 2), toFixed(fornecedor.DemandaInterna, 2), toFixed(fornecedor.PrecoDesejavel, 2)})
	}
	t.Render()
}

func printConsumidor(consumidores []comprador.EConsumidor) {
	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Id", "Prazo de contrato [s]", "Demanda Consumidor [kW]", "Máximo preço admissível [R$/kW]", "Tarifa desejável [R$/kW]"})
	for _, consumidor := range consumidores {
		t.AppendRow([]interface{}{consumidor.Id, consumidor.PrazoContrato, toFixed(consumidor.Demanda, 2), toFixed(consumidor.PrecoMaximo, 2), toFixed(consumidor.TarifaDesejavel, 2)})
	}
	t.Render()
}

func round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func toFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(round(num*output)) / output
}
