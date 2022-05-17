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
	Id               int     // id
	PrecoDesejavel   float64 // preco desejavel
	MenorPreco       float64 // menor preco admissivel
	CapacidadeCarga  float64 // capacidade de armazenamento
	EnergiaGerada    float64 // energia gerada
	EnergiaFornecida float64 // energia fornecida
	DemandaInterna   float64 // demanda interna
	FazOferta        bool    //variavel que indica se fornecedor fez uma oferta
	OfertaId         int     // id da oferta
}

func (c *Efornecedor) Inicia_Efornecedor() { // Inicialização da estrutura de dados
	rand.Seed(time.Now().UnixNano()) //limpa buffer para geração de números aleatórios
	c.PrecoDesejavel = GetRandFloat(100, 130)
	c.MenorPreco = GetRandFloat(50, 90)
	c.EnergiaGerada = GetRandFloat(5000, 10000)
	c.EnergiaFornecida = 0
	c.DemandaInterna = GetRandFloat(1000, 1500)
	c.CapacidadeCarga = c.EnergiaGerada - c.DemandaInterna
	c.FazOferta = false
	c.OfertaId = -1
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

func (c *Efornecedor) validaPropostaFeita(q *quadromensagens.QuadroMsg) {
	oferta := q.Mensagem[c.OfertaId]
	q.MensagemLock[c.OfertaId].Lock()
	defer q.MensagemLock[c.OfertaId].Unlock()
	if oferta.Status == quadromensagens.Aceite {
		c.AtualizaCapacidaDeCarga(oferta.CapacidadeFornecimento)
		oferta.Clean() //limpa oferta do aceite
		c.FazOferta = false
		c.OfertaId = -1
		fmt.Println("\nFornecedor:", c.Id, "- Teve oferta Aceita") //print mostrando qual fornecedor teve oferta aceita
	} else if oferta.Status == quadromensagens.Recusa { //se oferta é recusada
		fmt.Println("\nFornecedor:", c.Id, "- Teve oferta Recusada") //print mostrando qual fornecedor teve oferta recusada
		oferta.CleanFornecedor()                                     //limpa oferta do aceite
		c.FazOferta = false
		c.OfertaId = -1
	}

}

func (c *Efornecedor) fazProposta(q *quadromensagens.QuadroMsg) {

	for id, oferta := range q.Mensagem {
		if oferta.Status == quadromensagens.Livre {
			continue
		}
		q.MensagemLock[id].Lock()
		if oferta.Status == quadromensagens.Oferta && oferta.CodigoFornecedor == -1 { //Valida se tem uma oferta do comprador
			if oferta.DemandaSolicitada <= c.CapacidadeCarga && oferta.PrecoVenda <= c.PrecoDesejavel {
				fmt.Printf("\nFornecedor %d mandou oferta para o comprador  %d", c.Id, oferta.CodigoComprador)
				if c.CapacidadeCarga > oferta.DemandaSolicitada {
					oferta.CapacidadeFornecimento = GetRandFloat(100.15, 130.60) //limita o fornecimento para um randfloat
				} else {
					oferta.CapacidadeFornecimento = c.CapacidadeCarga //fornecer o que tem, não a demanda solicitada
				}
				oferta.Status = quadromensagens.Proposta
				oferta.CodigoFornecedor = c.Id
				oferta.PrecoVenda = c.PrecoDesejavel
				c.OfertaId = id
				c.FazOferta = true
				q.MensagemLock[id].Unlock()
				return
			}
		}
		q.MensagemLock[id].Unlock()
	}
}

func (c *Efornecedor) WorkFornecedorOferta(ctx context.Context, q *quadromensagens.QuadroMsg) {
	once := &sync.Once{} //Cria um type Once

	rand.Seed(time.Now().UnixNano())
	for {
		time.Sleep(time.Second * time.Duration(rand.Float64()+1))
		select {
		case <-ctx.Done():
			return
		default:
			if c.FazOferta && c.OfertaId >= 0 { //Valida se fazoferta é igual a true
				c.validaPropostaFeita(q) //Valida se a oferta foi aceita ou recusada
			} else if c.OfertaId == -1 {
				c.fazProposta(q) //Faz a proposta
			}
		}

		if c.CapacidadeCarga <= (c.EnergiaGerada * 0.1) { //se a capacidade de carga for 10% menor que a energia gerada, ele atualiza o preco
			once.Do(c.AtualizaPrecoDesejavel) //Onde.Do executa uma vez a atualização por "if"
		}
	}
}
