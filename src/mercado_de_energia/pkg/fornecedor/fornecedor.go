package fornecedor

import (
	"context"
	"fmt"
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
	c.CapacidadeCarga = c.EnergiaGerada - c.Demanda_Interna
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

func (c *Efornecedor) WorkFornecedorOferta(ctx context.Context, q quadromensagens.QuadroMsg) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			quadro := quadromensagens.MsgMerc{}
			if c.CapacidadeCarga <= (c.EnergiaGerada * 0.1) {
				c.AtualizaPrecoDesejavel()
			}
			if quadro.DemandaSolicitada <= c.CapacidadeCarga && quadro.PrecoVenda <= c.PrecoDesejavel {
				quadro.Status = quadromensagens.Oferta
				quadro.CodigoFornecedor = c.Id
				quadro.CapacidadeFornecimento = c.CapacidadeCarga
				q.MuRW.Lock()
				if len(q.Mensagem) < 8 { //valida se o tamanho do array é menor que 8
					q.Mensagem = append(q.Mensagem, quadro)
					go PrintThreads(c.Id)
				}
				q.MuRW.Unlock() //Desbloqueia o Mutex
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
