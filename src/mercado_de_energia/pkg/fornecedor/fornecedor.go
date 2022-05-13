package fornecedor

import (
	"context"
	"fmt"
	"math/rand"
	quadromensagens "mercado_de_energia/pkg/quadro_mensagens"
	"sync"
	"time"
)

// Criando um tipo de dados que será utilizado pelas funções  de calculo da equacao linear
type Efornecedor struct {
	Id               int                      // id
	PrecoDesejavel   float64                  // preco desejavel
	MenorPreco       float64                  // menor preco admissivel
	CapacidadeCarga  float64                  // capacidade de armazenamento
	EnergiaGerada    float64                  // energia gerada
	EnergiaFornecida float64                  // energia fornecida
	DemandaInterna   float64                  // demanda interna
	FazOferta        bool                     //variavel que indica se fornecedor fez uma oferta (se ==1)
	Oferta           *quadromensagens.MsgMerc // variavel que armazena a oferta
}

func (c *Efornecedor) Inicia_Efornecedor() { // Inicialização da estrutura de dados
	rand.Seed(time.Now().UnixNano()) //limpa buffer para geração de números aleatórios
	c.PrecoDesejavel = GetRandFloat(100, 200)
	c.MenorPreco = GetRandFloat(50, 100)
	c.EnergiaGerada = GetRandFloat(5000, 10000)
	c.EnergiaFornecida = 0
	c.DemandaInterna = GetRandFloat(1000, 1500)
	c.CapacidadeCarga = c.EnergiaGerada - c.DemandaInterna
	c.FazOferta = false
}

func (c *Efornecedor) AtualizaCapacidaDeCarga(energia_fornecida float64) { // Atualizacao Capacidade de Carga
	c.CapacidadeCarga -= energia_fornecida
	c.EnergiaFornecida += energia_fornecida
}

func (c *Efornecedor) AtualizaPrecoDesejavel() { // Atualizacao do pD
	c.MenorPreco -= (c.MenorPreco * 0.07)
	c.PrecoDesejavel = c.MenorPreco
}

func GetRandFloat(min, max float64) float64 { //Gerador de num aleatorios float
	return min + rand.Float64()*(max-min)
}

func (c *Efornecedor) WorkFornecedorOferta(ctx context.Context, q quadromensagens.QuadroMsg) {
	once := &sync.Once{}

	for {
		time.Sleep(time.Second * 1)
		select {
		case <-ctx.Done():
			return
		default:
			if c.FazOferta {
				if c.Oferta.Status == quadromensagens.Aceite {
					c.AtualizaCapacidaDeCarga(c.Oferta.CapacidadeFornecimento)
					c.Oferta.Clean()
					c.FazOferta = false
				} else if c.Oferta.Status == quadromensagens.Recusa {
					fmt.Println("\nFornecedor: ", c.Id, " - Oferta Recusada")
					c.Oferta.Status = quadromensagens.Oferta
					c.FazOferta = false
				}

			} else {
				for _, oferta := range q.Mensagem {
					if oferta == nil {
						continue
					}

					q.MuRW.Lock()

					if oferta.Status == quadromensagens.Oferta && oferta.CodigoFornecedor == -1 {
						fmt.Printf("\nFornecedor %d recebeu oferta do comprador  %d", c.Id, oferta.CodigoComprador)
						if oferta.DemandaSolicitada <= c.CapacidadeCarga && oferta.PrecoVenda <= c.PrecoDesejavel {
							oferta.Status = quadromensagens.Proposta
							oferta.CodigoFornecedor = c.Id

							if c.CapacidadeCarga > oferta.DemandaSolicitada {
								oferta.CapacidadeFornecimento = 2
							} else {
								oferta.CapacidadeFornecimento = c.CapacidadeCarga
							}

							c.Oferta = oferta
							c.FazOferta = true
						}
					}

					q.MuRW.Unlock()
				}
			}

			if c.CapacidadeCarga <= (c.EnergiaGerada * 0.1) {
				once.Do(c.AtualizaPrecoDesejavel)
			}
		}
	}
}

//Teste para print de threads
func PrintThreads(id int) {
	fmt.Printf("\nThread %d Execução em: ", id)
	printDate()
}
func printDate() { //Função para print data e hora atual
	currentTime := time.Now()
	fmt.Println(currentTime.Format("02/01/2006 15:04:05"))
}
