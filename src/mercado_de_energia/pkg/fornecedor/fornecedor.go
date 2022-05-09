package fornecedor

import (
	"context"
	"math/rand"
	quadromensagens "mercado_de_energia/pkg/quadro_mensagens"
	"time"
)

// Criando um tipo de dados que será utilizado pelas funções  de calculo da equacao linear
type Efornecedor struct {
	Id                int     // id
	PrecoDesejavel    float64 // preco desejavel
	MenorPreco        float64 // menor preco admissivel
	CapacidadeCarga   float64 // capacidade de armazenamento
	EnergiaGerada     float64 // energia gerada
	Energia_Fornecida float64 // energia fornecida
	Demanda_Interna   float64 // demanda interna
	FazOferta         int     //variavel que indica se fornecedor fez uma oferta (se ==1)
}

func (c *Efornecedor) Inicia_Efornecedor() { // Inicialização da estrutura de dados
	rand.Seed(time.Now().UnixNano()) //limpa buffer para geração de números aleatórios
	c.PrecoDesejavel = GetRandFloat(100, 200)
	c.MenorPreco = GetRandFloat(50, 100)
	c.CapacidadeCarga = c.EnergiaGerada - c.Demanda_Interna //- c.Energia_Fornecida //GetRandFloat(100, 200)
	c.EnergiaGerada = GetRandFloat(5000, 10000)
	c.Energia_Fornecida = 0
	c.Demanda_Interna = GetRandFloat(1000, 1500)
	c.FazOferta = 0
}

//void atualiza_pd(struct_Efornecedor *, double, double, double)
func (c *Efornecedor) AtualizaPrecoDesejavel() { // Atualizacao do pD
	c.MenorPreco -= (c.MenorPreco * 0.07)
	c.PrecoDesejavel = c.MenorPreco
}

func GetRandFloat(min, max float64) float64 { //Gerador de num aleatorios float
	return min + rand.Float64()*(max-min)
}

func (c *Efornecedor) WorkFornecedor(ctx context.Context, q quadromensagens.QuadroMsg) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			quadro := quadromensagens.MsgMerc{}
			if quadro.DemandaSolicitada <= c.CapacidadeCarga && quadro.PrecoVenda <= c.PrecoDesejavel {
				quadro.Status = quadromensagens.Oferta
				quadro.CodigoFornecedor = c.Id
				quadro.CapacidadeFornecimento = c.CapacidadeCarga
			}
			//get do quadro e validar valores e comprar e atualizar o quadro novamente
		}
	}
}
